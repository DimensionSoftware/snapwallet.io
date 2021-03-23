package wyre

import (
	"strings"
	"time"

	proto "github.com/khoerling/flux/api/lib/protocol"
)

func WyreTransferToProto(t *Transfer) *proto.WyreTransfer {
	return &proto.WyreTransfer{
		Id:             t.ID,
		Source:         strings.Split(t.Source, ":")[1],
		Dest:           strings.Split(t.Dest, ":")[1],
		SourceCurrency: t.SourceCurrency,
		DestCurrency:   t.DestCurrency,
		SourceAmount:   t.SourceAmount,
		DestAmount:     t.DestAmount,
		ExchangeRate:   t.ExchangeRate,
		Fees:           t.Fees,
		Blockhash:      t.BlockchainTx.Blockhash,
		NetworkTxId:    t.BlockchainTx.NetworkTxID,
		Status:         t.Status,
		CreatedAt:      time.Unix(t.CreatedAt, 0).Format(time.RFC3339),
		ExpiresAt:      time.Unix(t.ExpiresAt, 0).Format(time.RFC3339),
	}
}