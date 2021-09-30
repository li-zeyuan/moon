package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/li-zeyuan/micro/moon.common.api/errorenum"
	"github.com/li-zeyuan/micro/moon.common.api/sequence"
	"github.com/li-zeyuan/micro/moon.common.api/utils"
	"github.com/li-zeyuan/micro/user.logic.http/app/dao"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/app/model/inner"
	"github.com/li-zeyuan/micro/user.logic.http/config"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
	"gorm.io/gorm"
)

const (
	baseWXSessionUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

var Login = loginService{}

type loginService struct{}

func (l *loginService) VerifySingUp(infra *middleware.Infra, req *model.LoginApiPhoneLoginReq) error {
	if isLetterOrDigit, _ := regexp.MatchString(`^[A-Za-z0-9]{1,16}$`, req.Passport); !isLetterOrDigit {
		return errorenum.ErrorPassportLetterOrDigit
	}

	if len(req.Password) > 8 || len(req.Password) == 0 || req.Password != req.Password2 {
		return errorenum.ErrorPasswordInConformity
	}
	if isLetterOrDigit, _ := regexp.MatchString(`^[A-Za-z0-9]{1,16}$`, req.Password); !isLetterOrDigit {
		return errorenum.ErrorPasswordLetterOrDigit
	}

	//// 判断账号是否存在
	//profileMap, err := userdbrpc.GetProfileByPassport(infra.BaseInfra, []string{req.Passport})
	//if err != nil {
	//	return err
	//}
	//if len(profileMap) > 0 {
	//	return errorenum.ErrorPassportExist
	//}

	return nil
}

func (l *loginService) SingUp(infra *middleware.Infra, req *model.LoginApiPhoneLoginReq) error {

	return nil
}

func (l *loginService) WeChatLogin(infra *middleware.Infra, req *model.LoginApiWeChatLoginReq) (string, error) {
	wxSession, err := queryWxSession(infra, req.Code)
	if err != nil {
		return "", err
	}

	userDao := dao.NewUser(infra.DB)
	userProfile, err := userDao.GetByOpenid(infra, wxSession.OpenId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	if err == gorm.ErrRecordNotFound {
		userProfile = new(inner.UserProfileModel)
		userProfile.Uid = sequence.NewID()
		userProfile.Name = fmt.Sprintf("家普_%d", userProfile.Uid%1000000)
		userProfile.Gender = inner.GenderMan
		userProfile.Openid = wxSession.OpenId
		userProfile.SessionKey = wxSession.SessionKey
		err = userDao.Save(infra, []*inner.UserProfileModel{userProfile})
		if err != nil {
			return "", err
		}
	}

	token, err := utils.GenerateToken(userProfile.Uid, utils.RoleUser, time.Hour*24*30)
	if err != nil {
		return "", err
	}

	return token, nil
}

func queryWxSession(infra *middleware.Infra, code string) (*model.WXSessionRet, error) {
	url := fmt.Sprintf(baseWXSessionUrl, config.Conf.Wechat.AppId, config.Conf.Wechat.Secret, code)
	resp, err := http.Get(url)
	if err != nil {
		infra.Log.Error("get weChat session error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := model.WXSessionRet{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode:%d  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg)
	}

	return &wxResp, nil
}
