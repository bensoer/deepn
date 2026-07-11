package providers

import "context"

type Provider interface {
	Chat(ctx context.Context, req Request) (Response, error)
}
