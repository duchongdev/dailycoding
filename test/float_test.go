package test

import (
	"math"
	"testing"
)

func TestFloat(t *testing.T) {

	num := float32(10) / float32(3)
        //Round 四舍五入取整
	num2 := math.Round(float64(num) * 1000)
	t.Logf("num1 = %v", num)
	t.Logf("num2 = %v", num2)
}
