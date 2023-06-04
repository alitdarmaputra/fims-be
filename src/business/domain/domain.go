package domain

import (
	"github.com/alitdarmaputra/fims-be/src/business/domain/node"
	"github.com/alitdarmaputra/fims-be/src/business/domain/status"
	"github.com/alitdarmaputra/fims-be/src/business/domain/token"
	"github.com/alitdarmaputra/fims-be/src/business/domain/user"
)

type Domain struct {
	User   user.UserDom
	Token  token.TokenDom
	Node   node.NodeDom
	Status status.StatusDom
}

func Init() *Domain {
	userDom := user.InitUserDom()
	tokenDom := token.InitTokenDom()
	nodeDom := node.InitNodeDom()
	statusDom := status.InitStatusDom()
	return &Domain{
		User:   userDom,
		Token:  tokenDom,
		Node:   nodeDom,
		Status: statusDom,
	}
}
