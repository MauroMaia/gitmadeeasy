package model

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
