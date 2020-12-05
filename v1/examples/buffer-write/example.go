package main

import (
	"encoding/json"
	"math"
	"os"

	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
)

func main() {
	data := make([]byte, 64)
	offset := tally.UTally(0)

	j57.AppendBool(true, data, &offset)
	j57.AppendComplex128(complex(math.MaxFloat64, math.MaxFloat64), data, &offset)
	j57.AppendString("fubar", data, &offset)
	j57.AppendInt(33, data, &offset)
	j57.AppendFloat64(math.MaxFloat64, data, &offset)
	j57.AppendUint32(math.MaxUint32, data, &offset)
	j57.AppendInt16(math.MaxInt16, data, &offset)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	_ = enc.Encode(string(data[:offset]))

	// Outputs:
	// "$$.<=<[[=U=7(/.<=<[[=U=7(/$$$(*15==I9)$D.<=<[[=U=7(/)**TS+;&-'T"
}
