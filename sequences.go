// package sequences
// covers functonality of creating sequences of numbers
// Sources of study matierials
// https://stackoverflow.com/questions/24852317/how-to-count-decimal-places-of-float
// https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144#37533144
// https://stackoverflow.com/questions/24852317/how-to-count-decimal-places-of-float/24852368#24852368
package sequences

import (
	"constraints"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	seq_logger *log.Logger
)

// Runs initialisation at runtime start
func init() {
	seq_logger = log.New(os.Stderr, "sequences: ", 0)
}

// is_int checks if a numeric value is float or int
// returns true in case of float 'a' value
func is_int[T constraints.Float | constraints.Integer](a T) bool {
	return float64(a) != float64(int64(a))
}

// determines the number of decimal places of a float number
func numDecPlaces(v float64) int {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	//log.Println("HERE", s)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}

func round_to_arbitrary_precision[T constraints.Float | constraints.Integer, T2 constraints.Float | constraints.Integer](x T, y T2) float64 {
	x_float64 := float64(x)
	y_float64 := float64(y)
	return math.Round(x_float64*math.Pow(10, y_float64)) / math.Pow(10, y_float64)
}

// Seq_slice
// exported
// produce sequence of numbers from to by a number
func Seq_slice[T constraints.Float | constraints.Integer](from, to, by T) []T {

	decimal_by := is_int(by) // checking if 'by' is a int or float value. If it is float, determin the number of decimal places in order to round properly
	var number_decimals int  // declare and assign value 0, the default value for int
	if decimal_by {
		number_decimals = numDecPlaces(float64(by))
	}
	//log.Println("number_decimals", number_decimals)

	var out []T

	if from > to {

		if decimal_by {

			seq_logger.Println("'from'>'to' & T is of type float: float decrementing")
			for i := from; i >= to; i -= by {
				out = append(out, T(round_to_arbitrary_precision(i, number_decimals)))
			}

		} else {

			seq_logger.Println("'from'>'to' & T is of type int: int decrementing")
			for i := from; i >= to; i -= by {
				out = append(out, i)
			}
		}

	} else {

		if decimal_by {

			seq_logger.Println("'to'>'from' & T is of type float: float incrementing")
			for i := from; i <= to; i += by {
				out = append(out, T(round_to_arbitrary_precision(i, number_decimals)))
			}

		} else {

			seq_logger.Println("'to'>'from' & T is of type int: int incrementing")
			for i := from; i <= to; i += by {
				out = append(out, i)
			}
		}

	}

	return out
}
