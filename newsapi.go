package newsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ClientOption is used to set client configuration options.
type ClientOption func(c *Client)

// WithHTTPClient allows to set custom http client when newsapi client is
// making requests.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		c.client = hc
	}
}

// WithURL allows to set custom url that http calls are made to.
func WithURL(url string) ClientOption {
	return func(c *Client) {
		c.url = url
	}
}

// Client implements newsapi endpoints and allows to fetch articles and sources.
type Client struct {
	apiKey string
	url    string
	client *http.Client
}

// New creates fresh instance of newsapi client.
func New(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		url: "https://newsapi.org/v2/",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Everything retrieves articles based on provided parameters.
// Endpoint documentation can be found here: https://newsapi.org/docs/endpoints/everything
func (c *Client) Everything(ctx context.Context, pr EverythingParams) ([]Article, uint, error) {
	return c.getArticles(
		ctx,
		"everything",
		&pr,
	)
}

// TopHeadlines retrieves top headlines articles based on provided parameters.
// Endpoint documentation can be found here: https://newsapi.org/docs/endpoints/top-headlines
func (c *Client) TopHeadlines(ctx context.Context, pr TopHeadlinesParams) ([]Article, uint, error) {
	return c.getArticles(
		ctx,
		"top-headlines",
		&pr,
	)
}

// Sources retrieves available sources for top headlines and everything
// parameters.
// Endpoint documentation can be found here: https://newsapi.org/docs/endpoints/sources
func (c *Client) Sources(ctx context.Context, pr SourceParams) ([]Source, error) {
	statusCode, body, err := c.get(
		ctx,
		"top-headlines/sources",
		&pr,
	)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data := struct {
		Status  string   `json:"status"`
		Code    string   `json:"code"`
		Sources []Source `json:"sources"`
		Message string   `json:"message"`
	}{}

	if err = json.NewDecoder(body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Status != "ok" {
		return nil, &Error{
			StatusCode: statusCode,
			Code:       data.Code,
			Message:    data.Message,
		}
	}

	return data.Sources, nil
}

func (c *Client) getArticles(
	ctx context.Context,
	path string,
	pr params,
) ([]Article, uint, error) {

	statusCode, body, err := c.get(
		ctx,
		path,
		pr,
	)
	if err != nil {
		return nil, 0, err
	}
	defer body.Close()

	data := struct {
		Status       string    `json:"status"`
		Code         string    `json:"code"`
		TotalResults uint      `json:"totalResults"`
		Articles     []Article `json:"articles"`
		Message      string    `json:"message"`
	}{}

	if err = json.NewDecoder(body).Decode(&data); err != nil {
		return nil, 0, err
	}

	if data.Status != "ok" {
		return nil, 0, &Error{
			StatusCode: statusCode,
			Code:       data.Code,
			Message:    data.Message,
		}
	}

	return data.Articles, data.TotalResults, nil
}

func (c *Client) get(
	ctx context.Context,
	path string,
	pr params,
) (int, io.ReadCloser, error) {

	if err := pr.Validate(); err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s?%s", c.url, path, pr.rawQuery()),
		http.NoBody,
	)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("X-Api-Key", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, resp.Body, nil
}

type params interface {
	Validate() error
	rawQuery() string
}
