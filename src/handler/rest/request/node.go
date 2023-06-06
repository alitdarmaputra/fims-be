package request

type HTTPNodeCreateUpdateRequest struct {
	Title       string `json:"title"       binding:"required"`
	FigmaKey    string `json:"figma_key"   binding:"required"`
	NodeId      string `json:"node_id"     binding:"required"`
	Description string `json:"description"`
}

type HTTPNodeUpdateStatusRequest struct {
	StatusId uint `json:"status_id"`
}