package ratelimit_test

import (
	"github.com/lnyyj/ggrpc/middleware"
	"github.com/lnyyj/ggrpc/middleware/ratelimit"
	"google.golang.org/grpc"
)

// alwaysPassLimiter is an example limiter which implements Limiter interface.
// It does not limit any request because Limit function always returns false.
type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}

// Simple example of server initialization code.
func Example() {
	// Create unary/stream rateLimiters, based on token bucket here.
	// You can implement your own ratelimiter for the interface.
	limiter := &alwaysPassLimiter{}
	_ = grpc.NewServer(
		middleware.WithUnaryServerChain(
			ratelimit.UnaryServerInterceptor(limiter),
		),
		middleware.WithStreamServerChain(
			ratelimit.StreamServerInterceptor(limiter),
		),
	)
}
