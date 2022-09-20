package vhost

import (
	"fmt"
	"testing"
)

func ExampleInvalidName() {
	_, err := ReadFile("")
	fmt.Printf("Error : %v\n", err)

	//Output:
	// Error : invalid argument : file name is empty

}

func ExampleFileSystemNotMounted() {
	_, err := ReadFile("resource/readme.txt")
	fmt.Printf("Error : %v\n", err)

	//Output:
	// Error : invalid argument : file system has not been mounted

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

		{"MissingDelimiter", args{line: "missing:delimiter"}, "", "", true},

		{"KeyOnly", args{line: "key:only "}, "key:only", "", false},
		{"KeyValue", args{line: "key value"}, "key", "value", false},
		{"KeyValueLeadingSpaces", args{line: "key      value"}, "key", "value", false},
		{"KeyValueTrailingSpaces", args{line: "key value    "}, "key", "value    ", false},
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
