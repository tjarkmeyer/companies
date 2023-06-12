package controllers

import (
	"net/http"
)

// MiddlewareRoleCheck - checks if roles assigned to requesting user
func MiddlewareRoleCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Roles checks can happen here
		next.ServeHTTP(w, r)
	})
}
