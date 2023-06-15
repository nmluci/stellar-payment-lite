package dto

type MerchantQueryParams struct {
	MerchantID int64 `param:"merchantID"`
}

type MerchantRequest struct {
	Name         string `json:"merchant"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	MerchantCode string `json:"merchant_code"`
}

type MerchantResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"merchant"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	MerchantCode string `json:"merchant_code"`
}

type MerchantSettlement struct {
	ID           int64   `json:"settlement_id"`
	TrxID        int64   `json:"trx_id"`
	TrxDatetime  string  `json:"trx_datetime"`
	MerchantID   int64   `json:"merchant_id"`
	MerchantName string  `json:"merchant_name"`
	Nominal      float32 `json:"nominal"`
	Status       int64   `json:"status"`
}
