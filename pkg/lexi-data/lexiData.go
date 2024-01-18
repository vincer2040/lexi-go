package lexidata

type LexiDataT int

const (
	String LexiDataT = iota
	Simple
	Int
	Double
	Array
)

type LexiData struct {
	Type LexiDataT
	Data LexiDataI
}

type LexiDataI interface {
	lexiData()
}

type LexiSimple string

const (
	Ok   LexiSimple = "ok"
	None            = "none"
	Pong            = "pong"
)

func (LexiSimple) lexiData() {}

type LexiString struct {
	Str string
}

func (LexiString) lexiData() {}

type LexiInt struct {
	Integer int64
}

func (LexiInt) lexiData() {}

type LexiDouble struct {
	Double float64
}

func (LexiDouble) lexiData() {}

type LexiArray struct {
	Array []*LexiData
}

func (LexiArray) lexiData() {}
