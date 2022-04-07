package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/febzey/forestbot-api/pkg/configs"
	"github.com/febzey/forestbot-api/pkg/routes"
	"github.com/febzey/forestbot-api/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func FileServer(router *mux.Router) {
	dir, _ := utils.GetCurrentDirectory()
	webPath := filepath.Join(dir, "../web/dist")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(webPath))))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	routes.PublicRoutes(router)
	router.Use(mux.CORSMethodMiddleware(router))
	server := configs.ServerConfig(router)
	utils.StartServer(server)
}
