package server

import (
	"context"
	"fmt"
	"latihan_grpc/common/proto"
	"latihan_grpc/helper"
	"latihan_grpc/modules/repository"
)

type HandlerUser struct {
	proto.UnimplementedUsersServer
	RepoUser repository.RepoUser
}

func NewHandlerUser(repouser repository.RepoUser) *HandlerUser {
	return &HandlerUser{
		RepoUser: repouser,
	}
}
func (h *HandlerUser) ShowAll(ctx context.Context, UserIn *proto.User) (*proto.DataUsers, error) {

	data, err := h.RepoUser.Show(ctx)
	if err != nil {

		return nil, err
	}
	d := helper.BindAllUser(data)

	return &proto.DataUsers{
		Data: d.Data,
	}, nil
}

func (h *HandlerUser) CreateUser(ctx context.Context, in *proto.User) (*proto.User, error) {
	fmt.Printf("possition  CreateUser %T \n", h)
	UserIn := helper.BindProtoToUser(in)
	res, err := h.RepoUser.Create(ctx, UserIn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return helper.BindUser(res), err

}

func (h *HandlerUser) FindById(ctx context.Context, in *proto.User) (*proto.User, error) {

	UserIn := helper.BindProtoToUser(in)
	res, err := h.RepoUser.FindById(ctx, UserIn.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return helper.BindUser(res), err
}
func (h *HandlerUser) UpdateUser(ctx context.Context, in *proto.User) (*proto.User, error) {

	UserIn := helper.BindProtoToUser(in)
	res, err := h.RepoUser.Update(ctx, UserIn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return helper.BindUser(res), err
}
func (h *HandlerUser) DeleteUser(ctx context.Context, in *proto.User) (*proto.User, error) {

	UserIn := helper.BindProtoToUser(in)
	err := h.RepoUser.Delete(ctx, UserIn.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &proto.User{}, err
}
