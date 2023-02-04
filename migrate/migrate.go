package main

import (
	"github.com/bril3d/crud-api/initializers"
	"github.com/bril3d/crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
