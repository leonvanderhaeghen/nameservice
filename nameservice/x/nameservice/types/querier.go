package types

import "strings"


const QueryListName = "list-name"
const QueryGetName = "get-name"
const QueryResolveName = "resolve-name"

type QueryResResolve struct{
	Value string `json:"value"`
}

func (r QueryResResolve) String() string{
	return r.Value
}

type QueryResNames []string

func (n QueryResNames) String() string{
	return strings.Join(n[:],"\n")
}