package entity

import (
	"time"
)

type RefreshToken struct {
	Id           string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserId       string    `json:"user_id" bson:"user_id"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}

type SignTokens struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

type ResetTokens struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
