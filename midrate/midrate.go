package main

import (
	"github.com/belayhun-arage/go-crud/initializers"
	"github.com/belayhun-arage/go-crud/models"
)

func init(){
	initializers.LoadEnvVariables();
	initializers.ConnectionDB();
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}