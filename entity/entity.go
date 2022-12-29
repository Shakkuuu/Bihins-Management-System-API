package entity

import "time"

// User is user models property
type Bihin struct {
	ID          int       `json:"id"`
	Shishutsuid string    `json:"shishutsuid"`
	Dantaimei   string    `json:"dantaimei"`
	Bihin       string    `json:"bihin"`
	CreatedAt   time.Time `json:"createdat"`
}
