package indicator

import (
	"errors"
)
type Sma struct{
	Period int
	Multiplier float32
	ClosePrice float32
	
}

func (sma *Sma) calCulartorSma(data []float64) (float64, error){
	lengthData := len(data)
	if lengthData < sma.Period{
		return 0, errors.New("Not enough data to calculator")
	}
	smaData := []float64{}
	childArr := []float64{}
	for i, v := range data  {
		if i < lengthData -1 {
			childArr = append(childArr, v)
			//smaData = append(smaData, nil)
		}else{
			childArr = childArr[1:]
			childArr = append(childArr, v)
			smaOfChildArd := sma.Average(childArr)
			smaData = append(smaData, smaOfChildArd)
		}
	}
	return sma.Average(data), nil
}

func (sma Sma) Average(xs[]float64)float64 {
	total:=0.0
	for _,v:=range xs {
		total +=v
	}
	return total/float64(len(xs))
}