package vhost

import (
	"os"
	"strings"

	"github.com/idiomatic-go/common-lib/vhost/usr"
)

//var IsDevEnv usr.EnvValid = func() bool {
//	return IsEnvMatch(usr.RUNTIME_ENV, usr.DEV_ENV)
//}

var IsReviewEnv usr.EnvValid = func() bool {
	return IsEnvMatch(usr.REVIEW_ENV)
}

var IsTestEnv usr.EnvValid = func() bool {
	return IsEnvMatch(usr.TEST_ENV)
}

var IsStageEnv usr.EnvValid = func() bool {
	return IsEnvMatch(usr.STAGE_ENV)
}

var IsProdEnv usr.EnvValid = func() bool {
	return IsEnvMatch(usr.PROD_ENV)
}

func IsEnvMatch(val string) bool {
	target := os.Getenv(usr.RuntimeEnvKey)
	if len(target) == 0 || !strings.EqualFold(target, val) {
		return false
	}
	return true
}

func init() {
	usr.IsDevEnv = func() bool {
		return IsEnvMatch(usr.DEV_ENV)
	}
}
