package vuser

import "fmt"

const ServiceName = "vuser"

func ErrorCountKey(username string) string {
	return fmt.Sprintf("/login/error/count/%s", username)
}
