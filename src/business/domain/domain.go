package domain

import (
	"github.com/alitdarmaputra/fims-be/src/business/domain/figma"
	"github.com/alitdarmaputra/fims-be/src/business/domain/history"
	"github.com/alitdarmaputra/fims-be/src/business/domain/node"
	"github.com/alitdarmaputra/fims-be/src/business/domain/status"
	"github.com/alitdarmaputra/fims-be/src/business/domain/token"
	"github.com/alitdarmaputra/fims-be/src/business/domain/user"
	"github.com/alitdarmaputra/fims-be/src/config"
)

type Domain struct {
	User    user.UserDom
	Token   token.TokenDom
	Node    node.NodeDom
	Status  status.StatusDom
	Figma   figma.FigmaDom
	History history.HistoryDom
}

func Init(cfg *config.Api) *Domain {
	userDom := user.InitUserDom()
	tokenDom := token.InitTokenDom()
	nodeDom := node.InitNodeDom()
	statusDom := status.InitStatusDom()
	figmaDom := figma.InitFigmaDom(cfg)
	historyDom := history.InitHistoryDom()

	return &Domain{
		User:    userDom,
		Token:   tokenDom,
		Node:    nodeDom,
		Status:  statusDom,
		Figma:   figmaDom,
		History: historyDom,
	}
}
