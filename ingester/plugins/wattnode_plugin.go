package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)
func has_device(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.WattnodeState!= nil
}
var device_units = map[string]string{
	"EnergySum":	"kWh",
	"EnergyPosSUm":	"kWh",
	"EnergySumNR":	"kWh",
	"EnergyPosSumNr":	"kWh",
	"PowerSum":	"W",
	"PowerA":	"W",
	"PowerB":	"W",
	"PowerC":	"W",
	"VoltAvgLN":	"V",
	"VoltA":	"V",
	"VoltB":	"V",
	"VoltC":	"V",
	"VoltAvgLL":	"V",
	"VoltAB":	"V",
	"VoltBC":	"V",
	"VoltAC":	"V",
	"Freq":	"Hz",
	"EnergyA":	"kWh",
	"EnergyB":	"kWh",
	"EnergyC":	"kWh",
	"EnergyPosA":	"kWh",
	"EnergyPosB":	"kWh",
	"EnergyPosC":	"kWh",
	"EnergyNegSum":	"kWh",
	"EnergyNegSumNR":	"kWh",
	"EnergyNegA":	"kWh",
	"EnergyNegB":	"kWh",
	"EnergyNegC":	"kWh",
	"EnergyReacSum":	"kVARh",
	"EnergyReacA":	"kVARh",
	"EnergyReacB":	"kVARh",
	"EnergyReacC":	"kVARh",
	"EnergyAppSum":	"kVAh",
	"EnergyAppA":	"kVAh",
	"EnergyAppB":	"kVAh",
	"EnergyAppC":	"kVAh",
	"PowerFactorAvg":	"",
	"PowerFactorA":	"",
	"PowerFactorB":	"",
	"PowerFactorC":	"",
	"PowerReacSum":	"VAR",
	"PowerReacA":	"VAR",
	"PowerReacB":	"VAR",
	"PowerReacC":	"VAR",
	"PowerAppSum":	"VA",
	"PowerAppA":	"VA",
	"PowerAppB":	"VA",
	"PowerAppC":	"VA",
	"CurrentA":	"A",
	"CurrentB":	"A",
	"CurrentC":	"A",
	"Demand":	"W",
	"DemandMin":	"W",
	"DemandMax":	"W",
	"DemandApp":	"W",
	"DemandA":	"W",
	"DemandB":	"W",
	"DemandC":	"W",
}
var device_lookup = map[string]func(msg xbospb.XBOS) (float64, bool){

	"EnergySum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergySum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergySum.Value), true
		}
		return 0, false
	},
	"EnergyPosSUm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyPosSUm != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyPosSUm.Value), true
		}
		return 0, false
	},
	"EnergySumNR": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergySumNR != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergySumNR.Value), true
		}
		return 0, false
	},
	"EnergyPosSumNr": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyPosSumNr != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyPosSumNr.Value), true
		}
		return 0, false
	},
	"PowerSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerSum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerSum.Value), true
		}
		return 0, false
	},
	"PowerA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerA.Value), true
		}
		return 0, false
	},
	"PowerB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerB.Value), true
		}
		return 0, false
	},
	"PowerC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerC.Value), true
		}
		return 0, false
	},
	"VoltAvgLN": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltAvgLN != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltAvgLN.Value), true
		}
		return 0, false
	},
	"VoltA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltA.Value), true
		}
		return 0, false
	},
	"VoltB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltB.Value), true
		}
		return 0, false
	},
	"VoltC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltC.Value), true
		}
		return 0, false
	},
	"VoltAvgLL": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltAvgLL != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltAvgLL.Value), true
		}
		return 0, false
	},
	"VoltAB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltAB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltAB.Value), true
		}
		return 0, false
	},
	"VoltBC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltBC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltBC.Value), true
		}
		return 0, false
	},
	"VoltAC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.VoltAC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.VoltAC.Value), true
		}
		return 0, false
	},
	"Freq": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.Freq != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.Freq.Value), true
		}
		return 0, false
	},
	"EnergyA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyA.Value), true
		}
		return 0, false
	},
	"EnergyB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyB.Value), true
		}
		return 0, false
	},
	"EnergyC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyC.Value), true
		}
		return 0, false
	},
	"EnergyPosA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyPosA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyPosA.Value), true
		}
		return 0, false
	},
	"EnergyPosB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyPosB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyPosB.Value), true
		}
		return 0, false
	},
	"EnergyPosC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyPosC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyPosC.Value), true
		}
		return 0, false
	},
	"EnergyNegSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyNegSum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyNegSum.Value), true
		}
		return 0, false
	},
	"EnergyNegSumNR": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyNegSumNR != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyNegSumNR.Value), true
		}
		return 0, false
	},
	"EnergyNegA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyNegA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyNegA.Value), true
		}
		return 0, false
	},
	"EnergyNegB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyNegB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyNegB.Value), true
		}
		return 0, false
	},
	"EnergyNegC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyNegC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyNegC.Value), true
		}
		return 0, false
	},
	"EnergyReacSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyReacSum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyReacSum.Value), true
		}
		return 0, false
	},
	"EnergyReacA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyReacA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyReacA.Value), true
		}
		return 0, false
	},
	"EnergyReacB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyReacB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyReacB.Value), true
		}
		return 0, false
	},
	"EnergyReacC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyReacC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyReacC.Value), true
		}
		return 0, false
	},
	"EnergyAppSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyAppSum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyAppSum.Value), true
		}
		return 0, false
	},
	"EnergyAppA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyAppA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyAppA.Value), true
		}
		return 0, false
	},
	"EnergyAppB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyAppB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyAppB.Value), true
		}
		return 0, false
	},
	"EnergyAppC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.EnergyAppC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.EnergyAppC.Value), true
		}
		return 0, false
	},
	"PowerFactorAvg": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerFactorAvg != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerFactorAvg.Value), true
		}
		return 0, false
	},
	"PowerFactorA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerFactorA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerFactorA.Value), true
		}
		return 0, false
	},
	"PowerFactorB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerFactorB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerFactorB.Value), true
		}
		return 0, false
	},
	"PowerFactorC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerFactorC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerFactorC.Value), true
		}
		return 0, false
	},
	"PowerReacSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerReacSum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerReacSum.Value), true
		}
		return 0, false
	},
	"PowerReacA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerReacA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerReacA.Value), true
		}
		return 0, false
	},
	"PowerReacB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerReacB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerReacB.Value), true
		}
		return 0, false
	},
	"PowerReacC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerReacC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerReacC.Value), true
		}
		return 0, false
	},
	"PowerAppSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerAppSum != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerAppSum.Value), true
		}
		return 0, false
	},
	"PowerAppA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerAppA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerAppA.Value), true
		}
		return 0, false
	},
	"PowerAppB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerAppB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerAppB.Value), true
		}
		return 0, false
	},
	"PowerAppC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.PowerAppC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.PowerAppC.Value), true
		}
		return 0, false
	},
	"CurrentA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.CurrentA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.CurrentA.Value), true
		}
		return 0, false
	},
	"CurrentB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.CurrentB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.CurrentB.Value), true
		}
		return 0, false
	},
	"CurrentC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.CurrentC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.CurrentC.Value), true
		}
		return 0, false
	},
	"Demand": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.Demand != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.Demand.Value), true
		}
		return 0, false
	},
	"DemandMin": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.DemandMin != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.DemandMin.Value), true
		}
		return 0, false
	},
	"DemandMax": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.DemandMax != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.DemandMax.Value), true
		}
		return 0, false
	},
	"DemandApp": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.DemandApp != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.DemandApp.Value), true
		}
		return 0, false
	},
	"DemandA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.DemandA != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.DemandA.Value), true
		}
		return 0, false
	},
	"DemandB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.DemandB != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.DemandB.Value), true
		}
		return 0, false
	},
	"DemandC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.WattnodeState.DemandC != nil {
			return float64(msg.XBOSIoTDeviceState.WattnodeState.DemandC.Value), true
		}
		return 0, false
	},
}
func build_device(uri types.SubscriptionURI, name string, msg xbospb.XBOS) types.ExtractedTimeseries {
	
	if extractfunc, found := device_lookup[name]; found {
		if value, found := extractfunc(msg); found {
			var extracted types.ExtractedTimeseries
			time := int64(msg.XBOSIoTDeviceState.Time)
			extracted.Values = append(extracted.Values, value)
			extracted.Times = append(extracted.Times, time)
			extracted.UUID = types.GenerateUUID(uri, []byte(name))
			extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
			extracted.Tags = map[string]string{
				"unit": device_units[name],
				"name": name,
			}
			return extracted
		}
	}
return types.ExtractedTimeseries{}
}

func Extract(uri types.SubscriptionURI, msg xbospb.XBOS, add func(types.ExtractedTimeseries) error) error {
	if msg.XBOSIoTDeviceState != nil {
		if has_device(msg) {
			for name := range device_lookup {
				extracted := build_device(uri, name, msg)
				if err := add(extracted); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
