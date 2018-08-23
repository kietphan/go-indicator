package main

import (
	"fmt"
	"github.com/kietphan/go-indicators/indicator"
	"github.com/kietphan/go-indicators/untils"
)

func main()  {
	ema  := &indicator.Ema{
		Period: 10,
	}
	data := []float64{22.27, 22.19, 22.08, 22.17, 22.18, 22.13, 22.23, 22.43, 22.24, 22.29, 22.15, 22.39, 22.38, 22.61, 23.36, 24.05, 23.75, 23.83, 23.95, 23.63, 23.82, 23.87, 23.65, 23.19, 23.10, 23.33, 22.68, 23.10, 22.40, 22.17}
	fmt.Println(data)
	result, err := ema.CalCulartorEma(data)
	fmt.Println(err)
	fmt.Println("adasd",untils.Average([]float64{22.27, 22.19, 22.08, 22.17, 22.18, 22.13, 22.23, 22.43, 22.24, 22.29}))
	fmt.Println(result)
	data = []float64{44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08, 45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64, 46.21, 46.25, 45.71, 46.45, 45.78, 45.35, 44.03, 44.18, 44.22, 44.57, 43.42, 42.66, 43.13}
	rsi := &indicator.Rsi{
		Period: 14,
		Data: data,
	}
	result, _ = rsi.Calculator()
	fmt.Println("rsi")
	fmt.Println(result)
}
