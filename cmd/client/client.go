package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/o-sokol-o/as/pkg/domain"
	client "github.com/o-sokol-o/as/pkg/grpc_client"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация читалки конфига
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("configs/") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Fatalf("fatal error config file: %+v", err)
	}

	// Set up a connection to the server.
	target := viper.GetString("Server.uri")
	grpc, err := client.New(target)
	if err != nil {
		log.Fatalf("server: %s\n%v\n", target, err)
	}
	defer grpc.Close()

	err = grpc.Fetch(viper.GetString("Product.uri"))
	if err != nil {
		log.Fatalf("error grpc fetch:%v\n", err)
	}

	err = List(grpc,
		domain.Paging{NumPage: 1, ProductPerPage: 2},
		domain.Sorting{SortField: domain.Sort_Name, Direction: domain.SortNone})
	if err != nil {
		log.Fatalf("error grpc list: %v\n", err)
	}

	err = List(grpc,
		domain.Paging{NumPage: 1, ProductPerPage: 2},
		domain.Sorting{SortField: domain.Sort_Name, Direction: domain.SortAsc})
	if err != nil {
		log.Fatalf("error grpc list: %v\n", err)
	}

	err = List(grpc,
		domain.Paging{NumPage: 1, ProductPerPage: 2},
		domain.Sorting{SortField: domain.Sort_Name, Direction: domain.SortDesc})
	if err != nil {
		log.Fatalf("error grpc list: %v\n", err)
	}

	err = List(grpc,
		domain.Paging{NumPage: 1, ProductPerPage: 5},
		domain.Sorting{SortField: domain.Sort_Cost, Direction: domain.SortDesc})
	if err != nil {
		log.Fatalf("error grpc list: %v\n", err)
	}

}

func List(grpc *client.GrpcClient, p domain.Paging, s domain.Sorting) error {
	resp, err := grpc.List(p, s)
	if err != nil {
		return err
	}

	// Печатаем
	respJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(respJSON))

	return nil
}
