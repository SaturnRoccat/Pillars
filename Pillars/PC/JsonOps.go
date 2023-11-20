package PC

import (
	"encoding/json"
	"os"
)

func PullJson[T interface{}](PTJ string) (T, error) {
	var config T
	var jsonData []byte
	var err error

	jsonData, err = os.ReadFile(PTJ)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func WriteEmptyJson(PTJ string) error {

	return os.WriteFile(PTJ, []byte("{}"), 0777)
}
