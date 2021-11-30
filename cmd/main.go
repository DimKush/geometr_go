package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/DimKush/geometry_go"
	"github.com/DimKush/geometry_go/internal/handler"
	"github.com/DimKush/geometry_go/internal/repository"
	"github.com/DimKush/geometry_go/internal/service"
	"github.com/spf13/viper"
)

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}
}

func main() {
	serverInst := new(server.Server)
	if err := run(serverInst); err != nil {
		panic(err.Error())
	}

}

func run(server *server.Server) error {
	fmt.Println("Poop")
	dbConfig := repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Dbname:   viper.GetString("database.dbname"),
		Timezone: viper.GetString("database.timezone"),
		SSLMode:  viper.GetString("database.sslmode"),
	}

	db, err := repository.NewPostgresConnection(dbConfig)
	if err != nil {
		return fmt.Errorf("Cannot create db connection %v.\nReason: %s", dbConfig, err.Error())
	}

	rep := repository.InitRepository(db)
	services := service.InitService(rep)
	handlers := handler.InitHandler(services)

	//service.Audit = service.InitAudit(repository, viper.GetString("audit_level"))

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		return fmt.Errorf("Cannot run server on port : %s. Reason : %s", viper.GetString("port"), err.Error())
	}

	return nil
}

func initConfig() error {
	platform := strings.ToLower(runtime.GOOS)

	var confDirPath string
	switch platform {
	case "linux":
		{
			confDirPath = "/opt/geometry-go/conf"
		}
	case "windows":
		{
			confDirPath = "c:\\geometry-go\\conf"
		}
	default:
		{
			confDirPath = "./configs"
		}
	}

	viper.AddConfigPath(confDirPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
