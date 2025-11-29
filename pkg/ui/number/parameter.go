package number

import "fmt"

type Parameter struct {
	// Immutable
	name          string
	totalBitWidth int
	parts         int
	converter     Converter
	// Mutable
	data []byte
}

func NewUintParameter(totalBitWidth, partBitWidth, base int) *Parameter {
	return NewNamedUintParameter("", totalBitWidth, partBitWidth, base)
}

func NewNamedUintParameter(name string, totalBitWidth, partBitWidth, base int) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)
	converter := NewUintConverter(partBitWidth, base)

	return &Parameter{
		name:          name,
		totalBitWidth: totalBitWidth,
		parts:         parts,
		converter:     converter,
	}
}

func NewIntParameter(totalBitWidth, partBitWidth, base int) *Parameter {
	return NewNamedIntParameter("", totalBitWidth, partBitWidth, base)
}

func NewNamedIntParameter(name string, totalBitWidth, partBitWidth, base int) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)
	converter := NewIntConverter(partBitWidth, base)

	return &Parameter{
		name:          name,
		totalBitWidth: totalBitWidth,
		parts:         parts,
		converter:     converter,
	}
}

func NewFloatParameter(totalBitWidth, partBitWidth int) *Parameter {
	return NewNamedFloatParameter("", totalBitWidth, partBitWidth)
}

func NewNamedFloatParameter(name string, totalBitWidth, partBitWidth int) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)
	converter := NewFloatConverter(partBitWidth)

	return &Parameter{
		totalBitWidth: totalBitWidth,
		parts:         parts,
		converter:     converter,
	}
}

func (p *Parameter) Name() string {
	return p.name
}

func (p *Parameter) TotalBitWidth() int {
	return p.totalBitWidth
}

func (p *Parameter) PartBitWidth() int {
	return p.converter.GetBitWidth()
}

func (p *Parameter) Parts() int {
	return p.parts
}

func (p *Parameter) GetTextWidth() int {
	return p.converter.GetTextWidth()
}

func (p *Parameter) GetBitWidth() int {
	return p.converter.GetBitWidth()
}

func (p *Parameter) Base() int {
	return p.converter.GetBase()
}

func (p *Parameter) SetData(bytes []byte) {
	p.data = bytes
}

func (p *Parameter) DataAsString() string {
	return p.converter.BytesToString(p.data)
}

// This method is stateless and will be removed shortly
func (p *Parameter) StringToBytes(txt string) []byte {
	return p.converter.StringToBytes(txt)
}

// This method is stateless and will be removed shortly
func (p *Parameter) BytesToString(bytes []byte) string {
	return p.converter.BytesToString(bytes)
}

// TODO have a think about where this method should live
func (p *Parameter) Normalised(txt string) (string, bool) {
	return p.converter.Normalised(txt)
}

// TODO have a think about where this method should live
func (p *Parameter) IsStable(txt string) bool {
	return p.converter.IsStable(txt)
}

// TODO have a think about where this method should live
func (p *Parameter) InputAcceptor() func(string, rune) bool {
	return p.converter.InputAcceptor()
}

func partsCount(partBitWidth, totalBitWidth int) int {
	mustValidPartBitWidth(partBitWidth)
	mustValidTotalBitWidth(totalBitWidth)
	return totalBitWidth / partBitWidth
}

func mustValidTotalBitWidth(totalBitWidth int) {
	switch totalBitWidth {
	case 8, 16, 32, 64, 128, 256, 512:
	default:
		panic(fmt.Errorf("unsupported total bit width: %d", totalBitWidth))
	}
}

func mustValidPartBitWidth(partBitWidth int) {
	switch partBitWidth {
	case 8, 16, 32, 64:
	default:
		panic(fmt.Errorf("unsupported bit width value: %d", partBitWidth))
	}
}
