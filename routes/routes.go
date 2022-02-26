package routes

import (
	"github.com/ericomonteiro/go-my-training/controllers"
	"net/http"
)

func HandleRequest() {

	http.HandleFunc(
		"/",
		controllers.Home,
	)
	http.ListenAndServe(":8000", nil)

}
