package money

import (
	"log"
	"testing"
)

// TestMoneys test value conversion.
func TestMoneys(t *testing.T) {
	strPrice := "1.391.439,12345"

	price, err := Parse(strPrice, ",")
	if err != nil {
		log.Fatalf("Error, could not parse the value: %s", strPrice)
	}
	// log.Println("Price: ", price)
	log.Printf("Price: %#2.2f", price)
	log.Println("price integer: ", price.Int())
	log.Println("price decimal: ", price.Cents())
	var sumPrice float64 = 683945152
	units := 6318
	sumPriceMoney := Money(sumPrice)
	log.Printf("Average price raw: %.8f", sumPrice/float64(units))
	log.Printf("Average price : %.8f", sumPriceMoney.Divide(units))
}
