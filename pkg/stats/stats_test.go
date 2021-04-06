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

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "food"},
	}
	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
	}
	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {

	first := map[types.Category]types.Money{
		"auto":   100,
		"mobile": 5000,
	}

	second := map[types.Category]types.Money{
		"auto":   100,
		"mobile": -5000,
		"fun":    10,
	}
	expected := map[types.Category]types.Money{
		"auto":   0,
		"mobile": -10000,
		"fun":    10,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("something wrong expected: %v ,result: %v", expected, result)
	}
}
