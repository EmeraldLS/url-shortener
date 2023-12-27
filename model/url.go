package model

import "time"

type URL struct {
	OriginalUrl    string    `json:"original_url,omitempty" bson:"original_url,omitempty"`
	ShortUrl       string    `json:"short_url,omitempty" bson:"short_url,omitempty"`
	Id             string    `json:"id,omitempty" bson:"id,omitempty"`
	NoOfClicks     uint8     `json:"no_of_clicks,omitempty" bson:"no_of_clicks,omitempty"`
	CreatedAt      string    `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      string    `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	ExpirationTime time.Time `json:"expiration_time,omitempty" bson:"expiration_time,omitempty"`
}

type URL_SHORTENER_REQUEST struct {
	OriginalUrl string `json:"original_url,omitempty" validate:"required"`
}
