package domain

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"gorm.io/gorm"
)

type User interface {
	Create(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	FindOne(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
}
