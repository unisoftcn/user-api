package model

type MqQrcode struct {
	FromUser string `json:"fromUser"`
	AppId    string `json:"appId"`
	Data     string `json:"data"`
}
