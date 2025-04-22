package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"go.unistack.org/micro/v4/metadata"
)

func TestAppendResponseMetadata(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		md       metadata.Metadata
		expected context.Context
	}{
		{
			name:     "nil metadata",
			ctx:      context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
			md:       nil,
			expected: context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
		},
		{
			name:     "empty metadata",
			ctx:      context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
			md:       metadata.Metadata{},
			expected: context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
		},
		{
			name:     "context without response header key",
			ctx:      context.Background(),
			md:       metadata.Pairs("key1", "val1"),
			expected: context.Background(),
		},
		{
			name:     "context with nil response header value",
			ctx:      context.WithValue(context.Background(), rspHeaderKey{}, nil),
			md:       metadata.Pairs("key1", "val1"),
			expected: context.WithValue(context.Background(), rspHeaderKey{}, nil),
		},
		{
			name:     "context with response header value, but nil http.Header",
			ctx:      context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: nil}),
			md:       metadata.Pairs("key1", "val1"),
			expected: context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: nil}),
		},
		{
			name: "basic metadata append",
			ctx:  context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
			md:   metadata.Pairs("key1", "val1"),
			expected: context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{
				h: http.Header{
					"Key1": []string{"val1"},
				},
			}),
		},
		{
			name: "multiple values for same key",
			ctx:  context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
			md:   metadata.Pairs("key1", "val1", "key1", "val2"),
			expected: context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{
				h: http.Header{
					"Key1": []string{"val1", "val2"},
				},
			}),
		},
		{
			name: "multiple values for different keys",
			ctx:  context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{h: http.Header{}}),
			md:   metadata.Pairs("key1", "val1", "key1", "val2", "key2", "val3", "key2", "val4", "key3", "val5"),
			expected: context.WithValue(context.Background(), rspHeaderKey{}, &rspHeaderVal{
				h: http.Header{
					"Key1": []string{"val1", "val2"},
					"Key2": []string{"val3", "val4"},
					"Key3": []string{"val5"},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppendResponseMetadata(tt.ctx, tt.md)
			require.Equal(t, tt.expected, tt.ctx)
		})
	}
}
