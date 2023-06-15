package commonkey

type CtxKey string

const SCOPE_CTX_KEY CtxKey = "scope-ctx"

const ROLE_ADMIN = 1
const ROLE_USER = 2

var TRX_TYPE_MAP = map[int64]string{
	0: "P2P",
	1: "PAYMENT",
}
