package swag

import (
	"context"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/zcash-hackworks/zcon1-swag-api/api"
)

type Server struct {
}

func NewServer(dbPath string, zc *rpcclient.Client) (api.SwagAPIServer, error) {
	// TODO: CREATE AND POPULATE THE DATABASE
	return &Server{}, nil
}

func (s *Server) Redeem(ctx context.Context, req *api.Request) (*api.Response, error) {

	// TODO: QUERY THE DATABASE

	// TODO: SEND THE TRANSACTION IN RESPONSE

	return &api.Response{
		Success: true,
		Msg:     "",
	}, nil
}

func (s *Server) GracefulStop() {}

// SendTransaction forwards raw transaction bytes to a zcashd instance over JSON-RPC
//func (s *Server) SendTransaction(ctx context.Context, rawtx *walletrpc.RawTransaction) (*walletrpc.SendResponse, error) {
//	// sendrawtransaction "hexstring" ( allowhighfees )
//	//
//	// Submits raw transaction (serialized, hex-encoded) to local node and network.
//	//
//	// Also see createrawtransaction and signrawtransaction calls.
//	//
//	// Arguments:
//	// 1. "hexstring"    (string, required) The hex string of the raw transaction)
//	// 2. allowhighfees    (boolean, optional, default=false) Allow high fees
//	//
//	// Result:
//	// "hex"             (string) The transaction hash in hex

//	// Construct raw JSON-RPC params
//	params := make([]json.RawMessage, 1)
//	txHexString := hex.EncodeToString(rawtx.Data)
//	params[0] = json.RawMessage("\"" + txHexString + "\"")
//	result, rpcErr := s.client.RawRequest("sendrawtransaction", params)

//	var err error
//	var errCode int64
//	var errMsg string

//	// For some reason, the error responses are not JSON
//	if rpcErr != nil {
//		errParts := strings.SplitN(rpcErr.Error(), ":", 2)
//		errMsg = strings.TrimSpace(errParts[1])
//		errCode, err = strconv.ParseInt(errParts[0], 10, 32)
//		if err != nil {
//			// This should never happen. We can't panic here, but it's that class of error.
//			// This is why we need integration testing to work better than regtest currently does. TODO.
//			return nil, errors.New("SendTransaction couldn't parse error code")
//		}
//	} else {
//		errMsg = string(result)
//	}

//	// TODO these are called Error but they aren't at the moment.
//	// A success will return code 0 and message txhash.
//	return &walletrpc.SendResponse{
//		ErrorCode:    int32(errCode),
//		ErrorMessage: errMsg,
//	}, nil
//}
