package controller

import (
	"diary/model"
	"fmt"
	"regexp"
	"strconv"
)
//用户注册
func HandleRegister(user *model.User,backDoor string) model.Response  {
	var response model.Response
	response.Code = 0
	phone := user.Phone
	if phone == "" {
		response.Msg = "手机号不能为空"
		return response
	}
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	ok := rgx.MatchString(phone)
	back,err := strconv.ParseBool(backDoor)
	fmt.Println("ok",ok,"back",back,"err",err)
	if !ok && !back{
		response.Msg = "手机号格式不正确"
		return response
	}
	existUser := model.QueryUserByPhone(phone)
	if existUser != nil{
		response.Msg = "该账号已经注册"
		return response
	}
	fmt.Println("准备新增新用户",user)
	model.AddUser(user)
	newUser := model.QueryUserByPhone(user.Phone)
	fmt.Println("新增新用户成功",newUser)
	response.Msg = "新增新用户成功"
	response.Data = user
	return response
}

func HandleHello() model.Response {
	var response model.Response
	response.Code = 0
	response.Msg = "hello"
	return response
}
//用户查询
func HandleQueryUser(userid int) model.Response {
	var response model.Response
	response.Code = 0
	user := model.QueryUserById(userid)
	if user != nil{
		response.Msg = "查询成功"
	}else {
		response.Msg = "用户不存在"
	}
	response.Data =user
	return response
}
//用户删除
func HandleDeleteUser(userid int) model.Response {
	var response model.Response
	response.Code = 0
	user := model.DeleteUserById(userid)
	if user != nil{
		response.Msg = "删除成功"
	}else {
		response.Msg = "用户不存在"
	}
	return response
}

//用户登录
func HandleLogin(user *model.User,backDoor string) model.Response  {
	var response model.Response
	response.Code = 0
	phone := user.Phone
	if phone == "" {
		response.Msg = "手机号不能为空"
		return response
	}
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	ok := rgx.MatchString(phone)
	back,err := strconv.ParseBool(backDoor)
	fmt.Println("ok",ok,"back",back,"err",err)
	if !ok && !back{
		response.Msg = "手机号格式不正确"
		return response
	}
	existUser := model.QueryUserByPhone(phone)
	if existUser == nil{
		response.Msg = "用户不存在"
		return response
	}
	if user.Password != existUser.Password{
		response.Msg = "密码不正确"
		return response
	}
	response.Msg = "登录成功"
	response.Data = existUser
	return response
}

//用户列表查询
func HandleQueryUserList() model.Response {
	var response model.Response
	response.Code = 0
	userList := model.QueryUserListByWhere()
	if userList != nil{
		response.Msg = "查询成功"
	}else {
		response.Msg = "用户不存在"
	}
	response.Data = userList
	return response
}
