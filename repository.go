package main

import (

    "context"
    pb "github.com/dmanev/shippy-service-user/proto/user"
    "go.mongodb.org/mongo-driver/mongo"
    "gopkg.in/mgo.v2/bson"
)

type repository interface {
    FindAvailable(*pb.Specification) (*pb.user, error)
    Create(user *pb.user) error
}

type userRepository struct {
    collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of users,
// if capacity and max weight are below a users capacity and max weight,
// then return that user.
func (repository *userRepository) FindAvailable(spec *pb.Specification) (*pb.user, error) {
    filter := bson.D{{
        "capacity",
        bson.D{{
            "$lte",
            spec.Capacity,
        }, {
            "$lte",
            spec.MaxWeight,
        }},
    }}
    var user *pb.user
    if err := repository.collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
        return nil, err
    }
    return user, nil
}

// Create a new user
func (repository *userRepository) Create(user *pb.user) error {
    _, err := repository.collection.InsertOne(context.TODO(), user)
    return err
}
