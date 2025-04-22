package http

import (
	"context"
	"net/http"

	"go.unistack.org/micro/v4/metadata"
)

type (
	rspHeaderKey struct{}
	rspHeaderVal struct {
		h http.Header
	}
)

// AppendResponseMetadata adds metadata entries to an http.Header stored in the context.
// It expects the context to have a *rspHeaderVal value under the rspHeaderKey{} key.
// If the value is missing or invalid, the function does nothing.
//
// Note: This function is not thread-safe.
func AppendResponseMetadata(ctx context.Context, md metadata.Metadata) {
	if md == nil {
		return
	}

	header, ok := ctx.Value(rspHeaderKey{}).(*rspHeaderVal)
	if !ok || header == nil || header.h == nil {
		return
	}

	for key, values := range md {
		for _, value := range values {
			header.h.Add(key, value)
		}
	}
}
