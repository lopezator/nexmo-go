package nexmo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
)

type Server struct {
	s *httptest.Server
	// copied from httptest.Server
	URL string
	// URLs of incoming requests, in order
	URLs []*url.URL
	mu   sync.Mutex
}

func (s *Server) Close() {
	s.s.Close()
}

func (s *Server) CloseClientConnections() {
	s.s.CloseClientConnections()
}

func (s *Server) Start() {
	s.s.Start()
}

func newServer(response []byte, code int) *Server {
	serv := &Server{URLs: make([]*url.URL, 0)}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serv.mu.Lock()
		serv.URLs = append(serv.URLs, r.URL)
		serv.mu.Unlock()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(code)
		if _, err := w.Write(response); err != nil {
			panic(err)
		}
	}))
	serv.s = s
	serv.URL = s.URL
	return serv
}

// getServer returns a http server that returns the given bytes when requested,
// and a Client configured to make requests to that server.
func getServer(response []byte) (*Client, *Server) {
	s := newServer(response, 200)
	client := NewClient("AC123", "456", nil)
	client.Base = s.URL
	client.Messages.client.Base = s.URL

	return client, s
}

// useful trick: highlight the JSON range and hit `python -m json.tool` to
// pretty format it.
var sendMessageResponse = []byte(`
{
    "message-count": "1",
    "messages": [{
        "to": "19253920364",
        "message-id": "0D0000009A3D26AA",
        "status": "0",
        "remaining-balance": "59.43165823",
        "message-price": "0.06500000",
        "network": "21407"
    }]
}
`)

var sendCallResponse = []byte(`
{
	"call_id": "12abcdef111e1b3b4c5g4e23f1c222",
	"to": "34666666666",
	"status": "0",
	"error_text": "Success"
}
`)

const from = "11111111111"
const to = "11111111111"
