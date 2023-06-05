package response

import (
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type HTTPHistoryResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   string    `json:"created_by"`
	Description string    `json:"description"`
	HistoryType string    `json:"history_type"`
	FigmaUrl    string    `json:"figma_url"`
	StatusFrom  string    `json:"status_from"`
	StatusTo    string    `json:"status_to"`
}

func ToHistoryResponse(node model.History) HTTPHistoryResponse {
	response := HTTPHistoryResponse{}

	response.CreatedAt = node.CreatedAt
	response.UpdatedBy = node.User.Name
	response.Description = node.Description
	response.HistoryType = node.HistoryType

	if node.FigmaUrl.Valid {
		response.FigmaUrl = node.FigmaUrl.String
	}

	if node.StatusFromId.Valid {
		response.StatusFrom = node.StatusFrom.Name
	}

	if node.StatusToId.Valid {
		response.StatusTo = node.StatusTo.Name
	}

	return response
}

func ToHistoryResponses(histories []model.History) []HTTPHistoryResponse {
	historyResponses := []HTTPHistoryResponse{}
	for _, history := range histories {
		historyResponses = append(historyResponses, ToHistoryResponse(history))
	}
	return historyResponses
}
