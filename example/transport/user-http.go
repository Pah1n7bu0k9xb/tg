// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/seniorGolang/tg/example/implement"
	"github.com/seniorGolang/tg/example/interfaces"
)

type httpUser struct {
	log          logrus.FieldLogger
	errorHandler ErrorHandler
	svc          *serverUser
	base         interfaces.User
}

func NewUser(log logrus.FieldLogger, svcUser interfaces.User) (srv *httpUser) {

	srv = &httpUser{
		base: svcUser,
		log:  log,
		svc:  newServerUser(svcUser),
	}
	return
}

func (http httpUser) Service() MiddlewareSetUser {
	return http.svc
}

func (http *httpUser) WithLog(log logrus.FieldLogger) *httpUser {
	http.svc.WithLog(log)
	return http
}

func (http *httpUser) WithTrace() *httpUser {
	http.svc.WithTrace()
	return http
}

func (http *httpUser) WithMetrics() *httpUser {
	http.svc.WithMetrics()
	return http
}

func (http *httpUser) WithErrorHandler(handler ErrorHandler) *httpUser {
	http.errorHandler = handler
	return http
}

func (http *httpUser) SetRoutes(route *router.Router) {

	route.GET("/api/v2/user/info", http.serveGetUser)
	route.POST("/api/v2/user/file", http.serveUploadFile)
	route.PATCH("/api/v2/user/custom/response", http.serveCustomResponse)
	route.DELETE("/api/v2/user/custom", func(ctx *fasthttp.RequestCtx) {
		implement.CustomHandler(ctx, http.base)
	})
}