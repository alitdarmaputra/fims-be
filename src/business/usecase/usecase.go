package usecase

import (
	"github.com/alitdarmaputra/fims-be/src/business/domain"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/history"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/node"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/smtp"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/status"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/user"
	"github.com/alitdarmaputra/fims-be/src/config"
	"gorm.io/gorm"
)

type Usecase struct {
	User    user.UserUsecase
	Node    node.NodeUsecase
	History history.HistoryUsecase
	Status  status.StatusUsecase
}

func Init(
	dom *domain.Domain,
	db *gorm.DB,
	cfg *config.Api,
) *Usecase {
	smtpUsecase := smtp.InitSMTPUsecase(cfg.SMTP)
	userUsecase := user.InitUserUsecase(dom.User, smtpUsecase, dom.Token, db, cfg)
	nodeUsecase := node.InitNodeUsecase(
		db,
		cfg,
		dom.Node,
		dom.Status,
		dom.User,
		dom.Figma,
		dom.History,
		smtpUsecase,
	)
	historyUsecase := history.InitHistoryUsecase(dom.History, db)
	statusUsecase := status.InitStatusUsecase(db, cfg, dom.Status)
	return &Usecase{
		User:    userUsecase,
		Node:    nodeUsecase,
		History: historyUsecase,
		Status:  statusUsecase,
	}
}
