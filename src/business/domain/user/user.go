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
	Update(c context.Context, tx *gorm.DB, user model.User) (model.User, error)
	FindUnverifiedById(c context.Context, tx *gorm.DB, userId uint) (model.User, error)
	FindAll(c context.Context, tx *gorm.DB, offset, limit int, search string) ([]model.User, int)
}
