package bihin

import (
	"fmt"

	"github.com/gin-gonic/gin"

	bihin "Bihins-Management-System-API/service"
)

type Controller struct{}

// Index action: GET /bihins
func (pc Controller) Index(c *gin.Context) {
	var s bihin.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /bihins
func (pc Controller) Create(c *gin.Context) {
	var s bihin.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Showid action: GET /bihins/:id
func (pc Controller) Showid(c *gin.Context) {
	id := c.Params.ByName("id")
	var s bihin.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Showid action: GET /bihins/:dantaimei
func (pc Controller) ShowDantaimei(c *gin.Context) {
	dan := c.Params.ByName("dantaimei")
	var s bihin.Service
	p, err := s.GetByDantaimei(dan)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /bihins/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s bihin.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /bihins/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s bihin.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
