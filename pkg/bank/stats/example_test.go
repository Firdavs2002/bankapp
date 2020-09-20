package stats

import (
	"fmt"

	"github.com/Firdavs2002/bankapp/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   53_00,
			Category: "Cat",
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "Cat",
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}
