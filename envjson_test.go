package envjson

import (
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {

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
