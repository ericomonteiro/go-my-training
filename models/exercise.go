package models

import (
	"github.com/ericomonteiro/go-my-training/models/enum/muscle"
)

type Exercise struct {
	Id       int64         `json:"id"`
	Name     string        `json:"name"`
	Muscle   muscle.Muscle `json:"muscle"`
	ImageUrl string        `json:"image"`
}
