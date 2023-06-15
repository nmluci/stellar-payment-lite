package router

import (
	"github.com/labstack/echo/v4"
	ecMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/nmluci/stellar-payment-lite/cmd/webservice/handler"
	"github.com/nmluci/stellar-payment-lite/internal/config"
	"github.com/nmluci/stellar-payment-lite/internal/middleware"
	"github.com/nmluci/stellar-payment-lite/internal/service"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
	"github.com/rs/zerolog"
)

type InitRouterParams struct {
	Logger  zerolog.Logger
	Service service.Service
	Ec      *echo.Echo
	Conf    *config.Config
}

func HandleNotImplemented(c echo.Context) error {
	c.Response().Header().Add("Secret-Msg", "nothing to see here")

	return echttputil.WriteErrorResponse(c, errs.ErrRouteNotImplemented)
}

func Init(params *InitRouterParams) {
	params.Ec.Use(ecMiddleware.CORS(), middleware.ServiceVersioner)

	plainRouter := params.Ec.Group("")
	secureRouter := params.Ec.Group("", middleware.AuthorizationMiddleware(params.Service))

	plainRouter.GET(pingPath, handler.HandlePing(params.Service.Ping))
	plainRouter.OPTIONS(pingPath, handler.HandlePing(params.Service.Ping))

	// Auth
	plainRouter.POST(loginPath, handler.HandleAuthLogin(params.Service.AuthLogin))
	plainRouter.OPTIONS(loginPath, handler.HandleAuthLogin(params.Service.AuthLogin))

	// Users
	plainRouter.POST(UserRegisterPath, handler.HandleRegisterUser(params.Service.RegisterUser))
	plainRouter.OPTIONS(UserRegisterPath, handler.HandleRegisterUser(params.Service.RegisterUser))
	secureRouter.GET(UserIDBasepath, handler.HandleGetUserDetailByID(params.Service.GetUserDetailByID))
	secureRouter.OPTIONS(UserIDBasepath, handler.HandleGetUserDetailByID(params.Service.GetUserDetailByID))
	secureRouter.PUT(UserIDBasepath, handler.HandleUpdateUserDetail(params.Service.UpdateUserByID))
	secureRouter.OPTIONS(UserIDBasepath, handler.HandleUpdateUserDetail(params.Service.UpdateUserByID))

	// Customers
	secureRouter.GET(CustomerIDBasepath, handler.HandleGetCustomerDetailByID(params.Service.GetCustomerByID))
	secureRouter.OPTIONS(CustomerIDBasepath, handler.HandleGetCustomerDetailByID(params.Service.GetCustomerByID))
	secureRouter.PUT(CustomerIDBasepath, handler.HandleUpdateCustomerDetail(params.Service.UpdateCustomer))
	secureRouter.OPTIONS(CustomerIDBasepath, handler.HandleUpdateCustomerDetail(params.Service.UpdateCustomer))

	// Accounts
	secureRouter.GET(AccountBasepath, handler.HandleGetAccountByUser(params.Service.GetAccountByUser))
	secureRouter.OPTIONS(AccountBasepath, handler.HandleGetAccountByUser(params.Service.GetAccountByUser))
	secureRouter.POST(AccountBasepath, handler.HandleCreateAccount(params.Service.CreateAccount))
	secureRouter.OPTIONS(AccountBasepath, handler.HandleCreateAccount(params.Service.CreateAccount))
	secureRouter.GET(AccountTransactionPath, handler.HandleGetTransactionHistoryByUser(params.Service.GetTransactionHistoryByUser))
	secureRouter.OPTIONS(AccountTransactionPath, handler.HandleGetTransactionHistoryByUser(params.Service.GetTransactionHistoryByUser))

	// Merchants
	secureRouter.GET(MerchantBasepath, handler.HandleGetMerchant(params.Service.GetMerchants))
	secureRouter.OPTIONS(MerchantBasepath, handler.HandleGetMerchant(params.Service.GetMerchants))
	secureRouter.POST(MerchantBasepath, handler.HandleCreateMerchant(params.Service.CreateMerchant))
	secureRouter.OPTIONS(MerchantBasepath, handler.HandleCreateMerchant(params.Service.CreateMerchant))
	secureRouter.GET(MerchantSettlementPath, handler.HandleFindMerchantSettlements(params.Service.FindMerchantSettlements))
	secureRouter.OPTIONS(MerchantSettlementPath, handler.HandleFindMerchantSettlements(params.Service.FindMerchantSettlements))

	// Transactions
	secureRouter.POST(TransactionP2PPath, handler.HandleCreateTransactionP2P(params.Service.CreateTransactionP2P))
	secureRouter.OPTIONS(TransactionP2PPath, handler.HandleCreateTransactionP2P(params.Service.CreateTransactionP2P))
	secureRouter.POST(TransactionMerchantPath, handler.HandleCreateTransactionMerchant(params.Service.CreateTransactionMerchant))
	secureRouter.OPTIONS(TransactionMerchantPath, handler.HandleCreateTransactionMerchant(params.Service.CreateTransactionMerchant))

	plainRouter.Any("*", HandleNotImplemented)
}
