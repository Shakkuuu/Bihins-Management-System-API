package bihin

import (
	"github.com/gin-gonic/gin"

	db "Bihins-Management-System-API/dbgo"
	"Bihins-Management-System-API/entity"
)

type Service struct{}

type Bihin entity.Bihin

// GetAll is get all Bihin
func (s Service) GetAll() ([]Bihin, error) {
	db := db.GetDB()
	var b []Bihin

	if err := db.Find(&b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

// CreateModel is create Bihin model
func (s Service) CreateModel(c *gin.Context) (Bihin, error) {
	db := db.GetDB()
	var b Bihin

	if err := c.BindJSON(&b); err != nil {
		return b, err
	}

	if err := db.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

// GetByID is get a Bihin
func (s Service) GetByID(id string) (Bihin, error) {
	db := db.GetDB()
	var b Bihin

	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func (s Service) GetByDantaimei(dantaimei string) (Bihin, error) {
	db := db.GetDB()
	var b Bihin

	if err := db.Where("dantaimei LIKE ?", "%"+dantaimei+"%").Find(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

// UpdateByID is update a Bihin
func (s Service) UpdateByID(id string, c *gin.Context) (Bihin, error) {
	db := db.GetDB()
	var b Bihin

	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
		return b, err
	}

	if err := c.BindJSON(&b); err != nil {
		return b, err
	}

	db.Save(&b)

	return b, nil
}

// DeleteByID is delete a Bihin
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var b Bihin

	if err := db.Where("id = ?", id).Delete(&b).Error; err != nil {
		return err
	}

	return nil
}
