package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	// Crea una instancia del motor Gin
	r := gin.Default()

	// Define una ruta y su manejado
	r.GET("/users/:numempleado", usersHandler)
	r.POST("/users", userCreateHandler)
	r.DELETE("users/:numempleado", userDeleteHandler)


	// Inicia el servidor en el puerto 8080
	router.Run()
}