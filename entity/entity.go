package entity

import "time"

// User is user models property
type Bihin struct {
	ID          int       `json:"id"`
	KnrId       string    `json:"knrid"`
	Bihin       string    `json:"bihin"`
	Dantai      string    `json:"dantai"`
	Place       string    `json:"place"`
	Price       string    `json:"price"`
	Qty         string    `json:"qty"`
	CreatedAt   time.Time `json:"createdat"`
	PartNum     string    `json:"partnum"`
	Note        string    `json:"note"`
	Shishutsuid string    `json:"shishutsuid"`
}
