package domain

import (
	"github.com/alitdarmaputra/fims-be/src/business/domain/token"
	"github.com/alitdarmaputra/fims-be/src/business/domain/user"
)

type Domain struct {
	User  user.UserDom
	Token token.TokenDom
}

func Init() *Domain {
	userDom := user.InitUserDom()
	tokenDom := token.InitTokenDom()

	return &Domain{
		User:  userDom,
		Token: tokenDom,
	}
}
