package models

import (
	"github.com/ericomonteiro/go-my-training/models/enum/muscle"
)

type Exercise struct {
	Name     string        `json:"name"`
	Muscle   muscle.Muscle `json:"muscle"`
	ImageUrl string        `json:"image"`
}
