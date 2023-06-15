package model

import "time"

type Transaction struct {
	ID             int64     `db:"id"`
	AccountID      int64     `db:"account_id"`
	RecipientID    int64     `db:"recipient_id"`
	TrxType        int64     `db:"trx_type"`
	TrxDatetime    time.Time `db:"trx_datetime"`
	TrxStatus      int64     `db:"trx_status"`
	Nominal        float32   `db:"nominal"`
	TransactionFee float32   `db:"transaction_fee"`
}

type Settlement struct {
	ID            int64   `db:"id"`
	TransactionID int64   `db:"transaction_id"`
	MerchantID    int64   `db:"merchant_id"`
	Nominal       float32 `db:"nominal"`
	Status        int64   `db:"status"`
}

type Merchant struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	Address      string `db:"address"`
	Phone        string `db:"phone"`
	MerchantCode string `db:"merchant_code"`
}
