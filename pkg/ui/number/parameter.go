package number

import (
	"fmt"
)

type Parameter struct {
	// Immutable
	name          string
	totalBitWidth int
	converter     Converter
	// Mutable
	partData [][]byte
}

func NewNamedUintParameter(name string, totalBitWidth, partBitWidth, base int) *Parameter {
	converter := NewUintConverter(partBitWidth, base)

	return newParameter(name, totalBitWidth, partBitWidth, converter)
}

func NewNamedIntParameter(name string, totalBitWidth, partBitWidth, base int) *Parameter {
	converter := NewIntConverter(partBitWidth, base)

	return newParameter(name, totalBitWidth, partBitWidth, converter)
}

func NewNamedFloatParameter(name string, totalBitWidth, partBitWidth int) *Parameter {
	converter := NewFloatConverter(partBitWidth)

	return newParameter(name, totalBitWidth, partBitWidth, converter)
}

func newParameter(name string, totalBitWidth, partBitWidth int, converter Converter) *Parameter {
	parts := partsCount(partBitWidth, totalBitWidth)

	partData := make([][]byte, parts)
	for i := range partData {
		partData[i] = make([]byte, partBitWidth/8)
	}

	return &Parameter{
		// Immutable
		name:          name,
		totalBitWidth: totalBitWidth,
		converter:     converter,
		// Mutable
		partData: partData,
	}
}

func (p *Parameter) Name() string {
	return p.name
}

func (p *Parameter) GoType() string {
	if p.converter.GetBitWidth() == p.totalBitWidth {
		return p.converter.GoType()
	} else {
		return fmt.Sprintf("*[%d]%s", (p.totalBitWidth / p.converter.GetBitWidth()), p.converter.GoType())
	}
}

func (p *Parameter) TotalBitWidth() int {
	return p.totalBitWidth
}

func (p *Parameter) PartBitWidth() int {
	return p.converter.GetBitWidth()
}

func (p *Parameter) Parts() int {
	return len(p.partData)
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

func (p *Parameter) String() string {
	return fmt.Sprintf("%s %s", p.name, p.GoType())
}

func (p *Parameter) SetData(bytes []byte) {
	if (len(bytes) * 8) != p.totalBitWidth {
		panic(fmt.Errorf("bad data update, received %d bits, but need %d", len(bytes)*8, p.totalBitWidth))
	}

	if len(bytes)%len(p.partData) != 0 {
		panic(fmt.Errorf("set data with %d bytes, not cleanly divisible by %d parts", len(bytes), len(p.partData)))
	}

	bytesPer := len(bytes) / len(p.partData)
	for i := range p.partData {
		idx := i * bytesPer
		chunk := bytes[idx : idx+bytesPer]
		p.partData[i] = chunk
	}
}

func (p *Parameter) DataFromStrings(txts []string) {
	for i, txt := range txts {
		p.partData[i] = p.converter.StringToBytes(txt)
	}
}

func (p *Parameter) DataToStrings() []string {
	strs := make([]string, len(p.partData))

	for i, part := range p.partData {
		strs[i] = p.converter.BytesToString(part)
	}

	return strs
}

// TODO this is temporary and should be deleted shortly
func (p *Parameter) FlatData() []byte {
	flatData := []byte{}
	for _, part := range p.partData {
		flatData = append(flatData, part...)
	}

	return flatData
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
