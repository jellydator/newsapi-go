# newsapi
[![GoDoc](https://godoc.org/github.com/jellydator/newsapi-go?status.png)](https://godoc.org/github.com/jellydator/newsapi-go)
[![Test coverage](http://gocover.io/_badge/github.com/jellydator/newsapi-go)](https://gocover.io/github.com/jellydator/newsapi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jellydator/newsapi-go)](https://goreportcard.com/report/github.com/jellydator/newsapi-go)

Go client implementation for the NewsAPI.

## Installation
```
go get github.com/jellydator/newsapi-go
```

## Usage
Simply, we create a client using `NewClient` function, by default only an API
key is required, however some other parameters can be set using variadic
option functions.
```go
client := newsapi.NewClient("apiKey", newsapi.WithHTTPClient(&http.Client{
	Timeout: 5 * time.Second,
}))
```

## Endpoints

### Everything
`Everything` retrieves all articles based on provided parameters.
Full endpoint documentation can be viewed [here](https://newsapi.org/docs/endpoints/everything).
```go
articles, pageCount, err := client.Everything(context.Background(), newsapi.EverythingParams{
	Query: "cryptocurrency",
})
if err != nil {
	// handle error
}
// success
```

### Top Headlines
`TopHeadlines` retrieves top headlines articles based on provided parameters.
Full endpoint documentation can be viewed [here](https://newsapi.org/docs/endpoints/top-headlines).
```go
articles, pageCount, err := client.TopHeadlines(context.Background(), newsapi.TopHeadlinesParams{
	Query: "cryptocurrency",
})
if err != nil {
	// handle error
}
// success
```

### Sources
`Sources` retrieves available sources based on provided parameters.
Full endpoint documentation can be viewed [here](https://newsapi.org/docs/endpoints/sources).
```go
sources, err := client.Sources(context.Background(), newsapi.SourceParams{
	Categories: []newsapi.Category{
		newsapi.CategoryBusiness,
		newsapi.CategoryScience,
	},
})
if err != nil {
	// handle error
}
// success
```
