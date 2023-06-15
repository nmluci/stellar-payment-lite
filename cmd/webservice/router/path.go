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
)
