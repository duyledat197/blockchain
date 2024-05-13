package dispatcher

import "math/big"

type DispatcherTxRequest struct {
	PrivKey string   `json:"priv_key,omitempty"`
	From    string   `json:"from,omitempty"`
	To      string   `json:"to,omitempty"`
	Amount  *big.Int `json:"amount,omitempty"`
}
