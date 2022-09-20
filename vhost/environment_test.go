package vhost_test

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/idiomatic-go/common-lib/vhost/usr"
	"os"
)

func ExampleDevEnv() {
	fmt.Println(usr.IsDevEnv())
	os.Setenv(usr.RuntimeEnvKey, "dev")
	fmt.Println(usr.IsDevEnv())
	os.Setenv(usr.RuntimeEnvKey, "devrrr")
	fmt.Println(usr.IsDevEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleDevEnvOverride() {
	usr.IsDevEnv = func() bool { return false }
	fmt.Println(usr.IsDevEnv())
	os.Setenv(usr.RuntimeEnvKey, "dev")
	fmt.Println(usr.IsDevEnv())
	os.Setenv(usr.RuntimeEnvKey, "devrrr")
	fmt.Println(usr.IsDevEnv())

	// Output:
	// false
	// false
	// false
}

func ExampleProdEnv() {
	fmt.Println(vhost.IsProdEnv())
	os.Setenv(usr.RuntimeEnvKey, "prod")
	fmt.Println(vhost.IsProdEnv())
	os.Setenv(usr.RuntimeEnvKey, "production")
	fmt.Println(vhost.IsProdEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleReviewEnv() {
	fmt.Println(vhost.IsReviewEnv())
	os.Setenv(usr.RuntimeEnvKey, "review")
	fmt.Println(vhost.IsReviewEnv())
	os.Setenv(usr.RuntimeEnvKey, "revvrrr")
	fmt.Println(vhost.IsReviewEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleStageEnv() {
	fmt.Println(vhost.IsStageEnv())
	os.Setenv(usr.RuntimeEnvKey, "stage")
	fmt.Println(vhost.IsStageEnv())
	os.Setenv(usr.RuntimeEnvKey, "")
	fmt.Println(vhost.IsStageEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleTestEnv() {
	fmt.Println(vhost.IsTestEnv())
	os.Setenv(usr.RuntimeEnvKey, "test")
	fmt.Println(vhost.IsTestEnv())
	os.Setenv(usr.RuntimeEnvKey, "atvrrr")
	fmt.Println(vhost.IsTestEnv())

	// Output:
	// false
	// true
	// false
}
