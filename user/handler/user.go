package handler

import (
	"context"
	"github.com/machinism1011/microservice/user/domain/model"
	"github.com/machinism1011/microservice/user/domain/service"
	protoUser "github.com/machinism1011/microservice/user/proto/user"
)

type User struct{
	UserDataService service.IUserDataService
}

// register
func (u *User) Register(ctx context.Context, userRegisterRequest *protoUser.UserRegisterRequest, userRegisterResponse *protoUser.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}

	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, userLoginRequest *protoUser.UserLoginRequest, userLoginResponse *protoUser.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *protoUser.UserInfoRequest, userInfoResponse *protoUser.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = UserForResponse(userInfo)
	return nil
}

func UserForResponse(userModel *model.User) *protoUser.UserInfoResponse {
	response := &protoUser.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}

