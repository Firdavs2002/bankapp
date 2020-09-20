package stats

import (
	"github.com/Firdavs2002/bankapp/pkg/bank/types"
)

//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	countPayments := types.Money(len(payments))
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys / countPayments
}

// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys
}
