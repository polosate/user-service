package main

import (
	"golang.org/x/net/context"
	pb "github.com/polosate/user-service/proto/user"
)

type service struct {
	repo Repository
}

func (srv *service) Get(ctx  context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(ctx, req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := srv.repo.Create(ctx, req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := srv.repo.GetByEmail(ctx, req)
	if err != nil {
		return err
	}
	res.Token = "test_abc"
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}

