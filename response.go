package http

import (
	"context"
	"net/http"

	"go.unistack.org/micro/v4/metadata"
)

type (
	rspStatusCodeKey struct{}
	rspStatusCodeVal struct {
		code int
	}
)

// SetResponseStatusCode sets the status code in the context.
func SetResponseStatusCode(ctx context.Context, code int) {
	if rsp, ok := ctx.Value(rspStatusCodeKey{}).(*rspStatusCodeVal); ok {
		rsp.code = code
	}
}

// getResponseStatusCode retrieves the response status code from the context.
func getResponseStatusCode(ctx context.Context) int {
	code := http.StatusOK
	if rsp, ok := ctx.Value(rspStatusCodeKey{}).(*rspStatusCodeVal); ok {
		code = rsp.code
	}
	return code
}

type (
	rspMetadataKey struct{}
	rspMetadataVal struct {
		m metadata.Metadata
	}
)

// AppendResponseMetadata adds metadata entries to metadata.Metadata stored in the context.
// It expects the context to contain a *rspMetadataVal value under the rspMetadataKey{} key.
// If the value is missing or invalid, the function does nothing.
//
// Note: this function is not thread-safe. Synchronization is required if used from multiple goroutines.
func AppendResponseMetadata(ctx context.Context, md metadata.Metadata) {
	if md == nil {
		return
	}

	val, ok := ctx.Value(rspMetadataKey{}).(*rspMetadataVal)
	if !ok || val == nil || val.m == nil {
		return
	}

	for key, values := range md {
		val.m.Append(key, values...)
	}
}

// getResponseMetadata retrieves the metadata.Metadata stored in the context.
// If the value is missing, of the wrong type, or nil, it returns an empty metadata.Metadata.
//
// Note: this function is not thread-safe. Synchronization is required if used from multiple goroutines.
// If you plan to modify the returned metadata, make a full copy to avoid affecting shared state.
func getResponseMetadata(ctx context.Context) metadata.Metadata {
	val, ok := ctx.Value(rspMetadataKey{}).(*rspMetadataVal)
	if !ok || val == nil || val.m == nil {
		return metadata.Metadata{}
	}

	return val.m
}
