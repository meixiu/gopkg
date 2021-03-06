package httpclient

import (
	"net/url"

	"github.com/spf13/cast"

	"github.com/google/go-querystring/query"
)

// ParseUrlValues
func ParseUrlValues(data interface{}) (url.Values, error) {
	switch data.(type) {
	case url.Values:
		return data.(url.Values), nil
	case map[string]string:
		uv := make(url.Values)
		for k, v := range data.(map[string]string) {
			uv.Add(k, v)
		}
		return uv, nil
	case map[string]interface{}:
		uv := make(url.Values)
		for k, v := range data.(map[string]interface{}) {
			uv.Add(k, cast.ToString(v))
		}
		return uv, nil
	default:
		return query.Values(data)
	}
}
