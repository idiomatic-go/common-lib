package vhost

import (
	"os"
	"strings"
)

func GetEnv() string {
	return os.Getenv(RuntimeEnvKey)
}

//var IsDevEnv EnvValid = func() bool {
//	return IsEnvMatch(RUNTIME_ENV, DEV_ENV)
//}

var IsReviewEnv EnvValid = func() bool {
	return IsEnvMatch(REVIEW_ENV)
}

var IsTestEnv EnvValid = func() bool {
	return IsEnvMatch(TEST_ENV)
}

var IsStageEnv EnvValid = func() bool {
	return IsEnvMatch(STAGE_ENV)
}

var IsProdEnv EnvValid = func() bool {
	return IsEnvMatch(PROD_ENV)
}

func IsEnvMatch(val string) bool {
	target := os.Getenv(RuntimeEnvKey)
	if len(target) == 0 || !strings.EqualFold(target, val) {
		return false
	}
	return true
}

func init() {
	IsDevEnv = func() bool {
		return IsEnvMatch(DEV_ENV)
	}
}
