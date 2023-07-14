package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routes_exist(t *testing.T) {
	testApp := NewApp(nil)
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	var expectedRoutes []string
	expectedRoutes = append(expectedRoutes, "/authentication")

	for _, route := range expectedRoutes {
		routeExists(t, chiRoutes, route)
	}

}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false
	walkFunc := func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	}
	_ = chi.Walk(routes, walkFunc)
	if !found {
		t.Errorf("route %s not found", route)
	}
}
