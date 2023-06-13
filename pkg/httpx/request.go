package httpx

import (
	"context"
	"net/http"
	"net/url"
)

type RequestOptions struct {
	Ctx             context.Context
	BodyBinder      interface{}
	ResponseMiddler func(response *http.Response) error
	HttpMethod      string
	DestinationURL  string
	Queries         url.Values
	Headers         http.Header
	RequestBody     interface{}
}

func convArray(queries map[string][]string) []interface{} {
	result := make([]interface{}, 0, len(queries)*2) // todo: queries headers bad implements
	for h, v := range queries {
		result = append(result, h, v)
	}
	return result
}
