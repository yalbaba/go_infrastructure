// +build go1.9

package iris

import (
	"github.com/yalbaba/go_infrastructure/pkg/iris/context"
	"github.com/yalbaba/go_infrastructure/pkg/iris/core/host"
	"github.com/yalbaba/go_infrastructure/pkg/iris/core/router"
)

type (
	// Context is the middle-man server's "object" for the clients.
	//
	// A New context is being acquired from a sync.Pool on each connection.
	// The Context is the most important thing on the iris's http flow.
	//
	// Developers send responses to the client's request through a Context.
	// Developers get request information from the client's request by a Context.
	Context = context.Context
	// UnmarshalerFunc a shortcut, an alias for the `context#UnmarshalerFunc` type
	// which implements the `context#Unmarshaler` interface for reading request's body
	// via custom decoders, most of them already implement the `context#UnmarshalerFunc`
	// like the json.Unmarshal, xml.Unmarshal, yaml.Unmarshal and every library which
	// follows the best practises and is aligned with the Go standards.
	//
	// See 'context#UnmarshalBody` for more.
	//
	// Example: https://github.com/kataras/iris/blob/master/_examples/http_request/read-custom-via-unmarshaler/main.go
	UnmarshalerFunc = context.UnmarshalerFunc
	// A Handler responds to an HTTP request.
	// It writes reply headers and data to the Context.ResponseWriter() and then return.
	// Returning signals that the request is finished;
	// it is not valid to use the Context after or concurrently with the completion of the Handler call.
	//
	// Depending on the HTTP client software, HTTP protocol version,
	// and any intermediaries between the client and the iris server,
	// it may not be possible to read from the Context.Request().Body after writing to the context.ResponseWriter().
	// Cautious handlers should read the Context.Request().Body first, and then reply.
	//
	// Except for reading the body, handlers should not modify the provided Context.
	//
	// If Handler panics, the server (the caller of Handler) assumes that the effect of the panic was isolated to the active request.
	// It recovers the panic, logs a stack trace to the server error log, and hangs up the connection.
	Handler = context.Handler
	// Filter is just a type of func(Handler) bool which reports whether an action must be performed
	// based on the incoming request.
	//
	// See `NewConditionalHandler` for more.
	// An alias for the `context/Filter`.
	Filter = context.Filter
	// A Map is an alias of map[string]interface{}.
	Map = context.Map
	// Problem Details for HTTP APIs.
	// Pass a Problem value to `context.Problem` to
	// write an "application/problem+json" response.
	//
	// Read more at: https://github.com/kataras/iris/wiki/Routing-error-handlers
	//
	// It is an alias of the `context#Problem` type.
	Problem = context.Problem
	// ProblemOptions the optional settings when server replies with a Problem.
	// See `Context.Problem` method and `Problem` type for more details.
	//
	// It is an alias of the `context#ProblemOptions` type.
	ProblemOptions = context.ProblemOptions
	// JSON the optional settings for JSON renderer.
	//
	// It is an alias of the `context#JSON` type.
	JSON = context.JSON
	// XML the optional settings for XML renderer.
	//
	// It is an alias of the `context#XML` type.
	XML = context.XML
	// Supervisor is a shortcut of the `host#Supervisor`.
	// Used to add supervisor configurators on common Runners
	// without the need of importing the `core/host` package.
	Supervisor = host.Supervisor

	// Party is just a group joiner of routes which have the same prefix and share same middleware(s) also.
	// Party could also be named as 'Join' or 'Node' or 'Group' , Party chosen because it is fun.
	//
	// Look the `core/router#APIBuilder` for its implementation.
	//
	// A shortcut for the `core/router#Party`, useful when `PartyFunc` is being used.
	Party = router.Party
	// DirOptions contains the optional settings that
	// `FileServer` and `Party#HandleDir` can use to serve files and assets.
	// A shortcut for the `router.DirOptions`, useful when `FileServer` or `HandleDir` is being used.
	DirOptions = router.DirOptions
	// ExecutionRules gives control to the execution of the route handlers outside of the handlers themselves.
	// Usage:
	// Party#SetExecutionRules(ExecutionRules {
	//   Done: ExecutionOptions{Force: true},
	// })
	//
	// See `core/router/Party#SetExecutionRules` for more.
	// Example: https://github.com/kataras/iris/tree/master/_examples/mvc/middleware/without-ctx-next
	ExecutionRules = router.ExecutionRules
	// ExecutionOptions is a set of default behaviors that can be changed in order to customize the execution flow of the routes' handlers with ease.
	//
	// See `ExecutionRules` and `core/router/Party#SetExecutionRules` for more.
	ExecutionOptions = router.ExecutionOptions

	// CookieOption is the type of function that is accepted on
	// context's methods like `SetCookieKV`, `RemoveCookie` and `SetCookie`
	// as their (last) variadic input argument to amend the end cookie's form.
	//
	// Any custom or builtin `CookieOption` is valid,
	// see `CookiePath`, `CookieCleanPath`, `CookieExpires` and `CookieHTTPOnly` for more.
	//
	// An alias for the `context/Context#CookieOption`.
	CookieOption = context.CookieOption
	// N is a struct which can be passed on the `Context.Negotiate` method.
	// It contains fields which should be filled based on the `Context.Negotiation()`
	// server side values. If no matched mime then its "Other" field will be sent,
	// which should be a string or []byte.
	// It completes the `context/context.ContentSelector` interface.
	//
	// An alias for the `context/Context#N`.
	N = context.N
)
