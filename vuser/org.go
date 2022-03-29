package vuser

import (
	"strings"
	"vuuvv.cn/unisoftcn/orca/server"
)

func InAncestorOrg(token *server.AccessToken, orgPath string) bool {
	if token.IsSuper() {
		return true
	}
	return IsAncestorOrg(token.OrgPath, orgPath)
}

func IsAncestorOrg(ancestor string, descendant string) bool {
	return strings.HasPrefix(descendant, ancestor)
}
