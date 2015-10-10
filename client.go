package coding

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client is a local Client client.
type Client struct {
	Client      *http.Client
	AccessToken string
	APIPath     string

	ProjectService *ProjectService
	HookService    *HookService
}

const (
	DefaultTimeout = 30 * time.Second
	DafaultApiPath = "https://coding.net/api"
)

// NewClient creates a new Client client with a default timeout.
func NewClient(apiPath string, token string) *Client {
	return NewClientTimeout(apiPath, token, DefaultTimeout)
}

// NewClientTimeout acts like NewClient but takes a timeout.
func NewClientTimeout(apiPath string, token string, timeout time.Duration) *Client {
	apiPath = strings.TrimRight(apiPath, "/")
	c := &Client{
		Client:      &http.Client{Timeout: timeout},
		AccessToken: token,
		APIPath:     apiPath,
	}

	c.ProjectService = &ProjectService{client: c}
	c.HookService = &HookService{client: c}

	return c
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c := resp.StatusCode; !(200 <= c && c <= 299) {

		return nil, errors.New(fmt.Sprintf("response code %d", c))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err = CheckReponse(resp, data); err != nil {
		return resp, err
	}

	if v != nil {
		err = json.Unmarshal(data, v)
	}
	return resp, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method,
		c.APIPath+url+"?access_token="+c.AccessToken,
		body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type ErrorResponse struct {
	Response *http.Response
	Code     int               `json:"code"`
	Msg      map[string]string `json:"msg"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Code, r.Msg)

}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckReponse(r *http.Response, data []byte) error {
	errorResponse := &ErrorResponse{Response: r, Msg: make(map[string]string)}
	err := json.Unmarshal(data, errorResponse)
	if err == nil && errorResponse.Code != 0 {
		return errorResponse
	}
	return nil
}
