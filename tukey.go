// This package implements the Tukey Method (sometimes called Tukey Fences) for
// locating outliers within a sample.  The Tukey Method has the following
// advantages:
//
//  * detects multiple outliers at once
//  * detects high and low outliers at once
//  * implementation is easy to understand
//  * tunable parameter to adjust how strict the detection is
//  * not influenced by extreme values
//    (since it doesn't use mean or standard deviation)
//  * requires no assumptions about the population's distribution
package tukey // import "github.com/mndrix/tukey"

import (
	"math"
	"sort"
)

// Outliers returns outliers drawn from xs using Tukey Fences with the
// multiplier m. The traditional multiplier for outliers is 1.5 and for extreme
// outliers is 3.  There's no statistical basis for those multipliers, so feel
// free to use something that fits your preferences.
//
// The second and third values are the lower and upper fences, respectively,
// which were used to locate outliers.
func Outliers(m float64, xs []float64) ([]float64, float64, float64) {
	sort.Float64s(xs)

	lowerQuartile := Quantile(0.25, xs)
	upperQuartile := Quantile(0.75, xs)
	iqr := upperQuartile - lowerQuartile
	lowFence := lowerQuartile - m*iqr
	highFence := upperQuartile + m*iqr

	outliers := make([]float64, 0)
	for _, x := range xs {
		if x < lowFence || x > highFence {
			outliers = append(outliers, x)
		}
	}

	return outliers, lowFence, highFence
}

// Quantile returns the quantile with probability p for the slice xs using
// method 8, as recommended in the paper below.  The slice xs must be sorted
// from smallest to largest and may not be empty.
//
// See "Sample Quantiles in Statistical Packages" by Hyndman and Fan from which
// this implementation is derived. See also R documentation for the quantile()
// function which summarizes that paper.
func Quantile(p float64, xs []float64) float64 {
	if len(xs) == 0 {
		panic("Quantile() second argument may not be empty")
	}
	if len(xs) == 1 {
		return xs[0]
	}

	// log.Printf("xs = %#v\n", xs)
	// log.Printf("p = %.2f\n", p)
	n := float64(len(xs))
	// log.Printf("n = %.2f\n", n)

	// parameters for "Definition 8" (which the paper recommends)
	m := (p + 1) / 3
	// log.Printf("m = %.2f\n", m)

	// equations based on the parameters
	j := math.Floor(p*n + m)
	g := p*n + m - j
	gamma := g
	// log.Printf("j = %.2f\n", j)
	// log.Printf("g = %.2f\n", g)

	// j is 1-based but xs is 0-based, so lookups subtract 1
	quantile := (1-gamma)*xs[int(j)-1] + gamma*xs[int(j)]
	// log.Printf("quantile = %.2f\n\n", quantile)
	return quantile
}
