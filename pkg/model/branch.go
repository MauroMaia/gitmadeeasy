package model

import "encoding/json"

type Branch struct {
	name    string
	isLocal bool
}

func NewBranch(name string, isLocal bool) Branch {
	return Branch{name, isLocal}
}

func (b Branch) GetName() string {
	return b.name
}

func (b Branch) IsLocal() bool {
	return b.isLocal
}

func (b Branch) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Name    string `json:"name"`
		IsLocal bool   `json:"isLocal"`
	}{
		b.name,
		b.isLocal,
	})
}
