package metadata

import (
	"context"
)

// metaKey is the context value key we use to store metadata
type metaKey struct{}

// Metadata is our abstraction to store metadata
type Metadata map[string]string

// FromContext extracts the metadata from a context
func FromContext(ctx context.Context) (md Metadata, ok bool) {
	md, ok = ctx.Value(metaKey{}).(Metadata)
	return
}

// NewContext creates a new child context with metadata
func NewContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, metaKey{}, md)
}
