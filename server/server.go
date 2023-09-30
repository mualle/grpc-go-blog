package main

import (
	pb "github.com/mualle/grpc-go-blog/proto/blog/v1"
)

type Server struct {
	pb.UnimplementedBlogServiceServer
}
