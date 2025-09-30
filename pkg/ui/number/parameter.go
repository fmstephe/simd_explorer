package number

import "fmt"

type Parameter struct {
	totalBitWidth int
	parts         int
	converter     Converter
}

func NewUintParameter(totalBitWidth, partBitWidth, base int) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)
	converter := NewUintConverter(partBitWidth, base)

	return &Parameter{
		totalBitWidth: totalBitWidth,
		parts:         parts,
		converter:     converter,
	}
}

func NewIntParameter(totalBitWidth, partBitWidth, base int) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)
	converter := NewIntConverter(partBitWidth, base)

	return &Parameter{
		totalBitWidth: totalBitWidth,
		parts:         parts,
		converter:     converter,
	}
}

func NewFloatParameter(totalBitWidth, partBitWidth int) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)
	converter := NewFloatConverter(partBitWidth)

	return &Parameter{
		totalBitWidth: totalBitWidth,
		parts:         parts,
		converter:     converter,
	}
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

func (p *Parameter) Base() int {
	return p.converter.GetBase()
}

func (p *Parameter) Converter() Converter {
	return p.converter
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
		panic(fmt.Errorf("Unsupported total bit width: %d", totalBitWidth))
	}
}

func mustValidPartBitWidth(partBitWidth int) {
	switch partBitWidth {
	case 8, 16, 32, 64:
	default:
		panic(fmt.Errorf("Unsupported bit width value: %d", partBitWidth))
	}
}
