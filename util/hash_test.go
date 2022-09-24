package util

import "fmt"

func ExampleEmptyStrng() {
	hash1 := SimpleHash("")
	hash2 := SimpleHash("")
	fmt.Printf("Hash1 [] : %v\n", hash1)
	fmt.Printf("Hash2 [] : %v\n", hash2)
	fmt.Printf("Hash equality : %v\n", hash1 == hash2)

	//Output:
	// Hash1 [] : 0
	// Hash2 [] : 0
	// Hash equality : true
}

func ExampleLowEntropy() {
	hash1 := SimpleHash("1")
	hash2 := SimpleHash("1")
	fmt.Printf("Hash1 [1] : %v\n", hash1)
	fmt.Printf("Hash2 [1] : %v\n", hash2)
	fmt.Printf("Hash equality : %v\n", hash1 == hash2)

	hash1 = SimpleHash("1")
	hash2 = SimpleHash("2")
	fmt.Printf("Hash1 [1] : %v\n", hash1)
	fmt.Printf("Hash2 [2] : %v\n", hash2)
	fmt.Printf("Hash equality : %v\n", hash1 == hash2)

	//Output:
	// Hash1 [1] : 873244444
	// Hash2 [1] : 873244444
	// Hash equality : true
	// Hash1 [1] : 873244444
	// Hash2 [2] : 923577301
	// Hash equality : false
}

func ExampleHighEntropy() {
	s1 := "sdff9b99ae5nfhei446fadmnzacsxgdfo7k%^&$#@*kju(+={:>}/?'"
	s2 := "cvdfr63%$12#_+[:;'}bbmtfuoipcx2&6^$sd=;,<.>//??iop|~"
	hash1 := SimpleHash(s1)
	hash2 := SimpleHash(s1)
	fmt.Printf("Hash1 [%v] : %v\n", s1, hash1)
	fmt.Printf("Hash2 [%v] : %v\n", s1, hash2)
	fmt.Printf("Hash equality : %v\n", hash1 == hash2)

	hash1 = SimpleHash(s1)
	hash2 = SimpleHash(s2)
	fmt.Printf("Hash1 [%v] : %v\n", s1, hash1)
	fmt.Printf("Hash2 [%v] : %v\n", s2, hash2)
	fmt.Printf("Hash equality : %v\n", hash1 == hash2)

	//Output:
	// Hash1 [sdff9b99ae5nfhei446fadmnzacsxgdfo7k%^&$#@*kju(+={:>}/?'] : 3506137036
	// Hash2 [sdff9b99ae5nfhei446fadmnzacsxgdfo7k%^&$#@*kju(+={:>}/?'] : 3506137036
	// Hash equality : true
	// Hash1 [sdff9b99ae5nfhei446fadmnzacsxgdfo7k%^&$#@*kju(+={:>}/?'] : 3506137036
	// Hash2 [cvdfr63%$12#_+[:;'}bbmtfuoipcx2&6^$sd=;,<.>//??iop|~] : 3682702230
	// Hash equality : false

}
