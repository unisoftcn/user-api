package vuser

import (
	"fmt"
)

const CookieCaptchaId = "CaptchaId"
const CookieWxRedirectUrl = "WxRedirectUrl"
const CookieTrdUserId = "TrdUserId"

func SessionKeyCaptcha(captchaId string) string {
	return fmt.Sprintf("/session/captcha/%s", captchaId)
}
