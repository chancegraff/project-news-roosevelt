package api

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/pronuu/roosevelt/internal/config"
	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/internal/server"
	"github.com/pronuu/roosevelt/internal/utils"
	"github.com/pronuu/roosevelt/pkg/client"
	"github.com/pronuu/roosevelt/pkg/user"

	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

func main() {
	// Create context, done channel, and logger
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()
	lgr := utils.GetLogger("roosevelt")

	// Create store
	str, err := database.NewStore()
	if err != nil {
		lgr.Log("err", "Database connection failed")
		cancel()
		return
	}

	// Create services
	user := user.NewUserService(str, lgr)
	client := client.NewClientService(str, lgr)

	// Create router
	rtr := mux.NewRouter()
	api := rtr.PathPrefix("/api/v1").Subrouter()
	user.Route(api)
	client.Route(api)

	// Create config and server
	cfg := config.NewConfig()
	srv := server.NewHTTP(cfg, lgr, rtr)
	defer srv.Stop(ctx)

	// Start server and bind until exit
	go srv.Start(ctx)
	<-*done
	cancel()
}
