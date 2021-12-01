package docker_hub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DockerSearchResponse struct {
	NumPages   int           `json:"num_pages"`
	NumResults int           `json:"num_results"`
	PageSize   int           `json:"page_size"`
	Page       int           `json:"page"`
	Query      string        `json:"query"`
	Results    []DockerImage `json:"results"`
}

type DockerImage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StarCount   int    `json:"star_count"`
	IsTrusted   bool   `json:"is_trusted"`
	IsAutomated bool   `json:"is_automated"`
	IsOfficial  bool   `json:"is_official"`
}

var urlSearch = "https://registry.hub.docker.com/v1/search"

func GetImage(image string) []DockerImage {
	var resp DockerSearchResponse
	searchString := url.QueryEscape(image)
	urlSearch = urlSearch + "/?q=" + searchString
	fmt.Printf("%v", urlSearch)
	res, _ := http.Get(urlSearch)
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(bytes, &resp)

	return resp.Results
}

func DockerImageToArray(images []DockerImage) []string {
	n := len(images)
	arrTags := make([]string, n)
	for idx, image := range images {
		arrTags[idx] = image.Name
	}

	return arrTags
}
