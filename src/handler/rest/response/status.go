package response

import (
	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type HTTPStatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToStatusResponse(status model.Status) HTTPStatusResponse {
	return HTTPStatusResponse{
		ID:   status.ID,
		Name: status.Name,
	}
}

func ToStatusResponses(statuses []model.Status) []HTTPStatusResponse {
	statusResponses := []HTTPStatusResponse{}

	for _, status := range statuses {
		statusResponses = append(statusResponses, ToStatusResponse(status))
	}
	return statusResponses
}
