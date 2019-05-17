package main

import (
    "context"
    pb "github.com/dmanev/shippy-service-user/proto/user"
)

type handler struct {
    repository
}

// FindAvailable users
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

    // Find the next available user
    user, err := s.repository.FindAvailable(req)
    if err != nil {
        return err
    }

    // Set the user as part of the response message type
    res.user = user
    return nil
}

// Create a new user
func (s * handler) Create(ctx context.Context, req *pb.user, res *pb.Response) error {
    if err := s.repository.Create(req); err != nil {
        return err
    }
    res.user = req
    return nil
}
