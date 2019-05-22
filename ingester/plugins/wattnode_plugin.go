package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)

func has_device(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.Wattnode!= nil
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
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergySum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergySum.Value), true
		}
		return 0, false
	},
	"EnergyPosSUm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyPosSUm != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyPosSUm.Value), true
		}
		return 0, false
	},
	"EnergySumNR": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergySumNR != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergySumNR.Value), true
		}
		return 0, false
	},
	"EnergyPosSumNr": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyPosSumNr != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyPosSumNr.Value), true
		}
		return 0, false
	},
	"PowerSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerSum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerSum.Value), true
		}
		return 0, false
	},
	"PowerA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerA.Value), true
		}
		return 0, false
	},
	"PowerB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerB.Value), true
		}
		return 0, false
	},
	"PowerC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerC.Value), true
		}
		return 0, false
	},
	"VoltAvgLN": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltAvgLN != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltAvgLN.Value), true
		}
		return 0, false
	},
	"VoltA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltA.Value), true
		}
		return 0, false
	},
	"VoltB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltB.Value), true
		}
		return 0, false
	},
	"VoltC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltC.Value), true
		}
		return 0, false
	},
	"VoltAvgLL": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltAvgLL != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltAvgLL.Value), true
		}
		return 0, false
	},
	"VoltAB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltAB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltAB.Value), true
		}
		return 0, false
	},
	"VoltBC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltBC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltBC.Value), true
		}
		return 0, false
	},
	"VoltAC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.VoltAC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.VoltAC.Value), true
		}
		return 0, false
	},
	"Freq": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.Freq != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.Freq.Value), true
		}
		return 0, false
	},
	"EnergyA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyA.Value), true
		}
		return 0, false
	},
	"EnergyB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyB.Value), true
		}
		return 0, false
	},
	"EnergyC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyC.Value), true
		}
		return 0, false
	},
	"EnergyPosA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyPosA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyPosA.Value), true
		}
		return 0, false
	},
	"EnergyPosB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyPosB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyPosB.Value), true
		}
		return 0, false
	},
	"EnergyPosC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyPosC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyPosC.Value), true
		}
		return 0, false
	},
	"EnergyNegSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyNegSum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyNegSum.Value), true
		}
		return 0, false
	},
	"EnergyNegSumNR": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyNegSumNR != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyNegSumNR.Value), true
		}
		return 0, false
	},
	"EnergyNegA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyNegA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyNegA.Value), true
		}
		return 0, false
	},
	"EnergyNegB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyNegB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyNegB.Value), true
		}
		return 0, false
	},
	"EnergyNegC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyNegC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyNegC.Value), true
		}
		return 0, false
	},
	"EnergyReacSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyReacSum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyReacSum.Value), true
		}
		return 0, false
	},
	"EnergyReacA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyReacA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyReacA.Value), true
		}
		return 0, false
	},
	"EnergyReacB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyReacB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyReacB.Value), true
		}
		return 0, false
	},
	"EnergyReacC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyReacC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyReacC.Value), true
		}
		return 0, false
	},
	"EnergyAppSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyAppSum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyAppSum.Value), true
		}
		return 0, false
	},
	"EnergyAppA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyAppA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyAppA.Value), true
		}
		return 0, false
	},
	"EnergyAppB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyAppB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyAppB.Value), true
		}
		return 0, false
	},
	"EnergyAppC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.EnergyAppC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.EnergyAppC.Value), true
		}
		return 0, false
	},
	"PowerFactorAvg": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerFactorAvg != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerFactorAvg.Value), true
		}
		return 0, false
	},
	"PowerFactorA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerFactorA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerFactorA.Value), true
		}
		return 0, false
	},
	"PowerFactorB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerFactorB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerFactorB.Value), true
		}
		return 0, false
	},
	"PowerFactorC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerFactorC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerFactorC.Value), true
		}
		return 0, false
	},
	"PowerReacSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerReacSum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerReacSum.Value), true
		}
		return 0, false
	},
	"PowerReacA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerReacA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerReacA.Value), true
		}
		return 0, false
	},
	"PowerReacB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerReacB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerReacB.Value), true
		}
		return 0, false
	},
	"PowerReacC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerReacC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerReacC.Value), true
		}
		return 0, false
	},
	"PowerAppSum": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerAppSum != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerAppSum.Value), true
		}
		return 0, false
	},
	"PowerAppA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerAppA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerAppA.Value), true
		}
		return 0, false
	},
	"PowerAppB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerAppB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerAppB.Value), true
		}
		return 0, false
	},
	"PowerAppC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.PowerAppC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.PowerAppC.Value), true
		}
		return 0, false
	},
	"CurrentA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.CurrentA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.CurrentA.Value), true
		}
		return 0, false
	},
	"CurrentB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.CurrentB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.CurrentB.Value), true
		}
		return 0, false
	},
	"CurrentC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.CurrentC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.CurrentC.Value), true
		}
		return 0, false
	},
	"Demand": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.Demand != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.Demand.Value), true
		}
		return 0, false
	},
	"DemandMin": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.DemandMin != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.DemandMin.Value), true
		}
		return 0, false
	},
	"DemandMax": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.DemandMax != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.DemandMax.Value), true
		}
		return 0, false
	},
	"DemandApp": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.DemandApp != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.DemandApp.Value), true
		}
		return 0, false
	},
	"DemandA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.DemandA != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.DemandA.Value), true
		}
		return 0, false
	},
	"DemandB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.DemandB != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.DemandB.Value), true
		}
		return 0, false
	},
	"DemandC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Wattnode.DemandC != nil {
			return float64(msg.XBOSIoTDeviceState.Wattnode.DemandC.Value), true
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
