package transformer

import (
	"encoding/json"
	"testing"

	"github.com/revolutionchain/charon/pkg/internal"
	"github.com/revolutionchain/charon/pkg/revo"
	"github.com/revolutionchain/charon/pkg/utils"
)

func initializeProxyETHGetBlockByHash(revoClient *revo.Revo) ETHProxy {
	return &ProxyETHGetBlockByHash{revoClient}
}

func TestGetBlockByHashRequestNonceLength(t *testing.T) {
	if len(utils.RemoveHexPrefix(internal.GetTransactionByHashResponse.Nonce)) != 16 {
		t.Errorf("Nonce test data should be zero left padded length 16")
	}
}

func TestGetBlockByHashRequest(t *testing.T) {
	testETHProxyRequest(
		t,
		initializeProxyETHGetBlockByHash,
		[]json.RawMessage{[]byte(`"` + internal.GetTransactionByHashBlockHexHash + `"`), []byte(`false`)},
		&internal.GetTransactionByHashResponse,
	)
}

func TestGetBlockByHashTransactionsRequest(t *testing.T) {
	testETHProxyRequest(
		t,
		initializeProxyETHGetBlockByHash,
		[]json.RawMessage{[]byte(`"` + internal.GetTransactionByHashBlockHexHash + `"`), []byte(`true`)},
		&internal.GetTransactionByHashResponseWithTransactions,
	)
}
