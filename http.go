package nexmo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kevinburke/rest"
)

// The nexmo-go version. Run "make release" to bump this number.
const Version = "0.1"
const userAgent = "nexmo-go/" + Version

// The base URL serving the API. Override this for testing.
var BaseURL = "https://rest.nexmo.com"

type Client struct {
	*rest.Client

	ApiKey    string
	ApiSecret string

	// FullPath takes a path part (e.g. "Messages") and
	// returns the full API path, including the version (e.g.
	// "/sms").
	FullPath func(pathPart string) string

	// The API Client uses these resources
	Messages *MessageService
}

const defaultTimeout = 30*time.Second + 500*time.Millisecond

var defaultHttpClient *http.Client

func init() {
	defaultHttpClient = &http.Client{
		Timeout:   defaultTimeout,
		Transport: rest.DefaultTransport,
	}
}

// A message error returned by the Nexmo API.
type nexmoMessageError struct {
	MessageCount string `json:"message-count"`
	Messages     []struct {
		Status    string `json:"status"`
		ErrorText string `json:"error-text"`
	} `json:"messages"`
}

func parseNexmoError(resp *http.Response) error {
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := resp.Body.Close(); err != nil {
		return err
	}

	rerr := &nexmoMessageError{}
	err = json.Unmarshal(resBody, rerr)
	if err != nil {
		return fmt.Errorf("invalid response body: %s", string(resBody))
	}
	if rerr.Messages[0].ErrorText == "" {
		return fmt.Errorf("invalid response body: %s", string(resBody))
	}

	return &rest.Error{
		Title:  rerr.Messages[0].ErrorText,
		Type:   "Message Error",
		ID:     rerr.Messages[0].Status,
		Status: resp.StatusCode,
	}
}

// NewClient creates a Client for interacting with the Nexmo API. This is the
// main entrypoint for API interactions; view the methods on the subresources
// for more information.
func NewClient(apiKey string, apiSecret string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = defaultHttpClient
	}
	restClient := rest.NewClient("", "", BaseURL)
	restClient.Client = httpClient
	restClient.UploadType = rest.FormURLEncoded
	restClient.ErrorParser = parseNexmoError

	c := &Client{Client: restClient, ApiKey: apiKey, ApiSecret: apiSecret}
	c.FullPath = func(pathPart string) string {
		return "/" + pathPart
	}

	c.Messages = &MessageService{client: c}

	return c
}

func (c *Client) GetResource(ctx context.Context, pathPart string, sid string, v interface{}) error {
	sidPart := strings.Join([]string{pathPart, sid}, "/")

	return c.MakeRequest(ctx, "GET", sidPart, nil, v)
}

func (c *Client) CreateResource(ctx context.Context, pathPart string, data url.Values, v interface{}) error {
	return c.MakeRequest(ctx, "POST", pathPart, data, v)
}

func (c *Client) UpdateResource(ctx context.Context, pathPart string, sid string, data url.Values, v interface{}) error {
	sidPart := strings.Join([]string{pathPart, sid}, "/")

	return c.MakeRequest(ctx, "POST", sidPart, data, v)
}

func (c *Client) DeleteResource(ctx context.Context, pathPart string, sid string) error {
	sidPart := strings.Join([]string{pathPart, sid}, "/")
	err := c.MakeRequest(ctx, "DELETE", sidPart, nil, nil)
	if err == nil {
		return nil
	}
	rerr, ok := err.(*rest.Error)
	if ok && rerr.Status == http.StatusNotFound {
		return nil
	}

	return err
}

func (c *Client) ListResource(ctx context.Context, pathPart string, data url.Values, v interface{}) error {
	return c.MakeRequest(ctx, "GET", pathPart, data, v)
}

// Make a request to the Nexmo API.
func (c *Client) MakeRequest(ctx context.Context, method string, pathPart string, data url.Values, v interface{}) error {
	data.Add("api_key", c.ApiKey)
	data.Add("api_secret", c.ApiSecret)
	if !strings.HasPrefix(pathPart, "/") {
		pathPart = c.FullPath(pathPart)
	}
	rb := new(strings.Reader)
	if data != nil && (method == "POST" || method == "PUT") {
		rb = strings.NewReader(data.Encode())
	}
	if method == "GET" && data != nil {
		pathPart = pathPart + "?" + data.Encode()
	}
	req, err := c.NewRequest(method, pathPart, rb)
	if err != nil {
		return err
	}
	req = withContext(req, ctx)
	if ua := req.Header.Get("User-Agent"); ua == "" {
		req.Header.Set("User-Agent", userAgent)
	} else {
		req.Header.Set("User-Agent", userAgent+" "+ua)
	}

	return c.Do(req, &v)
}
