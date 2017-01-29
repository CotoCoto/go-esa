package esa

import (
	"encoding/json"
	"fmt"
	"time"
)

type Stargazer struct {
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	User      struct {
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Icon       string `json:"icon"`
	} `json:"user"`
}

type GetTeamPostStargazersRequest struct {
	PaginationRequest
}

type GetTeamPostStargazersResponse struct {
	Stargazers []Stargazer `json:"stargazers"`
	PaginationResponse
}

func (c *Client) GetTeamPostStarGazers(teamName string, number int, req *GetTeamPostStargazersRequest) (*GetTeamPostStargazersResponse, error) {
	r := c.get(fmt.Sprintf("/v1/teams/%s/posts/%d/stargazers", teamName, number))

	resp, body, errs := c.setPaginationParams(r, &req.PaginationRequest).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamPostStargazersResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type CreateTeamPostStartRequest struct {
	Body *string `json:"body"`
}

func (c *Client) CreateTeamPostStar(teamName string, postNumber int, req *CreateTeamPostStartRequest) error {
	resp, body, errs := c.post(fmt.Sprintf("/v1/teams/%s/posts/%d/star", teamName, postNumber)).Send(req).End()
	if len(errs) > 0 {
		return errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return c.parseError(body)
	}

	return nil
}
