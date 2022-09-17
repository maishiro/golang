package main

import (
	"os"
	"testing"
)

func TestLoadFileNone(t *testing.T) {
	dir, _ := os.Getwd()
	filename := dir + "/exampleNone.txt"
	contents, err := LoadFile(filename)
	if contents != "" {
		t.Fatalf(`Contents is not "" but it is %s`, contents)
	}
	if err == nil {
		t.Fatal(`Content file load  why success: ??`)
	}
}

func TestLoadFileExist(t *testing.T) {
	dir, _ := os.Getwd()
	filename := dir + "/exampleExist.txt"

	contentExpect := "abc"
	file, _ := os.Create(filename)
	defer func() {
		file.Close()
		os.Remove(filename)
	}()

	if file != nil {
		file.WriteString(contentExpect)
	}

	contents, err := LoadFile(filename)
	if contents != contentExpect {
		t.Fatalf(`Contents is not "%s" but it is %s`, contentExpect, contents)
	}
	if err != nil {
		t.Fatalf(`Content file load failed: %s`, err.Error())
	}
}

// add from IDE tool
func TestLoadFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "file none",
			args:    args{path: "noneFileName.txt"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LoadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
