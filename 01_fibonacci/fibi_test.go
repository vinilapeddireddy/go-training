package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)


func Test_Fib_negativeValue(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	fibonacci(-1)
    // Then
	expected := strconv.Quote("1")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func Test_Fib_1(t *testing.T) {
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	fibonacci(1)
    // Then
	expected := strconv.Quote("1")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}
