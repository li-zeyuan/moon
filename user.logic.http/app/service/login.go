package service

import (
	"regexp"

	"github.com/li-zeyuan/micro/micro.common.api/errorenum"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
)

var Login = loginService{}

type loginService struct{}

func (l *loginService) VerifySingUp(req *model.LoginApiSingUpReq) error {
	if len(req.Passport) > 8 || len(req.Passport) == 0 {
		return errorenum.ErrorPassportLength
	}
	if isLetterOrDigit, _ := regexp.MatchString(`^[A-Za-z0-9]$`, req.Passport); !isLetterOrDigit {
		return errorenum.ErrorPassportLetterOrDigit
	}

	if len(req.Password) > 8 || len(req.Password) == 0 || req.Password != req.Password2 {
		return errorenum.ErrorPasswordLength
	}
	if isLetterOrDigit, _ := regexp.MatchString(`^[A-Za-z0-9]$`, req.Password); !isLetterOrDigit {
		return errorenum.ErrorPasswordLetterOrDigit
	}

	// todo 判断账号是否存在

	return nil
}

func (l *loginService) SingUp() error {
	return nil
}
