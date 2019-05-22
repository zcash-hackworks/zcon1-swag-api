package swag

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/zcash-hackworks/zcon1-swag-api/api"
)

const (
	swagFromAddress   = "TODO HARDCODED FROM ADDRESS"
	swagDefaultAmount = 1
)

type Server struct {
	rpc *rpcclient.Client
}

func NewServer(dbPath string, rpc *rpcclient.Client) (api.SwagAPIServer, error) {
	// TODO: CREATE AND POPULATE THE DATABASE
	return &Server{
		rpc: rpc,
	}, nil
}

func (s *Server) Redeem(ctx context.Context, req *api.Request) (*api.Response, error) {

	// TODO: QUERY THE DATABASE

	return s.TriggerSend(ctx, req)
}

func (s *Server) GracefulStop() {}

type zSendManyDest struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

func (s *Server) TriggerSend(ctx context.Context, req *api.Request) (*api.Response, error) {
	// z_sendmany "fromaddress" [{"address":... ,"amount":...},...] ( minconf ) ( fee )
	//
	// Send multiple times. Amounts are decimal numbers with at most 8 digits of precision.
	// Change generated from a taddr flows to a new taddr address, while change generated from a zaddr returns to itself.
	// When sending coinbase UTXOs to a zaddr, change is not allowed. The entire value of the UTXO(s) must be consumed.
	// Before Sapling activates, the maximum number of zaddr outputs is 54 due to transaction size limits.
	//
	// Arguments:
	// 1. "fromaddress"         (string, required) The taddr or zaddr to send the funds from.
	// 2. "amounts"             (array, required) An array of json objects representing the amounts to send.
	//     [{
	//       "address":address  (string, required) The address is a taddr or zaddr
	//       "amount":amount    (numeric, required) The numeric amount in ZEC is the value
	//       "memo":memo        (string, optional) If the address is a zaddr, raw data represented in hexadecimal string format
	//     }, ... ]
	// 3. minconf               (numeric, optional, default=1) Only use funds confirmed at least this many times.
	// 4. fee                   (numeric, optional, default=0.0001) The fee amount to attach to this transaction.
	//
	// Result:
	// "operationid"          (string) An operationid to pass to z_getoperationstatus to get the result of the operation.
	//
	// Examples:
	// > zcash-cli z_sendmany "t1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd" '[{"address": "ztfaW34Gj9FrnGUEf833ywDVL62NWXBM81u6EQnM6VR45eYnXhwztecW1SjxA7JrmAXKJhxhj3vDNEpVCQoSvVoSpmbhtjf" ,"amount": 5.0}]'
	// > curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "z_sendmany", "params": ["t1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd", [{"address": "ztfaW34Gj9FrnGUEf833ywDVL62NWXBM81u6EQnM6VR45eYnXhwztecW1SjxA7JrmAXKJhxhj3vDNEpVCQoSvVoSpmbhtjf" ,"amount": 5.0}]] }' -H 'content-type: text/plain;' http://127.0.0.1:8232/

	// Construct raw JSON-RPC params. This is VERY UGLY because the btcd library lacks direct support for our RPC calls.

	zDests := make([]zSendManyDest, 1)
	zDests[0] = zSendManyDest{
		Address: req.Address,
		Amount:  swagDefaultAmount,
	}

	encodedDests, err := json.Marshal(zDests)
	if err != nil {
		return &api.Response{
			Success: false,
			Msg:     "Couldn't initiate transaction successfully.",
		}, err
	}

	params := make([]json.RawMessage, 2)
	params[0] = json.RawMessage("\"" + swagFromAddress + "\"")
	params[1] = json.RawMessage(encodedDests)
	result, rpcErr := s.rpc.RawRequest("z_sendmany", params)

	// For some reason, the error responses are not JSON
	if rpcErr != nil {
		errParts := strings.SplitN(rpcErr.Error(), ":", 2)
		errMsg := strings.TrimSpace(errParts[1])
		_, err = strconv.ParseInt(errParts[0], 10, 32)
		if err != nil {
			// This should never happen. We can't panic here, but it's that class of error.
			// This is why we need integration testing to work better than regtest currently does. TODO.
			return nil, errors.New("SendTransaction couldn't parse error code")
		}

		return &api.Response{
			Success: false,
			Msg:     errMsg,
		}, rpcErr
	}

	return &api.Response{
		Success: true,
		Msg:     string(result),
	}, nil
}
