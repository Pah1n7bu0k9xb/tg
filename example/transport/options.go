// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type ServiceRoute interface {
	SetRoutes(route *fiber.App)
}

type Option func(srv *Server)
type Handler = fiber.Handler
type ErrorHandler func(err error) error

func Service(svc ServiceRoute) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			svc.SetRoutes(srv.Fiber())
		}
	}
}

func JsonRPC(svc *httpJsonRPC) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			srv.httpJsonRPC = svc
			svc.SetRoutes(srv.Fiber())
		}
	}
}

func User(svc *httpUser) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			srv.httpUser = svc
			svc.SetRoutes(srv.Fiber())
		}
	}
}

func MaxBodySize(max int) Option {
	return func(srv *Server) {
		srv.config.BodyLimit = max
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(srv *Server) {
		srv.config.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(srv *Server) {
		srv.config.WriteTimeout = timeout
	}
}

func Use(args ...interface{}) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			srv.srvHTTP.Use(args...)
		}
	}
}
