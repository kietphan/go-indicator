package indicator

import (
	"github.com/kietphan/go-indicators/errors"
	"github.com/kietphan/go-indicators/untils"
	"fmt"
)

type Ema struct {
	Period      int
	Multiplier  float64
	emaPrevious float64
}

func (ema *Ema) CalCulartorEma(data []float64) ([]float64, error) {
	
	lengthData := len(data)
	if lengthData < ema.Period {
		return []float64{}, errors.New("Not enough data to calculator")
	}
	ema.calCulatorMultiplier()
	emaData := []float64{}
	childArr := []float64{}
	for i, v := range data {
		if i < ema.Period-1 {
			childArr = append(childArr, v)
			//smaData = append(smaData, nil)
		} else if i == ema.Period-1 {
			childArr = append(childArr, v)
			ema.emaPrevious = untils.Average(childArr)
			emaData = append(emaData, ema.emaPrevious)
		} else {
			newEma := (v-ema.emaPrevious)*ema.Multiplier + ema.emaPrevious
			ema.emaPrevious = newEma
			emaData = append(emaData, newEma)
		}
	}
	return emaData, nil
}

func (ema *Ema) calCulatorMultiplier() {
	multiplier := 2 / (float64(ema.Period) + 1)
	fmt.Println("multiplier", multiplier)
	ema.Multiplier = multiplier
}
