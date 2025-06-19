package routes

import "net/http"

func SetupRouter() http.Handler {
	return http.NewServeMux()
}
