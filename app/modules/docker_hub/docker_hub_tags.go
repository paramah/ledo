package docker_hub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DockerImageTag struct {
	Layer string `json:"layer"`
	Name  string `json:"name"`
}

var urlTags = "https://registry.hub.docker.com/v1/repositories"

func GetImageTags(image string) []DockerImageTag {
	urlTags = urlTags + "/" + image + "/tags"
	res, _ := http.Get(urlTags)
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var tags []DockerImageTag
	json.Unmarshal(bytes, &tags)

	return tags
}

func ImageTagsToArray(tags []DockerImageTag) []string {
	n := len(tags)
	arrTags := make([]string, n)
	for idx, tag := range tags {
		arrTags[idx] = tag.Name
	}

	return arrTags
}
