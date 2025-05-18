package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	isJsonExtension := strings.HasSuffix(path, ".json")
	if !isJsonExtension {
		return nil, errors.New(fmt.Sprintf("file %s is not json extension", path))
	}
	isValidJson := json.Valid(file)
	if !isValidJson {
		return nil, errors.New(fmt.Sprintf("file %s is not valid json", path))
	}

	return file, nil
}
