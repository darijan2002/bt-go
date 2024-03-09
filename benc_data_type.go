package main

type BencDataType int

const (
	BENC_INT BencDataType = iota
	BENC_STRING
	BENC_LIST
	BENC_DICTIONARY
	BENC_UNKNOWN
)

// String - Creating common behavior - give the type a String function
func (bdt BencDataType) String() string {
	return [...]string{"INT", "STRING", "UNKNOWN"}[bdt]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (bdt BencDataType) EnumIndex() int {
	return int(bdt)
}
