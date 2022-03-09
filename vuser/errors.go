package vuser

import (
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/request"
)

const (
	ErrorNotWxLogin = 20001
	ErrorCodeUserBindDuplicate = 20002
	ErrorPasswordMismatch = 20003
)

var (
	ErrorUserBindDuplicate = request.NewError(ErrorCodeUserBindDuplicate, "用户已绑定微信请解绑后再重新操作")
	ErrorPassword = request.NewError(ErrorPasswordMismatch, "用户名或密码错误")
)

func NewErrorPassword() error {
	return errors.WithStackAndSkip(ErrorPassword, 1)
}

//func ErrorNotLogin() error {
//	return errors.WithStack(&errors.Error{
//		Code:    errors.NotLogin,
//		Message: "请先登录",
//		NeedLog: false,
//	})
//}

//func ErrorIllegalAccess() error {
//	return errors.WithStack(&errors.Error{
//		Code:    errors.IllegalAccess,
//		Message: "无权访问",
//		NeedLog: false,
//	})
//}

//func ErrorIllegalUser() error {
//	return &request.Error{
//		Code:    errors.IllegalUser,
//		Message: "非法用户",
//		NeedLog: false,
//	}.WithStack()
//}

func ErrorWechatLogin() error {
	return &request.Error{
		Code:    ErrorNotWxLogin,
		Message: "请从微信客户端登录",
		NeedLog: false,
	}
}
