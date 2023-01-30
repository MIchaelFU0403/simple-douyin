package core

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"user/model"
	"user/services"
	"user/utils"
)

/**
用户登录service层
*/
func (*UserService) Login(ctx context.Context, req *services.DouyinUserLoginRequest, resp *services.DouyinUserLoginResponse) error {
	username := req.Username
	password := req.Password
	user, err := model.NewUserDaoInstance().FindUserByName(username)
	if err != nil {
		panic("调用UserDao的FindUserByName方法，根据用户名查询User失败")
		return err
	}

	if utils.Sha256(password) != user.Password {
		fmt.Println("用户名或密码错误")
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码错误"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "登录成功"
	resp.UserId = user.UserId
	resp.Token = "abcabc"
	return nil
}

/**
注册 service层
*/
func (*UserService) Register(ctx context.Context, req *services.DouyinUserRegisterRequest, resp *services.DouyinUserRegisterResponse) error {
	//查询用户名，没有错误（能查到）
	username := req.Username
	password := req.Password
	if _, err := model.NewUserDaoInstance().FindUserByName(username); err == nil {
		fmt.Println("用户名已经存在")
		resp.StatusCode = -1
		resp.StatusMsg = "用户名已经存在"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	user := &model.User{
		Name:           username,
		Password:       utils.Sha256(password),
		FollowingCount: 0,
		FollowerCount:  0,
		CreateAt:       time.Now(),
	}

	_, err := model.NewUserDaoInstance().CreateUser(user)
	if err != nil {
		fmt.Println("创建用户失败")
		resp.StatusCode = -1
		resp.StatusMsg = "创建用户失败"
		resp.UserId = -1
		resp.Token = ""
		return err
	}

	user, _ = model.NewUserDaoInstance().FindUserByName(username)

	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	resp.UserId = user.UserId
	resp.Token = ""
	return nil
}

/**
登录用户的详细信息 service层
*/
func (*UserService) UserInfo(ctx context.Context, req *services.DouyinUserRequest, resp *services.DouyinUserResponse) error {
	userId := req.UserId     //传入的参数
	tokenUserId := req.Token //token解析出来的userId

	//curUserId, err := strconv.ParseInt(currentUserId, 10, 64) //当前用户的id
	tokenUserIdConv, err := strconv.ParseInt(tokenUserId, 10, 64) //当前用户的id
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "类型转换失败"
		resp.User = &services.User{}
		return err
	}

	if userId != tokenUserIdConv {
		resp.StatusCode = 1
		resp.StatusMsg = "认证失败，请重新登陆"
		resp.User = &services.User{}
		return nil
	}

	//2. 根据userId查询User
	user, err := model.NewUserDaoInstance().FindUserById(userId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "查找用户信息时发生异常"
		resp.User = &services.User{}
		return err
	}

	isFollow := false

	resp.StatusCode = 0
	resp.StatusMsg = "查询用户信息成功"
	resp.User = BuildProtoUser(user, isFollow)
	return nil
}

func BuildProtoUser(item *model.User, isFollow bool) *services.User {
	user := services.User{
		Id:            item.UserId,
		Name:          item.Name,
		FollowCount:   item.FollowingCount,
		FollowerCount: item.FollowerCount,
		IsFollow:      isFollow,
	}
	return &user
}
