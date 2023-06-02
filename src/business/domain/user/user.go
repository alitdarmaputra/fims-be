package user

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type UserDom interface {
	Create(c context.Context, tx *gorm.DB, user model.User) (model.User, error)
	Save(c context.Context, tx *gorm.DB, user model.User) (model.User, error)
	FindOne(c context.Context, tx *gorm.DB, email string) (model.User, error)
	FindById(c context.Context, tx *gorm.DB, user_id uint) (model.User, error)
	Update(ctx context.Context, tx *gorm.DB, user model.User) (model.User, error)
	FindUnverifiedById(ctx context.Context, tx *gorm.DB, userId uint) (model.User, error)
}
