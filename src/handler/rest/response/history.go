package response

import (
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type HTTPHistoryResponse struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedBy    string    `json:"updated_by"`
	Description  string    `json:"description"`
	HistoryType  string    `json:"history_type"`
	FigmaUrl     string    `json:"figma_url"`
	SnapshotPath string    `json:"snapshot_path"`
	StatusFrom   string    `json:"status_from"`
	StatusTo     string    `json:"status_to"`
	NodeTitle    string    `json:"node_title"`
}

func ToHistoryResponse(node model.History) HTTPHistoryResponse {
	response := HTTPHistoryResponse{}

	response.CreatedAt = node.CreatedAt
	response.UpdatedBy = node.User.Name
	response.Description = node.Description
	response.HistoryType = node.HistoryType
	response.NodeTitle = node.Node.Title
	response.UpdatedBy = node.User.Name

	if node.FigmaUrl.Valid {
		response.FigmaUrl = node.FigmaUrl.String
	}

	if node.StatusFromId.Valid {
		response.StatusFrom = node.StatusFrom.Name
	}

	if node.StatusToId.Valid {
		response.StatusTo = node.StatusTo.Name
	}

	if node.SnapshotPath.Valid {
		response.SnapshotPath = node.SnapshotPath.String
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
