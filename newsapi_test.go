package newsapi

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WithHTTPClient(t *testing.T) {
	c := &Client{}
	client := &http.Client{}
	WithHTTPClient(client)(c)

	assert.Equal(t, client, c.client)
}

func Test_WithURL(t *testing.T) {
	c := &Client{}
	WithBaseURL("123")(c)

	assert.Equal(t, "123", c.baseURL)
}

func Test_New(t *testing.T) {
	client := &http.Client{}
	c := NewClient("321", WithBaseURL("123"), WithHTTPClient(client))

	assert.Equal(t, "321", c.apiKey)
	assert.Equal(t, "123", c.baseURL)
	assert.Equal(t, client, c.client)
}

func Test_Client_Everything(t *testing.T) {
	tstamp := time.Date(2022, 02, 22, 22, 22, 22, 0, time.UTC)

	tests := map[string]struct {
		Param     EverythingParams
		Resp      httpmock.Responder
		Articles  []Article
		PageCount uint
		Err       error
	}{
		"Invalid parameters": {
			Param: EverythingParams{},
			Resp:  httpmock.NewStringResponder(http.StatusOK, ""),
			Err:   ErrParamsScopeTooBroad,
		},
		"Newsapi returned invalid JSON": {
			Param: EverythingParams{
				Sources: []string{
					"testy",
				},
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "testy", req.URL.Query().Get("sources"))
				return httpmock.NewStringResponse(
					http.StatusOK,
					`{
					`,
				), nil
			},
			Err: assert.AnError,
		},
		"Newsapi returned an error": {
			Param: EverythingParams{
				QueryInTitle: "321",
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "321", req.URL.Query().Get("qInTitle"))
				return httpmock.NewStringResponse(
					http.StatusBadRequest,
					`{
						"status":"error",
						"code": "100",
						"message": "bad thing"
					}`,
				), nil
			},
			Err: &Error{
				HTTPCode: http.StatusBadRequest,
				APICode:  "100",
				Message:  "bad thing",
			},
		},
		"Succesfully fetched articles": {
			Param: EverythingParams{
				Query: "321",
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "321", req.URL.Query().Get("q"))
				return httpmock.NewStringResponse(
					http.StatusOK,
					`{
						"status":"ok",
						"totalResults":10,
						"articles":[
							{
								"source": {
									"id": "test",
									"name": "testtest"
								},
								"author": "david",
								"title": "btc",
								"description": "short description",
								"url": "test.com/david",
								"urlToImage": "test.com/image",
								"publishedAt": "2022-02-22T22:22:22Z",
								"content":"123 123 123"
							},
							{
								"source": {
									"id": "test2",
									"name": "testtest2"
								},
								"author": "vavid",
								"title": "eth",
								"description": "long description",
								"url": "test.com/vavid",
								"urlToImage": "test.com/image2",
								"publishedAt": "2022-02-22T22:22:22Z",
								"content":"123 423 143"
							}
						]
					}`,
				), nil
			},
			Articles: []Article{
				{
					Source: SourceID{
						ID:   "test",
						Name: "testtest",
					},
					Author:      "david",
					Title:       "btc",
					Description: "short description",
					URL:         "test.com/david",
					URLToImage:  "test.com/image",
					PublishedAt: tstamp,
					Content:     "123 123 123",
				},
				{
					Source: SourceID{
						ID:   "test2",
						Name: "testtest2",
					},
					Author:      "vavid",
					Title:       "eth",
					Description: "long description",
					URL:         "test.com/vavid",
					URLToImage:  "test.com/image2",
					PublishedAt: tstamp,
					Content:     "123 423 143",
				},
			},
			PageCount: 10,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			transport := httpmock.NewMockTransport()
			client := &Client{
				client: &http.Client{
					Transport: transport,
				},
				baseURL: "test/",
			}

			transport.RegisterResponder(http.MethodGet, "test/everything", test.Resp)

			articles, pageCount, err := client.Everything(context.Background(), test.Param)

			if errors.Is(test.Err, assert.AnError) {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.Err, err)
			}

			if err != nil {
				return
			}

			assert.Equal(t, test.Articles, articles)
			assert.Equal(t, test.PageCount, pageCount)
		})
	}
}

func Test_Client_TopHeadlines(t *testing.T) {
	tstamp := time.Date(2022, 02, 22, 22, 22, 22, 0, time.UTC)

	tests := map[string]struct {
		Param     TopHeadlinesParams
		Resp      httpmock.Responder
		Articles  []Article
		PageCount uint
		Err       error
	}{
		"Invalid parameters": {
			Param: TopHeadlinesParams{},
			Resp:  httpmock.NewStringResponder(http.StatusOK, ""),
			Err:   ErrParamsScopeTooBroad,
		},
		"Newsapi returned invalid JSON": {
			Param: TopHeadlinesParams{
				Sources: []string{
					"testy",
				},
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "testy", req.URL.Query().Get("sources"))
				return httpmock.NewStringResponse(
					http.StatusOK,
					`{
					`,
				), nil
			},
			Err: assert.AnError,
		},
		"Newsapi returned an error": {
			Param: TopHeadlinesParams{
				Query: "321",
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "321", req.URL.Query().Get("q"))
				return httpmock.NewStringResponse(
					http.StatusBadRequest,
					`{
						"status":"error",
						"code": "100",
						"message": "bad thing"
					}`,
				), nil
			},
			Err: &Error{
				HTTPCode: http.StatusBadRequest,
				APICode:  "100",
				Message:  "bad thing",
			},
		},
		"Succesfully fetched articles": {
			Param: TopHeadlinesParams{
				Query: "321",
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "321", req.URL.Query().Get("q"))
				return httpmock.NewStringResponse(
					http.StatusOK,
					`{
						"status":"ok",
						"totalResults":10,
						"articles":[
							{
								"source": {
									"id": "test",
									"name": "testtest"
								},
								"author": "david",
								"title": "btc",
								"description": "short description",
								"url": "test.com/david",
								"urlToImage": "test.com/image",
								"publishedAt": "2022-02-22T22:22:22Z",
								"content":"123 123 123"
							},
							{
								"source": {
									"id": "test2",
									"name": "testtest2"
								},
								"author": "vavid",
								"title": "eth",
								"description": "long description",
								"url": "test.com/vavid",
								"urlToImage": "test.com/image2",
								"publishedAt": "2022-02-22T22:22:22Z",
								"content":"123 423 143"
							}
						]
					}`,
				), nil
			},
			Articles: []Article{
				{
					Source: SourceID{
						ID:   "test",
						Name: "testtest",
					},
					Author:      "david",
					Title:       "btc",
					Description: "short description",
					URL:         "test.com/david",
					URLToImage:  "test.com/image",
					PublishedAt: tstamp,
					Content:     "123 123 123",
				},
				{
					Source: SourceID{
						ID:   "test2",
						Name: "testtest2",
					},
					Author:      "vavid",
					Title:       "eth",
					Description: "long description",
					URL:         "test.com/vavid",
					URLToImage:  "test.com/image2",
					PublishedAt: tstamp,
					Content:     "123 423 143",
				},
			},
			PageCount: 10,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			transport := httpmock.NewMockTransport()
			client := &Client{
				client: &http.Client{
					Transport: transport,
				},
				baseURL: "test/",
			}

			transport.RegisterResponder(http.MethodGet, "test/top-headlines", test.Resp)

			articles, pageCount, err := client.TopHeadlines(context.Background(), test.Param)

			if errors.Is(test.Err, assert.AnError) {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.Err, err)
			}

			if err != nil {
				return
			}

			assert.Equal(t, test.Articles, articles)
			assert.Equal(t, test.PageCount, pageCount)
		})
	}
}

func Test_Client_Sources(t *testing.T) {
	tests := map[string]struct {
		Param   SourceParams
		Resp    httpmock.Responder
		Sources []Source
		Err     error
	}{
		"Invalid parameters": {
			Param: SourceParams{
				Categories: []Category{
					"test",
				},
			},
			Resp: httpmock.NewStringResponder(http.StatusOK, ""),
			Err:  ErrInvalidCategory,
		},
		"Newsapi returned invalid JSON": {
			Param: SourceParams{
				Categories: []Category{
					CategoryBusiness,
				},
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "business", req.URL.Query().Get("category"))
				return httpmock.NewStringResponse(
					http.StatusOK,
					`{
					`,
				), nil
			},
			Err: assert.AnError,
		},
		"Newsapi returned an error": {
			Param: SourceParams{
				Categories: []Category{
					CategoryBusiness,
				},
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "business", req.URL.Query().Get("category"))
				return httpmock.NewStringResponse(
					http.StatusBadRequest,
					`{
						"status":"error",
						"code": "100",
						"message": "bad thing"
					}`,
				), nil
			},
			Err: &Error{
				HTTPCode: http.StatusBadRequest,
				APICode:  "100",
				Message:  "bad thing",
			},
		},
		"Succesfully fetched sources": {
			Param: SourceParams{
				Languages: []Language{
					LanguageChinese,
				},
			},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "zh", req.URL.Query().Get("language"))
				return httpmock.NewStringResponse(
					http.StatusOK,
					`{
						"status":"ok",
						"sources":[
							{
								"id": "test",
								"name": "testtest",
								"description": "short description",
								"url": "test.com/david",
								"category": "business",
								"language": "ru",
								"country": "pl"
							},
							{
								"id": "test2",
								"name": "testtest2",
								"description": "long description",
								"url": "test.com/vavid",
								"category": "general",
								"language": "de",
								"country": "nz"
							}
						]
					}`,
				), nil
			},
			Sources: []Source{
				{
					SourceID: SourceID{
						ID:   "test",
						Name: "testtest",
					},
					Description: "short description",
					URL:         "test.com/david",
					Category:    CategoryBusiness,
					Language:    LanguageRussian,
					Country:     CountryPoland,
				},
				{
					SourceID: SourceID{
						ID:   "test2",
						Name: "testtest2",
					},
					Description: "long description",
					URL:         "test.com/vavid",
					Category:    CategoryGeneral,
					Language:    LanguageGerman,
					Country:     CountryNewZealand,
				},
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			transport := httpmock.NewMockTransport()
			client := &Client{
				client: &http.Client{
					Transport: transport,
				},
				baseURL: "test/",
			}

			transport.RegisterResponder(http.MethodGet, "test/top-headlines/sources", test.Resp)

			sources, err := client.Sources(context.Background(), test.Param)

			if errors.Is(test.Err, assert.AnError) {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.Err, err)
			}

			if err != nil {
				return
			}

			assert.Equal(t, test.Sources, sources)
		})
	}
}

func Test_Client_get(t *testing.T) {
	tests := map[string]struct {
		StatusCode int
		Body       []byte
		Params     params
		Resp       httpmock.Responder
		NilContext bool
		Err        error
	}{
		"Validate returns an error": {
			Params: &EverythingParams{},
			Resp:   httpmock.NewBytesResponder(http.StatusBadRequest, []byte{1, 2, 3, 4}),
			Err:    ErrParamsScopeTooBroad,
		},
		"Invalid context": {
			Params:     &SourceParams{},
			NilContext: true,
			Err:        assert.AnError,
		},
		"Client do returns an error": {
			Params: &SourceParams{},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "777", req.Header.Get("X-Api-Key"))
				return nil, assert.AnError
			},
			Err: assert.AnError,
		},
		"Successful request": {
			StatusCode: http.StatusBadRequest,
			Body:       []byte{1, 2, 3, 4},
			Params:     &SourceParams{},
			Resp: func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "777", req.Header.Get("X-Api-Key"))
				return httpmock.NewBytesResponse(
					http.StatusBadRequest,
					[]byte{1, 2, 3, 4},
				), nil
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			transport := httpmock.NewMockTransport()
			client := &Client{
				client: &http.Client{
					Transport: transport,
				},
				baseURL: "test/",
				apiKey:  "777",
			}

			transport.RegisterResponder(http.MethodGet, "test/123", test.Resp)

			var ctx context.Context
			if !test.NilContext {
				ctx = context.Background()
			}

			statusCode, body, err := client.get(
				ctx,
				"123",
				test.Params,
			)

			if body != nil {
				defer body.Close()
			}

			if errors.Is(test.Err, assert.AnError) {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.Err, err)
			}

			if err != nil {
				return
			}

			buf := &bytes.Buffer{}
			_, err = io.Copy(buf, body)
			require.NoError(t, err)

			assert.Equal(t, test.StatusCode, statusCode)
			assert.Equal(t, test.Body, buf.Bytes())
		})
	}
}
