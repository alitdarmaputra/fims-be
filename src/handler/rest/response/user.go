package response

import (
	"database/sql"
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type HTTPUserDetailResponse struct {
	ID         uint         `json:"id"`
	Email      string       `json:"email"`
	Name       string       `json:"name"`
	ProfileImg string       `json:"profile_img"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}

func ToUserResponse(user model.User) HTTPUserDetailResponse {
	return HTTPUserDetailResponse{
		ID:         user.ID,
		Name:       user.Name,
		ProfileImg: user.ProfileImg,
	}
}

func ToUserResponses(users []model.User) []HTTPUserDetailResponse {
	var usersResponses []HTTPUserDetailResponse = []HTTPUserDetailResponse{}
	for _, user := range users {
		usersResponses = append(usersResponses, ToUserResponse(user))
	}
	return usersResponses
}

type TotalUserResponse struct {
	Total int64 `json:"total"`
}
