package helper

import (
	"latihan_grpc/common/proto"
	"latihan_grpc/modules/model"
)

func BindUser(user model.User) *proto.User {
	return &proto.User{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Jenkel: user.Jenkel,
	}
}
func BindAllUser(users []model.User) *proto.DataUsers {
	var dataUser proto.DataUsers
	for _, v := range users {
		dataUser.Data = append(dataUser.Data, BindUser(v))
	}
	return &dataUser
}

func BindProtoToUser(user *proto.User) model.User {
	return model.User{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Jenkel: user.Jenkel,
	}
}
