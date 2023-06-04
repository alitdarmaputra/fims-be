package figma

type FigmaDom interface {
	GetFileNodes(fileKey string, nodeId string) (HTTPFigmaFileResponse, error)
	GetImage(fileKey string, nodeId string) (string, error)
}
