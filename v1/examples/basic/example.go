package main

import (
	"encoding/json"
	"math"
	"os"

	"github.com/foxcapades/gojasc/v1/pkg/j57"
)

func main() {
	data := j57.SerializeComplex128(complex(math.MaxFloat64, math.MaxFloat64))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	_ = enc.Encode(string(data))

	// Outputs:
	// ".<=<[[=U=7(/.<=<[[=U=7(/"
}
