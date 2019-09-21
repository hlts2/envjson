package envjson

import (
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	type input struct {
		paths []string
	}

	type want struct {
		isErr bool
		m     map[string]string
	}

	type test struct {
		input input
		want  want
	}

	tests := []test{
		{
			input: input{
				paths: []string{"./test/testdata/.envjson_1", "./test/testdata/.envjson_2"},
			},
			want: want{
				isErr: false,
				m: map[string]string{
					"debug": "true",
					"name":  "envjson",
				},
			},
		},

		{
			input: input{
				paths: []string{"./test/testdata/.envjson_invalid"},
			},
			want: want{
				isErr: true,
				m:     map[string]string{},
			},
		},

		{
			input: input{
				paths: []string{"./test/testdata/not_exists_file"},
			},
			want: want{
				isErr: true,
				m:     map[string]string{},
			},
		},
	}

	for i, test := range tests {
		err := Load(test.input.paths...)

		isErr := !(err == nil)

		if got, want := isErr, test.want.isErr; got != want {
			t.Errorf("tests[%d] - Load isErr is wrong, want: %v, but got: %v", i, want, got)
		}

		for k, v := range test.want.m {
			if got, want := os.Getenv(k), v; got != want {
				t.Errorf("tests[%d] - Load key: %v is wrong, want: %v, but got: %v", i, k, want, got)
			}
		}
	}
}

func TestSetEnv(t *testing.T) {
	input := map[string]interface{}{
		"debug": true,
		"db": map[string]string{
			"user":   "user_1",
			"pass":   "pass_1",
			"dbname": "dbname_1",
		},
	}

	if err := setEnv(input); err != nil {
		t.Fatalf("setEnv is err: %v", err)
	}

	if got, want := os.Getenv("debug"), "true"; want != got {
		t.Errorf("env debug is wrong, want: %v, but got: %v", want, got)
	}

	if got, want := os.Getenv("db"), `{"dbname":"dbname_1","pass":"pass_1","user":"user_1"}`; want != got {
		t.Errorf("env db is wrong, want: %v, but got: %v", want, got)
	}
}

func TestFilenamesOrDefault(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{},
			want:  []string{".envjson"},
		},
		{
			input: []string{"input.json"},
			want:  []string{"input.json"},
		},
	}

	for i, test := range tests {
		got := filenamesOrDefault(test.input...)

		if ok := reflect.DeepEqual(got, test.want); !ok {
			t.Errorf("tests[%d] - filenamesOrDefault is wrong, want: %v, but got: %v", i, test.want, got)
		}
	}
}
