package indicator

import (
	"github.com/kietphan/go-indicators/errors"
	"github.com/zimmski/tavor/log"
)

type Rsi struct {
	previousAverageGain float64
	previousAverageLoss float64
	Period      int
	Data        []float64
}

func (rsi *Rsi)Calculator()([]float64, error) {
	lengthData := len(rsi.Data)
	if lengthData < rsi.Period {
		return []float64{}, errors.New("Not enough data to calculator")
	}
	rsiData := []float64{}
	for i, v := range rsi.Data {
		if i >= rsi.Period {
			if i == rsi.Period{
				rsi.calFirstGainLost()
				
			}else {
				gain, loss := rsi.calGainLoss(v,rsi.Data[i-1])
				rsi.previousAverageLoss = (rsi.previousAverageLoss * float64(rsi.Period -1) + loss)/float64(rsi.Period)
				rsi.previousAverageGain = (rsi.previousAverageGain * float64(rsi.Period -1) + gain)/float64(rsi.Period)
			}
			rs := rsi.previousAverageGain / rsi.previousAverageLoss
			log.Println("rs", rs)
			rsiValue := 100 - 100/(rs + 1)
			rsiData = append(rsiData, rsiValue)
		}
	}
	return rsiData, nil
}


func (rsi *Rsi)calFirstGainLost(){
	sumGain := 0.0
	sumLoss := 0.0
	for i:= 0; i <= rsi.Period ; i ++{
		if i == 0{
			sumGain = 0
			sumLoss = 0
		}else{
			if rsi.Data[i] < rsi.Data[i -1] {
				sumLoss += (rsi.Data[i - 1] - rsi.Data[i])
			}else {
				sumGain += (rsi.Data[i] - rsi.Data[i -1])
			}
		}
	}
	rsi.previousAverageLoss =  sumLoss / float64(rsi.Period)
	rsi.previousAverageGain =  sumGain / float64(rsi.Period)
	log.Println("previousAverageLoss", rsi.previousAverageLoss)
	log.Println("previousAverageGain", rsi.previousAverageGain)
	
}

func (rsi *Rsi)calGainLoss(currentData float64, previousData float64) (float64, float64 ) {
	loss := 0.0
	gain := 0.0
	if currentData < previousData{
		loss = previousData - currentData
	}else{
		gain = currentData - previousData
	}
	return gain, loss


}
