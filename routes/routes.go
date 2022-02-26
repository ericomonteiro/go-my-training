package routes

import (
	"github.com/ericomonteiro/go-my-training/controllers"
	"net/http"
)

func HandleRequest() {

	http.HandleFunc("/", controllers.Home)

	http.HandleFunc("/api/v1/exercise", controllers.GetExercises)

	http.ListenAndServe(":8000", nil)

}
