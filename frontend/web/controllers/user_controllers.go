package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"goproject/commons"
	"goproject/datamodels"
	"goproject/encrypt"
	"goproject/services"
	"strconv"
)

type UserController struct {
	Context iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Context.FormValue("nickName")
		userName = c.Context.FormValue("userName")
		password = c.Context.FormValue("password")
	)

	user := &datamodels.User{
		NickName:     nickName,
		UserName:     userName,
		HashPassword: password,
	}

	_, err := c.Service.AddUser(user)
	if err != nil {
		c.Context.Redirect("/user/error")
		return
	}
	c.Context.Redirect("/user/login")
	return
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	//1.获取用户提交的表单信息
	var (
		userName = c.Context.FormValue("userName")
		password = c.Context.FormValue("password")
	)

	//2、验证账号密码正确
	user, isOK := c.Service.IsPwdSuccess(userName, password)
	if !isOK {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	//3、写入用户 ID 到 cookie 中
	userIDString := strconv.FormatInt(user.ID, 10)
	commons.GlobalCookie(c.Context, "uid", userIDString)
	userIDByte := []byte(userIDString)
	userIDStringEncoded, err := encrypt.EncodePassword(userIDByte)
	if err != nil {
		c.Context.Application().Logger().Error(err)
	}
	//4、写入加密 Cookie 到用户浏览器中
	commons.GlobalCookie(c.Context, "sign", userIDStringEncoded)
	//c.Session.Set("userID", strconv.FormatInt(user.ID, 10))
	return mvc.Response{
		Path: "/product/",
	}
}
