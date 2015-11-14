package httputil

import (
	"net/http"
	"strings"
	"net/url"
)

// contains all the datas to send to the API endpoint

type ApiRequest struct {
	Method string
	Headers http.Header
	Body interface{}
	Path string
	QueryString string
	FormValues url.Values
}

func (t *ApiRequest) SetHeaderParam(name string, values ...string) error {
	if t.Headers == nil {
		t.Headers = make(http.Header)
	}
	t.Headers.Set(name, values[0])
	return nil
}

// Set the path parameter,
// t.SetPathParam("name", "value") - /name/value
func (t *ApiRequest) SetPathParam(name string, value string) error {
	if len(t.Path) == 0 || !strings.HasSuffix(t.Path, "/") {
		t.Path += "/"
	}
	if len(name) == 0 {
		return nil
	}
	t.Path += name
	if len(value) == 0 {
		return nil
	}
	t.Path += "/"+value
	return nil
}

// Set the query parameter,
// t.SetQueryParam("name", "value") - {path}?name=value
func (t *ApiRequest) SetQueryParam(name string, values ...string) error {
	v := url.QueryEscape(values[0])
	if len(v)>0 {
		if len(t.QueryString) != 0 {
			t.QueryString += "&"
		}
		t.QueryString += name+"="+v
	}
	return nil
}

func (t *ApiRequest) SetFormParam(name string, values ...string) error {
	if len(name)==0 || len(values[0])==0 {
		return nil
	}
	if t.FormValues == nil {
		t.FormValues = url.Values{}
	}
	t.FormValues.Set(name, values[0])
	return nil
}

func (t *ApiRequest) SetFileParam(_ string, _ string) error { return nil }

func (t *ApiRequest) SetBodyParam(body interface{}) error {
	t.Body = body
	return nil
}

func (t *ApiRequest) GetPath() string {
	if len(t.QueryString) != 0 {
		return t.Path + "?" + t.QueryString
	}
	return t.Path
}

func (t *ApiRequest) GetFormValues() url.Values {
	return t.FormValues
}

