package vhost_test

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost"
	"os"
)

func ExampleDevEnv() {
	fmt.Println(vhost.IsDevEnv())
	os.Setenv(vhost.RuntimeEnvKey, "dev")
	fmt.Println(vhost.IsDevEnv())
	os.Setenv(vhost.RuntimeEnvKey, "devrrr")
	fmt.Println(vhost.IsDevEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleDevEnvOverride() {
	vhost.IsDevEnv = func() bool { return false }
	fmt.Println(vhost.IsDevEnv())
	os.Setenv(vhost.RuntimeEnvKey, "dev")
	fmt.Println(vhost.IsDevEnv())
	os.Setenv(vhost.RuntimeEnvKey, "devrrr")
	fmt.Println(vhost.IsDevEnv())

	// Output:
	// false
	// false
	// false
}

func ExampleProdEnv() {
	fmt.Println(vhost.IsProdEnv())
	os.Setenv(vhost.RuntimeEnvKey, "prod")
	fmt.Println(vhost.IsProdEnv())
	os.Setenv(vhost.RuntimeEnvKey, "production")
	fmt.Println(vhost.IsProdEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleReviewEnv() {
	fmt.Println(vhost.IsReviewEnv())
	os.Setenv(vhost.RuntimeEnvKey, "review")
	fmt.Println(vhost.IsReviewEnv())
	os.Setenv(vhost.RuntimeEnvKey, "revvrrr")
	fmt.Println(vhost.IsReviewEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleStageEnv() {
	fmt.Println(vhost.IsStageEnv())
	os.Setenv(vhost.RuntimeEnvKey, "stage")
	fmt.Println(vhost.IsStageEnv())
	os.Setenv(vhost.RuntimeEnvKey, "")
	fmt.Println(vhost.IsStageEnv())

	// Output:
	// false
	// true
	// false
}

func ExampleTestEnv() {
	fmt.Println(vhost.IsTestEnv())
	os.Setenv(vhost.RuntimeEnvKey, "test")
	fmt.Println(vhost.IsTestEnv())
	os.Setenv(vhost.RuntimeEnvKey, "atvrrr")
	fmt.Println(vhost.IsTestEnv())

	// Output:
	// false
	// true
	// false
}
