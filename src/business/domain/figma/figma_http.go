package figma

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/entity"
)

type HTTPFigmaFileResponse struct {
	Name         string                 `json:"name"`
	LastModified time.Time              `json:"lastModified"`
	Version      string                 `json:"version"`
	Nodes        map[string]interface{} `json:"nodes"`
}

type HTTPFigmaImageResponse struct {
	Images map[string]interface{} `json:"images"`
}

func (domain *FigmaDomImpl) httpGetFileNodes(
	fileKey string,
	nodeId string,
) (HTTPFigmaFileResponse, error) {
	figmaFileResponse := HTTPFigmaFileResponse{}

	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodGet,
		domain.cfg.Figma.FigmaApiBaseUrl+"/v1/files/"+fileKey+"/nodes",
		nil,
	)

	if err != nil {
		return figmaFileResponse, err
	}

	q := req.URL.Query()
	q.Add("ids", nodeId)

	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-FIGMA-TOKEN", domain.cfg.Figma.FigmaToken)

	response, err := client.Do(req)
	if err != nil {
		return figmaFileResponse, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusBadRequest {
		return figmaFileResponse, entity.NewBadRequestError("Invalid payload")
	}

	if response.StatusCode == http.StatusNotFound {
		return figmaFileResponse, entity.NewNotFoundError("Node not found")
	}

	if response.StatusCode == http.StatusInternalServerError {
		return figmaFileResponse, entity.NewBadGateWayError("Gateway error")
	}

	err = json.NewDecoder(response.Body).Decode(&figmaFileResponse)
	if err != nil {
		return figmaFileResponse, err
	}

	if figmaFileResponse.Nodes[strings.ReplaceAll(nodeId, "-", ":")] == nil {
		return figmaFileResponse, entity.NewNotFoundError("Node not foudn")
	}

	return figmaFileResponse, nil
}

func (domain *FigmaDomImpl) httpGetImage(
	fileKey string,
	nodeId string,
) (string, error) {
	figmaImageResponse := HTTPFigmaImageResponse{}

	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodGet,
		domain.cfg.Figma.FigmaApiBaseUrl+"/v1/images/"+fileKey,
		nil,
	)

	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("ids", nodeId)

	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-FIGMA-TOKEN", domain.cfg.Figma.FigmaToken)

	response, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusBadRequest {
		return "", entity.NewBadRequestError("Invalid payload")
	}

	if response.StatusCode == http.StatusNotFound {
		return "", entity.NewNotFoundError("Node not found")
	}

	if response.StatusCode == http.StatusInternalServerError {
		return "", entity.NewBadGateWayError("Gateway error")
	}

	err = json.NewDecoder(response.Body).Decode(&figmaImageResponse)
	if err != nil {
		return "", err
	}

	newNodeId := strings.ReplaceAll(nodeId, "-", ":")

	if figmaImageResponse.Images[newNodeId] == nil {
		return "", entity.NewNotFoundError("Node not found")
	}

	figmaImageUrl := figmaImageResponse.Images[newNodeId].(string)

	// save image to local
	fileUrl, err := url.Parse(figmaImageUrl)
	if err != nil {
		return "", err
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// create directory
	staticPath := filepath.Join("public", "upload")
	err = os.MkdirAll(staticPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// create file
	fileFigmaPath := filepath.Join("public", "upload", fmt.Sprintf("%s.png", fileName))
	file, err := os.Create(fileFigmaPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// get file data
	imageRes, err := client.Get(figmaImageUrl)
	if err != nil {
		return "", err
	}
	defer imageRes.Body.Close()

	// Write content
	_, err = io.Copy(file, imageRes.Body)
	if err != nil {
		return "", err
	}

	return fileFigmaPath, nil
}
