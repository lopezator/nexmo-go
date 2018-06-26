// +build go1.7

package nexmo

import (
	"context"
	"net/http"
)

func withContext(ctx context.Context, r *http.Request) *http.Request {
	return r.WithContext(ctx)
}
