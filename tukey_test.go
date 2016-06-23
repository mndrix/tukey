package tukey // import "github.com/mndrix/tukey"

import (
	"fmt"
	"sort"
	"testing"
	"testing/quick"
)

func TestTukey(t *testing.T) {
	// these expected results are from R's function quantile(..., type=8)
	data := []float64{7, 3, 4, 2, 9, 8, 7, 4, 3, 20, 9, 7, 400}
	outliers, low, high := Outliers(1.5, data)
	if len(outliers) != 2 || outliers[0] != 20 || outliers[1] != 400 {
		t.Errorf("unexpected outliers: %#v", outliers)
	}
	if fmt.Sprintf("%.2f", low) != "-4.33" || fmt.Sprintf("%.2f", high) != "17.00" {
		t.Errorf("unexpected fences: %.2f %.2f", low, high)
	}
}

func TestQuick(t *testing.T) {
	// quartiles must have proper order
	f := func(xs []float64) (ok bool) {
		if len(xs) == 0 {
			return true // don't run this test
		}
		sort.Float64s(xs)

		lowerQuartile := Quantile(0.25, xs)
		median := Quantile(0.5, xs)
		upperQuartile := Quantile(0.75, xs)
		return lowerQuartile <= median && median <= upperQuartile
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
