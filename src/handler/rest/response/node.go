package response

import (
	"fmt"
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type HTTPNodeResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	FigmaUrl  string    `json:"figma_url"`
	CreatedBy string    `json:"created_by"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type HTTPNodeDetailResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	FigmaUrl    string    `json:"figma_url"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	AssigneeId  uint      `json:"assignee_id"`
}

func ToNodeResponse(node model.Node, figmaBaseUrl string) HTTPNodeResponse {
	return HTTPNodeResponse{
		ID:        node.ID,
		Title:     node.Title,
		FigmaUrl:  fmt.Sprintf("%s/file/%s?node-id=%s", figmaBaseUrl, node.FigmaKey, node.NodeId),
		CreatedBy: node.User.Name,
		Status:    node.Status.Name,
		CreatedAt: node.CreatedAt,
	}
}

func ToNodeDetailResponse(node model.Node, figmaBaseUrl string) HTTPNodeDetailResponse {
	return HTTPNodeDetailResponse{
		ID:          node.ID,
		Title:       node.Title,
		FigmaUrl:    fmt.Sprintf("%s/file/%s?node-id=%s", figmaBaseUrl, node.FigmaKey, node.NodeId),
		Description: node.Description,
		CreatedBy:   node.User.Name,
		Status:      node.Status.Name,
		CreatedAt:   node.CreatedAt,
		AssigneeId:  node.AssigneeId,
	}
}

func ToNodeResponses(nodes []model.Node, figmaBaseUrl string) []HTTPNodeResponse {
	var nodeResponses []HTTPNodeResponse = []HTTPNodeResponse{}
	for _, node := range nodes {
		nodeResponses = append(nodeResponses, ToNodeResponse(node, figmaBaseUrl))
	}
	return nodeResponses
}
