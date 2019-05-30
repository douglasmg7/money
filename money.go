package money

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Money is a value with decimal precision of two.
type Money float64

func main() {
	// price, err := Parse("000,3", "g")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("Price: ", price)

	// price2, err := Parse("40.979,00", true)
	// if err != nil {
	// 	log.Fatalln("Error, could not parse the value.")
	// }
	// fmt.Println("price integer: ", price.Int())
	// fmt.Println("price decimal: ", price.Cents())
	// total := price + price2
	// log.Printf("Total type: %T", total)
	// log.Println("Total: ", total)
	// log.Println("Total: ", total/2)
}

// Divide return money divided.
func (m *Money) Divide(val int) Money {
	return New(float64(*m) / float64(val))
}

// Cents return the cents part.
func (m *Money) Cents() string {
	vals := strings.Split(fmt.Sprintf("%.2f", *m), ".")
	return vals[1]
}

// Int return the integer part.
func (m *Money) Int() string {
	return fmt.Sprintf("%.0f", *m)
}

// New create a Money value.
func New(val float64) Money {
	sVal := fmt.Sprintf("%.2f", val)
	newVal, err := strconv.ParseFloat(sVal, 64)
	if err != nil {
		log.Fatalf("New(), Could not convert parse to float")
	}
	return Money(newVal)
}

// Parse from a string.
func Parse(str string, sep string) (Money, error) {
	switch sep {
	case ",":
		str = strings.Replace(str, ".", "", -1)
		str = strings.Replace(str, ",", ".", -1)
	case ".":
	case "":
	default:
		return 0, errors.New(`Invalid separator, must be "", "," or "."`)
	}
	str = strings.TrimSpace(str)
	val64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return Money(math.Round(val64*100) / 100), nil
}
