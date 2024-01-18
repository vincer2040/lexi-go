package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/vincer2040/lexi-go/internal/util"
	lexigo "github.com/vincer2040/lexi-go/pkg/lexi-go"
)

type Parser struct {
	input    []byte
	inputLen int
	pos      int
	ch       byte
}

func New(input []byte, input_len int) Parser {
	p := Parser{input, input_len, 0, 0}
	p.readByte()
	return p
}

func (p *Parser) Parse() (*lexigo.LexiData, error) {
	return p.parseData()
}

func (p *Parser) parseData() (*lexigo.LexiData, error) {
	switch p.ch {
	case util.STRING_TYPE_BYTE:
		return p.parseString()
	case util.INT_TYPE_BYTE:
		return p.parseInt()
	case util.DOUBLE_TYPE_BYTE:
		return p.parseDouble()
	case util.ARRAY_TYPE_BYTE:
		return p.parseArray()
	default:
		break
	}
	return nil, errors.New("unknown data type byte")
}

func (p *Parser) parseString() (*lexigo.LexiData, error) {
	if !p.expectPeekToBeNum() {
		return nil, errors.New("expected length")
	}

	length := p.parseLength()

	if !p.curByteIs('\r') {
		return nil, errors.New("expected ret car")
	}

	if !p.expectPeek('\n') {
		return nil, errors.New("expected new line")
	}

	p.readByte()

	var builder strings.Builder

	for i := 0; i < length; i++ {
		builder.WriteByte(p.ch)
		p.readByte()
	}

	if !p.curByteIs('\r') {
		return nil, errors.New("expected ret car")
	}

	if !p.expectPeek('\n') {
		return nil, errors.New("expected new line")
	}

	p.readByte()

	str := builder.String()

	res := &lexigo.LexiData{
		Type: lexigo.String,
		Data: lexigo.LexiString{Str: str},
	}

	return res, nil
}

func (p *Parser) parseInt() (*lexigo.LexiData, error) {
	p.readByte()
	var builder strings.Builder
	for p.ch != '\r' && p.ch != 0 {
		builder.WriteByte(p.ch)
		p.readByte()
	}

	i, err := strconv.ParseInt(builder.String(), 10, 64)
	if err != nil {
		return nil, err
	}

	if !p.curByteIs('\r') {
		return nil, errors.New("expected ret car")
	}

	if !p.expectPeek('\n') {
		return nil, errors.New("expected new line")
	}

	p.readByte()

	res := &lexigo.LexiData{
		Type: lexigo.Int,
		Data: lexigo.LexiInt{Integer: i},
	}

	return res, nil
}

func (p *Parser) parseDouble() (*lexigo.LexiData, error) {
	p.readByte()
	var builder strings.Builder
	for p.ch != '\r' && p.ch != 0 {
		builder.WriteByte(p.ch)
		p.readByte()
	}

	if !p.curByteIs('\r') {
		return nil, errors.New("expected ret car")
	}

	if !p.expectPeek('\n') {
		return nil, errors.New("expected new line")
	}

	dbl, err := strconv.ParseFloat(builder.String(), 64)

	if err != nil {
		return nil, err
	}

	p.readByte()

	res := &lexigo.LexiData{
		Type: lexigo.Double,
		Data: lexigo.LexiDouble{Double: dbl},
	}

	return res, nil
}

func (p *Parser) parseArray() (*lexigo.LexiData, error) {
	if !p.expectPeekToBeNum() {
		return nil, errors.New("expected length")
	}
	length := p.parseLength()

	if !p.curByteIs('\r') {
		return nil, errors.New("expected ret car")
	}

	if !p.expectPeek('\n') {
		return nil, errors.New("expected new line")
	}

	p.readByte()

	var data []*lexigo.LexiData

	for i := 0; i < length; i++ {
		parsed, err := p.parseData()
		if err != nil {
			return nil, err
		}

		data = append(data, parsed)
	}

	res := &lexigo.LexiData{
		Type: lexigo.Array,
		Data: lexigo.LexiArray{Array: data},
	}

	return res, nil
}

func (p *Parser) parseLength() int {
	res := 0
	for p.ch != '\r' && p.ch != 0 {
		res = (res * 10) + (int(p.ch) - int('0'))
		p.readByte()
	}
	return res
}

func (p *Parser) peekByte() byte {
	if p.pos >= p.inputLen {
		return 0
	}
	return p.input[p.pos]
}

func (p *Parser) curByteIs(ch byte) bool {
	return p.ch == ch
}

func (p *Parser) peekByteIs(ch byte) bool {
	return p.peekByte() == ch
}

func (p *Parser) expectPeek(ch byte) bool {
	if !p.peekByteIs(ch) {
		return false
	}
	p.readByte()
	return true
}

func (p *Parser) expectPeekToBeNum() bool {
	peek := p.peekByte()
	if '0' <= peek && peek <= '9' {
		p.readByte()
		return true
	}
	return false
}

func (p *Parser) readByte() {
	if p.pos >= p.inputLen {
		p.ch = 0
		return
	}
	p.ch = p.input[p.pos]
	p.pos++
}
