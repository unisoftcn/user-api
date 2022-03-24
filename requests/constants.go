package requests

import "fmt"

func ServiceUrl(path string) string {
	//return fmt.Sprintf("http://user.park%s", path)
	return fmt.Sprintf("http://localhost:6090%s", path)
}
