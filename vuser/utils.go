package vuser

import (
	"go.uber.org/zap"
	"strings"
)

// Ancestors 由机构的path获取其所有父辈机构的path, 包含自己
func Ancestors(path string) []string {
	if path == "" {
		return []string{""}
	}

	parts := strings.Split(path, ":")
	if len(parts) <= 1 {
		return []string{"", path}
	}

	var ret []string
	cur := ""
	ret = append(ret, cur)

	for _, v := range parts {
		if cur == "" {
			cur = v
		} else {
			cur = cur + ":" + v
		}
		ret = append(ret, cur)
	}

	return ret
}

type UserCanLog interface {
	GetId() int64
	GetName() string
}

func LogAction(user UserCanLog, actionType string, data interface{}) {
	zap.L().Info(actionType, zap.Int64("userId", user.GetId()), zap.String("username", user.GetName()), zap.Any("data", data))
}
