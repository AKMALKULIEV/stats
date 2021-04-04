package stats

import (
	"reflect"
	"testing"

	"github.com/AKMALKULIEV/bank/v2/pkg/types"
)

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