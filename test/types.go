package test

type sampleInt int

type sampleUint uint

type sampleString string

type sampleStringPtr *string

type sampleStringSlice []string

type sampleChanBool chan bool

type zeroer interface {
	IsZero() bool
}

type sampleZeroer struct {
	positive bool
}

var _ zeroer = &sampleZeroer{}

func (p *sampleZeroer) IsZero() bool {
	return p.positive
}

type sampleStruct struct {
	test   string
	interf interface{}
	Test   string
}
