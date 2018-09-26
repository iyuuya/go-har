package har

import (
	"encoding/json"
	"io/ioutil"

	"github.com/iyuuya/go-har/internal/types"
)

func LoadHARFromFile(fileName string) (*types.HAR, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return LoadHARFromBytes(data)
}

func LoadHARFromStr(raw string) (*types.HAR, error) {
	return LoadHARFromBytes([]byte(raw))
}

func LoadHARFromBytes(bytes []byte) (*types.HAR, error) {
	tmp := new(types.HAR)

	err := json.Unmarshal(bytes, tmp)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}
