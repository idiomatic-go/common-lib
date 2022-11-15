package vhost

import (
	"os"
	"strings"
)

var runtimeKey string = RuntimeEnvKey
var IsDevEnv FuncBool

func init() {
	IsDevEnv = func() bool {
		target := GetEnv()
		if len(target) == 0 || strings.EqualFold(target, DevEnv) {
			return true
		}
		return false
	}
}

// GetEnv - function to get the vhost runtime environment
func GetEnv() string {
	s := os.Getenv(runtimeKey)
	if s == "" {
		return DevEnv
	}
	return s
}

// SetEnv - function to set the vhost runtime environment
func SetEnv(s string) {
	os.Setenv(runtimeKey, s)
}

//var IsDevEnv EnvValid = func() bool {
//	return IsEnvMatch(RUNTIME_ENV, DEV_ENV)
//}

/*
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
	target := os.Getenv(runtimeKey)
	if len(target) == 0 || !strings.EqualFold(target, val) {
		return false
	}
	return true
}

*/
