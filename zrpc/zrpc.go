package zrpc

import (
	"net"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

func NewZRPCFromConf(confPath string) (*rpcclient.Client, error) {
	cfg, err := ini.Load(confPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	rpcaddr := cfg.Section("").Key("rpcbind").String()
	rpcport := cfg.Section("").Key("rpcport").String()
	username := cfg.Section("").Key("rpcuser").String()
	password := cfg.Section("").Key("rpcpassword").String()

	return NewZRPCFromCreds(net.JoinHostPort(rpcaddr, rpcport), username, password)
}

func NewZRPCFromCreds(addr, username, password string) (*rpcclient.Client, error) {
	// Connect to local zcash RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         addr,
		User:         username,
		Pass:         password,
		HTTPPostMode: true, // Zcash only supports HTTP POST mode
		DisableTLS:   true, // Zcash does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	return rpcclient.New(connCfg, nil)
}
