package util

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed resource/*
var fsys embed.FS

func _ExampleReadFile() {
	_, err0 := FSReadFile(nil, "")
	fmt.Printf("Error : %v\n", err0)

	buf, err := FSReadFile(fsys, "bad-path/config_bad.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Println(string(buf))
	}

	buf, err = FSReadFile(fsys, "postgresql/config_dev.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Println(string(buf))
	}

	//Output:
	// Error : invalid argument : path is empty
	// Error : open resource/bad-path/config_bad.txt: file does not exist
	// this is the test environment
	// env : dev
	// next  : second value
	// timeout : 10020

}

func _ExampleReadMap() {
	_, err0 := FSReadMap(fsys, "")
	fmt.Printf("Error : %v\n", err0)

	m, err := FSReadMap(fsys, "postgresql/config_dev.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Printf("Map [config_dev.txt]: %v\n", m)
	}

	m, err = FSReadMap(fsys, "postgresql/config_test.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Printf("Map [config_test.txt]: %v\n", m)
	}

	// Should override and return config_test.txt
	//lookupEnv = func(name string) (string, error) { return "stage", nil }
	//m, err = ReadMap("postgresql/config_{env}.txt")
	//if err != nil {
	//	fmt.Printf("Error : %v\n", err)
	//} else {
	//	fmt.Printf("Map : %v\n", m)
	//}

	//Output:
	// Error : invalid argument : path is empty
	// Map [config_dev.txt]: map[env:dev
	//  next:second value
	//  timeout:10020]
	// Map [config_test.txt]: map[env:test
	//  thelast:line of the file]

}

func TestParseLine(t *testing.T) {
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
