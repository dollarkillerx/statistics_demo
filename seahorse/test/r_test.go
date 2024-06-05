package test

import (
	"fmt"
	"math"
	"testing"
)

func TestRp(t *testing.T) {
	r := math.Round(0.012345*100) / 100
	fmt.Println(r)
}
