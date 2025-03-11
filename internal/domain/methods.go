package domain

const (
	SET Method = "SET"
	GET Method = "GET"
	DEL Method = "DEL"
)

const (
	SetLen = 3
	GetLen = 2
	DelLen = 2
)

type Query struct {
	Method Method
	Args   []string
}

type Method string
