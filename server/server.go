package server

import (
	"fmt"
	"github.com/99-66/SLx-Api/routes"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func RunAPI() error {
	host := "0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	address := fmt.Sprintf("%s:%s", host, port)

	return RunAPIWithMiddleware(&address)
}

func RunAPIWithMiddleware(address *string) error {
	g := gin.Default()

	// Set CORS Middlewares
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	g.Use(cors.New(corsConf))

	// Set Routes
	routes.InitRouters(g)

	return endless.ListenAndServe(*address, g)
}
