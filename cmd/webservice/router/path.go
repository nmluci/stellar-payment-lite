package router

const (
	basePath = "/v1"
	pingPath = basePath + "/ping"

	// Auth
	authBasepath = basePath + "/auth"
	loginPath    = authBasepath + "/login"

	// User
	UserBasepath     = basePath + "/users"
	UserIDBasepath   = UserBasepath + "/:userID"
	UserRegisterPath = UserBasepath + "/register"

	// Customers
	CustomerBasepath   = basePath + "/customers"
	CustomerIDBasepath = CustomerBasepath + "/:customerID"

	// Accounts
	AccountBasepath        = basePath + "/accounts"
	AccountIDBasepath      = AccountBasepath + "/:accountID"
	AccountTransactionPath = AccountIDBasepath + "/transactions"

	// Merchants
	MerchantBasepath       = basePath + "/merchants"
	MerchantIDBasepath     = MerchantBasepath + "/:merchantID"
	MerchantSettlementPath = MerchantIDBasepath + "/settlements"

	// Transaction
	TransactionBasepath     = basePath + "/transactions"
	TransactionP2PPath      = TransactionBasepath + "/p2p"
	TransactionMerchantPath = TransactionBasepath + "/merchants"
)
