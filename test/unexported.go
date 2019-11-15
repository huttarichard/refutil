package test

// SampleType represent category
// for samples
type SampleType uint

const (
	// SampleInt sample int type
	SampleInt SampleType = iota
	// SampleUint sample uint type
	SampleUint
	// SampleString sample string type
	SampleString
	// SampleStringPtr sample string pointer type
	SampleStringPtr
	// SampleStringNil sample string pointer type which is nil
	SampleStringNil
	// SampleStringSlice sample string slice type
	SampleStringSlice
	// SampleChanBool sample booelan chan  type
	SampleChanBool
	// SampleEmptyZeroer sample zeroer struct type which is nil
	SampleEmptyZeroer
	// SamplePositiveZeroer sample zeroer struct  type
	SamplePositiveZeroer
	// SampleZeroer sample zeroer struct  type
	SampleZeroer
	// SampleStruct sample struct  type
	SampleStruct
)

var sampleStringValue = "test"

var samples = map[SampleType]interface{}{
	SampleInt:            sampleInt(1),
	SampleUint:           sampleUint(1),
	SampleString:         sampleString(sampleStringValue),
	SampleStringNil:      sampleStringPtr((*string)(nil)),
	SampleStringPtr:      sampleStringPtr((*string)(&sampleStringValue)),
	SampleStringSlice:    sampleStringSlice([]string{sampleStringValue}),
	SampleChanBool:       sampleChanBool(make(chan bool)),
	SampleZeroer:         &sampleZeroer{positive: false},
	SamplePositiveZeroer: &sampleZeroer{positive: true},
	SampleEmptyZeroer:    (*sampleZeroer)(nil),
	SampleStruct: &sampleStruct{
		Test:   "test",
		test:   "test",
		interf: "test",
	},
}

// Sample return sample interface.
// This is intended for testing unexported types in refutil package
func Sample(t SampleType) interface{} {
	m, ok := samples[t]
	if !ok {
		panic("trying to get unregistered sample")
	}
	return m
}
