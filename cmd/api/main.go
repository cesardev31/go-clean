package main

import (
	"github/cesardev31/go-clean/cmd/api/handlers/player"
	"github/cesardev31/go-clean/internal/repositories/mongo"
	playerRepository "github/cesardev31/go-clean/internal/repositories/mongo/player"
	PlayerService "github/cesardev31/go-clean/internal/services/player"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .nev file")
	}
	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	playerRepo := playerRepository.Repository{
		Client: client,
	}

	playerSrv := PlayerService.Service{
		Repo: playerRepo,
	}

	playerHandler := player.Handler{
		PlayerService: playerSrv,
	}

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	log.Fatalln(ginEngine.Run(":8001"))
}
