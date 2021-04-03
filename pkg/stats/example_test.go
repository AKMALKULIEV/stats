package stats

import (
	"fmt"

	"github.com/AKMALKULIEV/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 0,
			Amount: 50,
			Category: "",
			Status: "FAIL",
		},
		{
			ID: 0,
			Amount: 50,
			Category: "",
			Status: "",
		},
	}
	fmt.Println(Avg(payments))

	//Output:
	//50
}
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 0,
			Amount: 50,
			Category: "",
			Status: "FAIL",
		},
		{
			ID: 0,
			Amount: 50,
			Category: "Abs",
			Status: "",
		},
	}
    category := "Abs"
	fmt.Println(TotalInCategory(payments,types.Category(category) ))

	//Output:
	//50
}
