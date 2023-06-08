package response

import (
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type HTTPHistoryResponse struct {
	NodeId           uint      `json:"node_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedBy        string    `json:"updated_by"`
	UpdatedByProfile string    `json:"updated_by_profile"`
	Description      string    `json:"description"`
	HistoryType      string    `json:"history_type"`
	FigmaUrl         string    `json:"figma_url"`
	SnapshotPath     string    `json:"snapshot_path"`
	StatusFrom       string    `json:"status_from"`
	StatusTo         string    `json:"status_to"`
	NodeTitle        string    `json:"node_title"`
}

func ToHistoryResponse(history model.History) HTTPHistoryResponse {
	response := HTTPHistoryResponse{}

	response.CreatedAt = history.CreatedAt
	response.UpdatedBy = history.User.Name
	response.UpdatedByProfile = history.User.ProfileImg
	response.Description = history.Description
	response.HistoryType = history.HistoryType
	response.NodeTitle = history.Node.Title
	response.UpdatedBy = history.User.Name
	response.NodeId = history.NodeId

	if history.FigmaUrl.Valid {
		response.FigmaUrl = history.FigmaUrl.String
	}

	if history.StatusFromId.Valid {
		response.StatusFrom = history.StatusFrom.Name
	}

	if history.StatusToId.Valid {
		response.StatusTo = history.StatusTo.Name
	}

	if history.SnapshotPath.Valid {
		response.SnapshotPath = history.SnapshotPath.String
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
