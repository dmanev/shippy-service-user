package main

import (
    "context"
    "fmt"

    pb "github.com/dmanev/shippy-service-user/proto/user"
    "github.com/micro/go-micro"
    "log"
    "os"
)

type service struct {
    repo repository
}

const (
    defaultHost = "datastore:27017"
)

func main() {
    srv := micro.NewService(
        micro.Name("shippy.service.user"),
    )

    srv.Init()

    uri := os.Getenv("DB_HOST")
    if uri == "" {
        uri = defaultHost
    }
    client, err := CreateClient(uri)
    if err != nil {
        log.Panic(err)
    }
    defer client.Disconnect(context.TODO())

    userCollection := client.Database("shippy").Collection("user")
    repository := &userRepository{
        userCollection,
    }

    // Register our implementation with
    pb.RegisteruserServiceHandler(srv.Server(), &handler{repository})

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
