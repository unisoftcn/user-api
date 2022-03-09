package vuser

import (
	"fmt"
)

const CookieCaptchaId = "CaptchaId"

func SessionKeyCaptcha(captchaId string) string {
	return fmt.Sprintf("/session/captcha/%s", captchaId)
}
