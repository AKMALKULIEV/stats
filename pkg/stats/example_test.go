package stats

import (
	"fmt"

	"github.com/AKMALKULIEV/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       0,
			Amount:   50,
			Category: "",
		},
		{
			ID:       0,
			Amount:   50,
			Category: "",
		},
	}
	fmt.Println(Avg(payments))

	//Output:
	//50
}
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       0,
			Amount:   50,
			Category: "",
		},
		{
			ID:       0,
			Amount:   50,
			Category: "Abs",
		},
	}
    category := "Abs"
	fmt.Println(TotalInCategory(payments,types.Category(category) ))

	//Output:
	//50
}
