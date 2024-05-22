package clusters

import (
	"errors"
	"github.com/opensaucerer/goaxios"
	"net/http"
)

func (cl *ClustersSDK) makeRequest(method string, target interface{}) {
	cl.err = nil

	req := goaxios.GoAxios{
		Url:    cl.URL,
		Method: method,
		// Headers: map[string]string{
		// 	"X-API-KEY": cl.APIKey,
		// },
	}

	if target != nil {
		req.ResponseStruct = target
	}

	if method == http.MethodPost {
		req.Body = cl.body
	}

	res := req.RunRest()

	if res.Error != nil {
		cl.err = errors.Join(res.Error, UnrecoverableError)
		return
	}

	switch res.Response.StatusCode {
	case 404:
		cl.err = ErrorNotFound
	case 500:
		cl.err = ClustersInternalError
	default:
		if res.Response.StatusCode >= 300 {
			cl.err = UnrecoverableError
		} else if string(res.Bytes) == "null" {
			cl.err = ErrorNotFound
		}
	}
}
