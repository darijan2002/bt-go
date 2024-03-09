package main

type BencDataType int

const (
	INT BencDataType = iota
	STRING
	LIST
	DICTIONARY
	UNKNOWN
)

// String - Creating common behavior - give the type a String function
func (bdt BencDataType) String() string {
	return [...]string{"INT", "STRING", "UNKNOWN"}[bdt]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (bdt BencDataType) EnumIndex() int {
	return int(bdt)
}
