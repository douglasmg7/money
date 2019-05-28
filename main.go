package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Money is a value with decimal precision of two.
type Money float32

func main() {
	fmt.Println("vim-go")

	price, err := Parse("000,3", "g")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Price: ", price)

	// price2, err := Parse("3,01", true)
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

// Cents return the cents part.
func (m *Money) Cents() string {
	// _, frac := math.Modf(float64(*m))
	// return fmt.Sprintf("%.2f", frac)

	vals := strings.Split(fmt.Sprintf("%.2f", float32(*m)), ".")
	return vals[1]
}

// Int return the integer part.
func (m *Money) Int() string {
	// _, frac := math.Modf(float64(*m))
	// return fmt.Sprintf("%.2f", frac)

	return fmt.Sprintf("%.0f", float32(*m))
}

// New create a Money value.
func New(val float64) Money {
	return Money(math.Round(val*100) / 100)
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
	val64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return Money(math.Round(val64*100) / 100), nil
}
