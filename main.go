package main

import (
	"fmt"
	// "log"
	// "gostudy/application/config"
	"gostudy/application/app/controllers"
	"gostudy/application/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}