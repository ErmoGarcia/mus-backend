package main

import (
	"log"

	"github.com/ErmoGarcia/mus-backend/controllers/auth"
	"github.com/ErmoGarcia/mus-backend/db"
	"github.com/ErmoGarcia/mus-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = db.ConnectDataBase()
	if err != nil {
		log.Fatalf("Error loading database")
	}

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)

	protected := r.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/profile", auth.CurrentUser)

	r.Run(":8080")

}
