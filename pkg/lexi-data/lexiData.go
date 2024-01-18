package lexidata

import (
	"fmt"
)

type LexiDataT int

const (
	String LexiDataT = iota
	Simple
	Int
	Double
	Array
    Error
)

type LexiData struct {
	Type LexiDataT
	Data LexiDataI
}

type LexiDataI interface {
	lexiData()
    Print()
}

type LexiSimple string

const (
	Ok   LexiSimple = "ok"
	None            = "none"
	Pong            = "pong"
)

func (LexiSimple) lexiData() {}

func (simple LexiSimple) Print() {
    fmt.Println(simple)
}

type LexiString struct {
	Str string
}

func (LexiString) lexiData() {}

func (s LexiString) Print() {
    fmt.Println(s.Str)
}

type LexiInt struct {
	Integer int64
}

func (LexiInt) lexiData() {}

func (i LexiInt) Print() {
    fmt.Println(i.Integer)
}

type LexiDouble struct {
	Double float64
}

func (LexiDouble) lexiData() {}

func (d LexiDouble) Print() {
    fmt.Println(d.Double)
}

type LexiArray struct {
	Array []*LexiData
}

func (LexiArray) lexiData() {}

func (arr LexiArray) Print() {
    for _, cur := range arr.Array {
        cur.Data.Print()
    }
}
