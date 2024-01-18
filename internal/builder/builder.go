package builder

import (
	"strconv"

	"github.com/vincer2040/lexi-go/internal/util"
)

type Builder struct {
	buf []byte
}

func New() Builder {
    buf := []byte{}
    return Builder{buf}
}

func (b *Builder) AddPing() *Builder {
    b.buf = append(b.buf, util.SIMPLE_TYPE_BYTE)
    b.buf = append(b.buf, 'P')
    b.buf = append(b.buf, 'I')
    b.buf = append(b.buf, 'N')
    b.buf = append(b.buf, 'G')
    b.addEnd()
    return b
}

func (b *Builder) AddArray(length int) *Builder {
	b.buf = append(b.buf, util.ARRAY_TYPE_BYTE)
	b.addLength(length)
	b.addEnd()
	return b
}

func (b *Builder) AddString(str string) *Builder {
	b.buf = append(b.buf, util.STRING_TYPE_BYTE)
	b.addLength(len(str))
	b.addEnd()

	for _, ch := range str {
		b.buf = append(b.buf, byte(ch))
	}

	b.addEnd()

	return b
}

func (b *Builder) AddInt(integer int) *Builder {
    b.buf = append(b.buf, util.INT_TYPE_BYTE)
    int_str := strconv.Itoa(integer)
    for _, ch := range int_str {
        b.buf = append(b.buf, byte(ch))
    }
    b.addEnd()
    return b
}

func (b *Builder) AddDouble(dbl float64) *Builder {
    b.buf = append(b.buf, util.DOUBLE_TYPE_BYTE)
    dbl_str := strconv.FormatFloat(dbl, 'f', -1, 64)
    for _, ch := range dbl_str {
        b.buf = append(b.buf, byte(ch))
    }
    b.addEnd()
    return b
}

func (b *Builder) addLength(length int) {
	lengthString := strconv.Itoa(length)
	for _, ch := range lengthString {
		b.buf = append(b.buf, byte(ch))
	}
}

func (b *Builder) addEnd() {
	b.buf = append(b.buf, '\r')
	b.buf = append(b.buf, '\n')
}
