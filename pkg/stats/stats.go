package stats

import (
	"github.com/AKMALKULIEV/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	
	num := 0
	for _, payment := range payments {
		sum += payment.Amount

	}
	for _, payment := range payments {
		if payment.Amount < 0  && payment.Status == "FAIL" {
			continue
		}
		num += 1
	}
	return sum / types.Money(num)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category && payment.Status == "FAIL" {
			continue
		}
		mPayments := payment.Amount
		sum += mPayments
	}
	return sum
}

// func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
// 	categories := map[types.Category]types.Money{}
	

// 	for _, payment := range payments {
// 		categorynum = len(payments.Category{})
// 		categories[payment.Category] = (categories[payment.Category] + payment.Amount)/categorynum
		
// 	}
// 	return categories
// }

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	count := map[types.Category]types.Money{}
	result := map[types.Category]types.Money{}
  
	for _, payment := range payments {
	  result[payment.Category] += payment.Amount
	  count[payment.Category]++
	}
  
	for key := range result {
	  result[key] /= count[key]
	}
  
	return result
  }