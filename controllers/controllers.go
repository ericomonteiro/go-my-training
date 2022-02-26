package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ericomonteiro/go-my-training/models"
	"github.com/ericomonteiro/go-my-training/models/enum/muscle"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func GetExercises(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(exercises())
}

func exercises() []models.Exercise {
	return []models.Exercise{
		{Name: "Supino reto", Muscle: muscle.CHEST, ImageUrl: "https://treinomestre.com.br/wp-content/uploads/2017/08/supino-fechado-cp.jpg"},
	}
}
