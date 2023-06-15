package dto

type AccountQueryParams struct {
	AccountID int64 `param:"accountID"`
}

type AccountRequest struct {
	CustomerID  int64   `json:"customer_id"`
	AccountNo   string  `json:"account_no"`
	AccountType int64   `json:"account_type"`
	CardNumber  string  `json:"card_number"`
	CVV         string  `json:"cvv"`
	PIN         string  `json:"pin"`
	Balance     float32 `json:"balance"`
}

type AccountResponse struct {
	AccountID   int64   `json:"account_id"`
	AccountNo   string  `json:"account_no"`
	AccountType int64   `json:"account_type"`
	CardNumber  string  `json:"card_number"`
	Balance     float32 `json:"balance"`
}

type TransactionRequest struct {
	AccountID      int64   `json:"account_id"`
	RecipientID    int64   `json:"recipient_id"`
	TrxType        int64   `json:"trx_type"`
	TrxDatetime    string  `json:"trx_datetime"`
	TrxStatus      int64   `json:"trx_status"`
	Nominal        float32 `json:"nominal"`
	TransactionFee float32 `json:"transaction_fee"`
}

type TransactionResponse struct {
	ID            int64   `db:"trx_id"`
	AccountID     int64   `db:"account_id"`
	AccountOwner  string  `db:"account_owner"`
	RecipientID   int64   `db:"recipient_id"`
	RecipientName string  `db:"recipient_name"`
	TrxType       string  `db:"trx_type"`
	TrxDatetime   string  `db:"trx_datetime"`
	Nominal       float32 `db:"nominal"`
	TrxFee        float32 `db:"trx_fee"`
}
