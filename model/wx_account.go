package model

import (
	"fmt"
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/orm"
	"github.com/vuuvv/orca/utils"

	url2 "net/url"
)

type WxAccount struct {
	orm.Id
	Name           string     `json:"name"`
	OrgPath        string     `json:"orgPath"`
	Type           string     `json:"type"`
	AppId          string     `json:"appId"`
	AppSecret      string     `json:"appSecret"`
	Token          string     `json:"token"`
	EncodingAESKey string     `json:"encodingAesKey"`
	OauthAddress   string     `json:"oauthAddress"`
	ClientAddress  string     `json:"clientAddress"`
	Custom         bool       `json:"custom"`
	Template       map[string]string `json:"template"`
	Site           *Site
}

func (a *WxAccount) GetOauthUrl() (string, error) {
	url, err := a.Site.Url()
	if err != nil {
		return "", err
	}
	return utils.UrlJoin(url, a.OauthAddress, fmt.Sprintf("/oauth?app=%s", a.AppId))
}

// GetEntryUrl 获取微信入口Url
func (a *WxAccount) GetEntryUrl(redirect string, token string) (string, error) {
	url, err := a.Site.Url()
	if err != nil {
		return "", err
	}
	params := url2.Values{}
	params.Add("app", a.AppId)
	if redirect != "" {
		params.Add("r", redirect)
	}
	if token != "" {
		params.Add("token", token)
	}

	return utils.UrlJoin(url, a.OauthAddress, fmt.Sprintf("/?%s", params.Encode()))
}

func (a *WxAccount) GetClientUrl() (string, error) {
	url, err := a.Site.Url()
	if err != nil {
		return "", err
	}
	return utils.UrlJoin(url, a.ClientAddress)
}

func (a *WxAccount) GetTemplateId(name string) (id string, err error) {
	if a.Template != nil {
		if id, ok := a.Template[name]; ok {
			return id, nil
		}
	}
	return "", errors.New(fmt.Sprintf("模板[%s]未设置", name))
}
