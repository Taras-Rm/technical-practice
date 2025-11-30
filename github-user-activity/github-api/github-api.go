package githubapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IGithubAPI interface {
	GetUserActivity(username string) ([]PublicEvent, error)
}

type GithubAPI struct {
	ActivityURI string
}

func NewGithubAPI() IGithubAPI {
	return &GithubAPI{
		ActivityURI: "https://api.github.com/users/<username>/events",
	}
}

func (api *GithubAPI) GetUserActivity(username string) ([]PublicEvent, error) {
	uri := api.prepareActivityURI(username)

	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("username not found")
		}
		return nil, fmt.Errorf("failed request. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	publicEvents, err := api.parseGetUserActivityResponse(body)
	if err != nil {
		return nil, err
	}

	return publicEvents, nil
}

func (api *GithubAPI) prepareActivityURI(username string) string {
	return strings.Replace(api.ActivityURI, "<username>", username, 1)
}

func (api *GithubAPI) parseGetUserActivityResponse(data []byte) ([]PublicEvent, error) {
	var events []PublicEvent

	err := json.Unmarshal(data, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
