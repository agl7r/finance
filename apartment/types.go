package apartment

import (
	"encoding/json"
	"io/ioutil"
)

type Type struct {
	Id    int
	Title string
}

func GetTypes() []*Type {
	var types []*Type

	content, _ := ioutil.ReadFile("var/apartment/types.json")
	_ = json.Unmarshal(content, &types)

	return types
}
