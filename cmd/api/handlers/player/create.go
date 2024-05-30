package player

import (
	"github/cesardev31/go-clean/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePlayer(c *gin.Context) {
	var playerCreateParams domain.Player
	if err := c.BindJSON(&playerCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insertedID, err := h.PlayerService.Create(playerCreateParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	c.JSON(200, gin.H{"player_id": insertedID})

}
