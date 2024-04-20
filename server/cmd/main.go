package main

import (
	"github.com/danielsobrado/ainovelprompter/pkg/api"
	"github.com/danielsobrado/ainovelprompter/pkg/config"
	"github.com/danielsobrado/ainovelprompter/pkg/db"
	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/gin-contrib/cors"
)

func main() {
	config.LoadConfig()
	logging.InitLogger()

	dbConn, err := db.ConnectDB()
	if err != nil {
		logging.Logger.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	router := api.SetupRouter(dbConn)

	// Configure CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	port := config.GetString("server.port")
	logging.Logger.Infof("Server started on port %s", port)
	err = router.Run(":" + port)
	if err != nil {
		logging.Logger.Fatalf("Failed to start the server: %v", err)
	}
}
