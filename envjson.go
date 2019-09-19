package envjson

import (
	"encoding/json"
	"os"
)

const DefaultFileName = ".envjson"

func Load(filenames ...string) error {
	filenames = filenamesOrDefault(filenames...)

	for _, filename := range filenames {
		if err := load(filename); err != nil {
			return err
		}
	}
	return nil
}

func load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var m map[string]interface{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		return err
	}

	if err := setEnv(m); err != nil {
		return err
	}

	return nil
}

func setEnv(m map[string]interface{}) error {
	for key, val := range m {
		if str, ok := val.(string); ok {
			if err := os.Setenv(key, str); err != nil {
				return err
			}
			continue
		}

		b, err := json.Marshal(val)
		if err != nil {
			return err
		}

		if err := os.Setenv(key, string(b)); err != nil {
			return err
		}
	}

	return nil
}

func filenamesOrDefault(filenames ...string) []string {
	if len(filenames) == 0 {
		return []string{DefaultFileName}
	}
	return filenames
}
