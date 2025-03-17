// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router interface {
	AddRoute() Route
	AddNotFoundHandler(handler http.Handler)
	AddMethodNotAllowedHandler(handler http.Handler)
}

type defaultRouter struct {
	i *mux.Router
}

var _ Router = &defaultRouter{}

type DefaultRouterOption func(router *defaultRouter)

func NewDefaultRouter(pathPrefix string, options ...DefaultRouterOption) *defaultRouter {
	router := &defaultRouter{}

	for _, option := range options {
		option(router)
	}

	if router.i == nil {
		router.i = mux.NewRouter()
	}

	router.i.SkipClean(true)

	if pathPrefix != "" {
		router.i = router.i.PathPrefix(pathPrefix).Subrouter()
	}

	return router
}

// WithGorillaMuxRouter sets underlying gorilla mux router to provided one
func WithGorillaMuxRouter(gorillaRouter *mux.Router) DefaultRouterOption {
	return func(router *defaultRouter) {
		router.i = gorillaRouter
	}
}

func (r *defaultRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.i.ServeHTTP(writer, request)
}

func (r *defaultRouter) AddNotFoundHandler(handler http.Handler) {
	r.i.NotFoundHandler = handler
}

func (r *defaultRouter) AddMethodNotAllowedHandler(handler http.Handler) {
	r.i.MethodNotAllowedHandler = handler
}

func (r *defaultRouter) AddRoute() Route {
	muxRoute := r.i.NewRoute()
	return &defaultRoute{muxRoute}
}

type Route interface {
	Methods(methods ...string) Route
	Path(tpl string) Route
	Queries(pairs ...string) Route
	Headers(pairs ...string) Route
	HeadersRegexp(pairs ...string) Route
	MatcherFunc(func(r *http.Request, match *RouteMatch) bool) Route
	Handler(handler http.Handler) Route
}

type RouteMatch struct {
	Route   Route
	Handler http.Handler
	Vars    map[string]string
}

var _ Route = &defaultRoute{}

type defaultRoute struct {
	*mux.Route
}

func (route *defaultRoute) Methods(methods ...string) Route {
	route.Route = route.Route.Methods(methods...)
	return route
}

func (route *defaultRoute) Path(tpl string) Route {
	route.Route = route.Route.Path(tpl)
	return route
}

func (route *defaultRoute) Queries(pairs ...string) Route {
	route.Route = route.Route.Queries(pairs...)
	return route
}

func (route *defaultRoute) Headers(pairs ...string) Route {
	route.Route = route.Route.Headers(pairs...)
	return route
}

func (route *defaultRoute) HeadersRegexp(pairs ...string) Route {
	route.Route = route.Route.HeadersRegexp(pairs...)
	return route
}

func (route *defaultRoute) MatcherFunc(f func(r *http.Request, match *RouteMatch) bool) Route {
	route.Route = route.Route.MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		routeMatch := &RouteMatch{
			Route:   route,
			Handler: match.Handler,
			Vars:    match.Vars,
		}
		return f(request, routeMatch)
	})
	return route
}

func (route *defaultRoute) Handler(handler http.Handler) Route {
	route.Route = route.Route.Handler(handler)
	return route
}
