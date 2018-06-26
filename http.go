package nexmo

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kevinburke/rest"
)

// Version of the program
const Version = "0.1"
const userAgent = "nexmo-go/" + Version

// BaseURL serving the API. Override this for testing.
var BaseURL = "https://rest.nexmo.com"

// Client holds all the necessary data to handle nexmo API
type Client struct {
	*rest.Client

	APIKey    string
	APISecret string

	// FullPath takes a path part (e.g. "Messages") and
	// returns the full API path, including the version (e.g.
	// "/sms").
	FullPath func(pathPart string) string

	// The API Client uses these resources
	Messages *MessageService
}

const defaultTimeout = 30*time.Second + 500*time.Millisecond

var defaultHTTPClient *http.Client

func init() {
	defaultHTTPClient = &http.Client{
		Timeout:   defaultTimeout,
		Transport: rest.DefaultTransport,
	}
}

// NewClient creates a Client for interacting with the Nexmo API. This is the
// main entrypoint for API interactions; view the methods on the subresources
// for more information.
func NewClient(apiKey string, apiSecret string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = defaultHTTPClient
	}
	restClient := rest.NewClient("", "", BaseURL)
	restClient.Client = httpClient
	restClient.UploadType = rest.FormURLEncoded

	c := &Client{Client: restClient, APIKey: apiKey, APISecret: apiSecret}
	c.FullPath = func(pathPart string) string {
		return "/" + pathPart
	}

	c.Messages = &MessageService{client: c}

	return c
}

// TODO(david.lopez) functions took from: https://github.com/kevinburke/twilio-go/blob/master/http.go
// Add GetResource, UpdateResource, DeleteResource... as needed

// CreateResource handles POST requests
func (c *Client) CreateResource(ctx context.Context, pathPart string, data url.Values, v interface{}) error {
	return c.MakeRequest(ctx, "POST", pathPart, data, v)
}

// MakeRequest makes a request to the Nexmo API.
func (c *Client) MakeRequest(ctx context.Context, method string, pathPart string, data url.Values, v interface{}) error {
	data.Add("api_key", c.APIKey)
	data.Add("api_secret", c.APISecret)
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
	req = withContext(ctx, req)
	if ua := req.Header.Get("User-Agent"); ua == "" {
		req.Header.Set("User-Agent", userAgent)
	} else {
		req.Header.Set("User-Agent", userAgent+" "+ua)
	}

	return c.Do(req, &v)
}
