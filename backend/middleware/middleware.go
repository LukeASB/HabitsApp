package middleware

import (
	"dohabits/data"
	"dohabits/logger"
	"dohabits/middleware/session"
	"net/http"
)

type Middleware struct {
	jwtTokens  session.IJWTTokens
	csrfTokens session.ICSRFToken
	logger     logger.ILogger
}

type IMiddleware interface {
	MiddlewareList(handler http.HandlerFunc, dependencies data.Middleware) http.HandlerFunc
	protectedMiddlewareList() []func(http.HandlerFunc) http.HandlerFunc
	chainMiddleware(handler http.HandlerFunc, middlewares []func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc
}

func NewMiddleware(jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken, logger logger.ILogger) *Middleware {
	return &Middleware{
		jwtTokens:  jwtTokens,
		csrfTokens: csrfTokens,
		logger:     logger,
	}
}

func (mw *Middleware) MiddlewareList(handler http.HandlerFunc, dependencies data.Middleware) http.HandlerFunc {
	middlewares := []func(http.HandlerFunc) http.HandlerFunc{
		HTTPMethodValidation(dependencies.HTTPMethod, mw.logger),
		JSONMiddleware(mw.logger),
	}

	if dependencies.IsProtected {
		middlewares = append(middlewares, mw.protectedMiddlewareList()...)
	}

	if dependencies.CSRFRequired {
		mw.logger.DebugLog("middleware.MiddlewareList - CSRF Token commented out for debug")
		// middlewares = append(middlewares, CSRFToken(mw.csrfTokens, mw.logger)) // Commented out for debug
	}

	middlewares = append(middlewares, ErrorHandlingMiddleware(mw.logger))

	return mw.chainMiddleware(handler, middlewares)
}

func (mw *Middleware) protectedMiddlewareList() []func(http.HandlerFunc) http.HandlerFunc {
	return []func(http.HandlerFunc) http.HandlerFunc{
		session.AuthMiddleware(mw.jwtTokens, mw.logger),
	}
}

// Chain the listed middlewares in reverse order
func (mw *Middleware) chainMiddleware(handler http.HandlerFunc, middlewares []func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
