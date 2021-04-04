package stats

import (
	"reflect"
	"testing"

	"github.com/AKMALKULIEV/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 30, Category: "auto"},
		{ID: 1, Amount: 10, Category: "auto"},
		{ID: 1, Amount: 10, Category: "mobile"},
		{ID: 1, Amount: 10, Category: "fun"},
	}
	expected := map[types.Category]types.Money{
		"auto":   20,
		"mobile": 10,
		"fun":    10,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("something wrong expected: %v ,result: %v", expected, result)
	}
}
