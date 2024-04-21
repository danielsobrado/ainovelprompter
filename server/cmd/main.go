package main

import (
	"github.com/danielsobrado/ainovelprompter/pkg/api"
	"github.com/danielsobrado/ainovelprompter/pkg/config"
	"github.com/danielsobrado/ainovelprompter/pkg/db"
	"github.com/danielsobrado/ainovelprompter/pkg/logging"
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

	port := config.GetString("server.port")
	logging.Logger.Infof("Server started on port %s", port)
	err = router.Run(":" + port)
	if err != nil {
		logging.Logger.Fatalf("Failed to start the server: %v", err)
	}
}
