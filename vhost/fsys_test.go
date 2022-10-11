package vhost

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed resource/*
var content embed.FS

func init() {
	MountFS(content)
}

func _ExampleFileSystemNotMounted() {
	_, err := ReadFile("resource/readme.txt")
	fmt.Printf("Error : %v\n", err)

	//Output:
	// Error : invalid argument : file system has not been mounted
}

func _ExampleReadFile() {
	_, err0 := ReadFile("")
	fmt.Printf("Error : %v\n", err0)

	buf, err := ReadFile("bad-path/config_bad.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Println(string(buf))
	}

	buf, err = ReadFile("postgresql/config_{env}.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Println(string(buf))
	}

	// Should override and return config_test.txt
	lookupEnv = func(name string) (string, error) { return "stage", nil }
	buf, err = ReadFile("postgresql/config_{env}.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Println(string(buf))
	}

	//Output:
	// Error : invalid argument : path is empty
	// Error : open resource/bad-path/config_bad.txt: file does not exist
	// Error : invalid argument : template variable is invalid: env
	// env : test
}

func ExampleReadMap() {
	_, err0 := ReadMap("")
	fmt.Printf("Error : %v\n", err0)

	m, err := ReadMap("postgresql/config_dev.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Printf("Map [config_dev.txt]: %v\n", m)
	}

	m, err = ReadMap("postgresql/config_test.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Printf("Map [config_test.txt]: %v\n", m)
	}

	// Should override and return config_test.txt
	lookupEnv = func(name string) (string, error) { return "stage", nil }
	m, err = ReadMap("postgresql/config_{env}.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Printf("Map : %v\n", m)
	}

	//Output:
	// Error : invalid argument : path is empty
	// Error : open resource/bad-path/config_bad.txt: file does not exist
	// Error : invalid argument : template variable is invalid: env
	// env : test
}

func _TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"BlankLine", args{line: ""}, "", "", false},
		{"LeadingSpace", args{line: " "}, "", "", false},
		{"LeadingSpaces", args{line: "       "}, "", "", false},

		{"Comment", args{line: comment}, "", "", false},
		{"LeadingSpaceComment", args{line: " " + comment}, "", "", false},
		{"LeadingSpacesComment", args{line: "       " + comment}, "", "", false},

		{"MissingDelimiter", args{line: "missing delimiter"}, "", "", true},

		{"KeyOnly", args{line: "key-only :"}, "key-only", "", false},
		{"KeyValue", args{line: "key  : value"}, "key", "value", false},
		{"KeyValueLeadingSpaces", args{line: "key:      value"}, "key", "value", false},
		{"KeyValueTrailingSpaces", args{line: "key :value    "}, "key", "value    ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseLine() got = [%v], want [%v]", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseLine() got1 = [%v], want [%v]", got1, tt.want1)
			}
		})
	}
}
