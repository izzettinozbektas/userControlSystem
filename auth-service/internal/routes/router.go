package routes

import (
	"net/http"
	"github.com/izzettinozbektas/userControlSystem/auth-service/internal/controllers"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", controllers.Register)

	return mux
}
