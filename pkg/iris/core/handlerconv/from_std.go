package handlerconv

import (
	"fmt"
	"net/http"

	"github.com/yalbaba/go_infrastructure/pkg/iris/context"
)

// FromStd converts native http.Handler & http.HandlerFunc to context.Handler.
//
// Supported form types:
// 		 .FromStd(h http.Handler)
// 		 .FromStd(func(w http.ResponseWriter, r *http.Request))
// 		 .FromStd(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc))
func FromStd(handler interface{}) context.Handler {
	switch h := handler.(type) {
	case context.Handler:
		{
			//
			// it's already a iris handler
			//
			return h
		}

	case http.Handler:
		//
		// handlerFunc.ServeHTTP(w,r)
		//
		{
			return func(ctx context.Context) {
				h.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
			}
		}

	case func(http.ResponseWriter, *http.Request):
		{
			//
			// handlerFunc(w,r)
			//
			return FromStd(http.HandlerFunc(h))
		}

	case func(http.ResponseWriter, *http.Request, http.HandlerFunc):
		{
			//
			// handlerFunc(w,r, http.HandlerFunc)
			//
			return FromStdWithNext(h)
		}

	default:
		{
			//
			// No valid handler passed
			//
			panic(fmt.Errorf(`
			Passed argument is not a func(context.Context) neither one of these types:
			- http.Handler
			- func(w http.ResponseWriter, r *http.Request)
			- func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
			---------------------------------------------------------------------
			It seems to be a  %T points to: %v`, handler, handler))
		}

	}
}

// FromStdWithNext receives a standar handler - middleware form - and returns a
// compatible context.Handler wrapper.
func FromStdWithNext(h func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)) context.Handler {
	return func(ctx context.Context) {
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx.ResetRequest(r)
			ctx.Next()
		})

		h(ctx.ResponseWriter(), ctx.Request(), next)
	}
}
