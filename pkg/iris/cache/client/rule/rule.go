package rule

import (
	"liveearth/infrastructure/pkg/iris/context"
)

// Rule a superset of validators
type Rule interface {
	Claim(ctx context.Context) bool
	Valid(ctx context.Context) bool
}