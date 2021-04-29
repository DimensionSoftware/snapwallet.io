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
		SourceName:     t.SourceName,
		DestName:       t.DestName,
		Message:        t.Message,
		ExchangeRate:   t.ExchangeRate,
		Fees:           t.Fees,
		Status:         t.Status,
		BlockchainTxId: t.BlockchanTxID,
		CreatedAt:      formatEpochMSAsRFC3339(t.CreatedAt),
		ClosedAt:       formatEpochMSAsRFC3339(t.ClosedAt),
	}
}

func formatEpochMSAsRFC3339(epochMS int64) string {
	if epochMS == 0 {
		return ""
	}

	return time.Unix(epochMS/1000, 0).Format(time.RFC3339)
}
