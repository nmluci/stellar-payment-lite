package indto

import "time"

type TranasctionHistory struct {
	ID            int64     `db:"trx_id"`
	AccountID     int64     `db:"account_id"`
	AccountOwner  string    `db:"account_owner"`
	RecipientID   int64     `db:"recipient_id"`
	RecipientName string    `db:"recipient_name"`
	TrxType       int64     `db:"trx_type"`
	TrxDatetime   time.Time `db:"trx_datetime"`
	Nominal       float32   `db:"nominal"`
	TrxFee        float32   `db:"trx_fee"`
}

type SettlementDetail struct {
	ID           int64     `db:"settlement_id"`
	TrxID        int64     `db:"trx_id"`
	TrxDatetime  time.Time `db:"trx_datetime"`
	MerchantID   int64     `db:"merchant_id"`
	MerchantName string    `db:"merchant_name"`
	Nominal      float32   `db:"nominal"`
	Status       int64     `db:"status"`
}
