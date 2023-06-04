package figma

import (
	"github.com/alitdarmaputra/fims-be/src/config"
)

type FigmaDomImpl struct {
	cfg *config.Api
}

func InitFigmaDom(cfg *config.Api) FigmaDom {
	return &FigmaDomImpl{
		cfg: cfg,
	}
}

func (domain *FigmaDomImpl) GetFileNodes(
	fileKey string,
	nodeId string,
) (HTTPFigmaFileResponse, error) {
	return domain.httpGetFileNodes(fileKey, nodeId)
}

func (domain *FigmaDomImpl) GetImage(
	fileKey string,
	nodeId string,
) (string, error) {
	return domain.httpGetImage(fileKey, nodeId)
}
