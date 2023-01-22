package apartment

import (
	"encoding/json"
	"io/ioutil"
)

type Type struct {
	Id    int `json:"id"`
	Title string `json:"title"`
}

func (_type *Type) String() string {
	return _type.Title
}

type Types []*Type

func (types *Types) GetById(id int) *Type {
	for _, t := range *types {
		if t.Id == id {
			return t
		}
	}
	return nil
}

func GetTypes() *Types {
	types := new(Types)

	content, _ := ioutil.ReadFile("var/apartment/types.json")
	_ = json.Unmarshal(content, &types)

	return types
}
