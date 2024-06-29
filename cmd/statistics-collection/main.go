package main

import (
	"fmt"
	"log"

	"github.com/zhetkerbaevan/statistics_collection/cmd/api"
	"github.com/zhetkerbaevan/statistics_collection/internal/config"
	"github.com/zhetkerbaevan/statistics_collection/internal/db"
)

func main() {
	db, err := db.NewPostgreSQLStorage(config.Config{
		DBHost:     config.Envs.DBHost,
		DBPort:     config.Envs.DBPort,
		DBUser:     config.Envs.DBUser,
		DBName:     config.Envs.DBName,
		DBPassword: config.Envs.DBPassword,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to db")

	//Start server
	server := api.NewAPIServer(db, ":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}