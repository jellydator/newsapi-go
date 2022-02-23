package newsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const _defaultBaseURL = "https://newsapi.org/v2/"

// Client handles request sending to newsapi.
type Client struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// ClientOption is used to set client configuration options.
type ClientOption func(c *Client)

// WithHTTPClient sets custom http client.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		c.client = hc
	}
}

// WithBaseURL sets custom base url.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// New creates a fresh instance of newsapi client.
func New(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: _defaultBaseURL,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Everything retrieves articles by the provided parameters.
// The uint return value indicates the number of available articles. The
// length of the returned slice may be less than this value; additional calls
// need to be make to retrieve other available articles.
// Endpoint documentation can be found here:
// https://newsapi.org/docs/endpoints/everything
func (c *Client) Everything(ctx context.Context, pr EverythingParams) ([]Article, uint, error) {
	return c.getArticles(ctx, "everything", &pr)
}

// TopHeadlines retrieves top headlines articles by the provided parameters.
// The uint return value indicates the number of available articles. The
// length of the returned slice may be less than this value; additional calls
// need to be make to retrieve other available articles.
// Endpoint documentation can be found here:
// https://newsapi.org/docs/endpoints/top-headlines
func (c *Client) TopHeadlines(ctx context.Context, pr TopHeadlinesParams) ([]Article, uint, error) {
	return c.getArticles(ctx, "top-headlines", &pr)
}

// Sources retrieves available sources for top headlines and everything
// endpoints by the provided parameters.
// Endpoint documentation can be found here:
// https://newsapi.org/docs/endpoints/sources
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
			HTTPCode: statusCode,
			APICode:  data.Code,
			Message:  data.Message,
		}
	}

	return data.Sources, nil
}

// getArticles retrieves articles by the provided path and parameters.
// The uint return value indicates the number of available articles. The
// length of the returned slice may be less than this value; additional calls
// need to be make to retrieve other available articles.
func (c *Client) getArticles(ctx context.Context, endpoint string, pr params) ([]Article, uint, error) {
	statusCode, body, err := c.get(
		ctx,
		endpoint,
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
			HTTPCode: statusCode,
			APICode:  data.Code,
			Message:  data.Message,
		}
	}

	return data.Articles, data.TotalResults, nil
}

// get sends a GET request to the provided endpoint.
func (c *Client) get(ctx context.Context, endpoint string, pr params) (int, io.ReadCloser, error) {
	if err := pr.validate(); err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s?%s", c.baseURL, endpoint, pr.rawQuery()),
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

// params is an interface is used to process query parameters.
type params interface {
	// validate should validate the params.
	validate() error

	// rawQuery should build a raw query from the params.
	rawQuery() string
}
