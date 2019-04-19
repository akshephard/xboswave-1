package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
    //"reflect"
)
func has_device(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.WeatherPrediction != nil
}
var device_units = map[string]string{
	"time":	"seconds",
	"icon":	"",
	"precipIntensity":	"inches per hour",
	"precipIntensityError":	"",
	"precipProbability":	"",
	"precipAccumulation":	"inches",
	"precipType":	"",
	"temperature":	"F",
	"apparentTemperature":	"F",
	"dewPoint":	"F",
	"humidity":	"",
	"pressure":	"millibars",
	"windSpeed":	"miles per hour",
	"windGust":	"miles per hour",
	"windBearing":	"degree",
	"cloudCover":	"",
	"uvIndex":	"",
	"visibility":	"miles",
	"ozone":	"Dobson",
}
/*
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
*/

func Extract(uri types.SubscriptionURI, msg xbospb.XBOS, add func(types.ExtractedTimeseries) error) error {
	if msg.XBOSIoTDeviceState != nil {
		if has_device(msg) {
            //fmt.Printf("The length of the the message is: %d", len(msg.XBOSIoTDeviceState.WeatherPrediction.Predictions))
			for _, _prediction := range msg.XBOSIoTDeviceState.WeatherPrediction.Predictions {
                prediction := _prediction.Prediction
                //fmt.Printf("part of the message: %v", prediction)
                //fmt.Println(reflect.TypeOf(prediction))
				var extracted types.ExtractedTimeseries
				var name string
				time := int64(msg.XBOSIoTDeviceState.Time)
				step := (int64(_prediction.PredictionTime) - time) / 1e9
				extracted.Times = append(extracted.Times, time)

            	if prediction.time != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.time.Value))
            		name = "time"
            	} else {
            	continue
            	}
            	if prediction.icon != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.icon.Value))
            		name = "icon"
            	} else {
            	continue
            	}
            	if prediction.precipIntensity != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.precipIntensity.Value))
            		name = "precipintensity"
            	} else {
            	continue
            	}
            	if prediction.precipIntensityError != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.precipIntensityError.Value))
            		name = "precipintensityerror"
            	} else {
            	continue
            	}
            	if prediction.precipProbability != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.precipProbability.Value))
            		name = "precipprobability"
            	} else {
            	continue
            	}
            	if prediction.precipAccumulation != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.precipAccumulation.Value))
            		name = "precipaccumulation"
            	} else {
            	continue
            	}
            	if prediction.precipType != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.precipType.Value))
            		name = "preciptype"
            	} else {
            	continue
            	}
            	if prediction.temperature != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.temperature.Value))
            		name = "temperature"
            	} else {
            	continue
            	}
            	if prediction.apparentTemperature != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.apparentTemperature.Value))
            		name = "apparenttemperature"
            	} else {
            	continue
            	}
            	if prediction.dewPoint != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.dewPoint.Value))
            		name = "dewpoint"
            	} else {
            	continue
            	}
            	if prediction.humidity != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.humidity.Value))
            		name = "humidity"
            	} else {
            	continue
            	}
            	if prediction.pressure != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.pressure.Value))
            		name = "pressure"
            	} else {
            	continue
            	}
            	if prediction.windSpeed != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.windSpeed.Value))
            		name = "windspeed"
            	} else {
            	continue
            	}
            	if prediction.windGust != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.windGust.Value))
            		name = "windgust"
            	} else {
            	continue
            	}
            	if prediction.windBearing != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.windBearing.Value))
            		name = "windbearing"
            	} else {
            	continue
            	}
            	if prediction.cloudCover != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.cloudCover.Value))
            		name = "cloudcover"
            	} else {
            	continue
            	}
            	if prediction.uvIndex != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.uvIndex.Value))
            		name = "uvindex"
            	} else {
            	continue
            	}
            	if prediction.visibility != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.visibility.Value))
            		name = "visibility"
            	} else {
            	continue
            	}
            	if prediction.ozone != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.ozone.Value))
            		name = "ozone"
            	} else {
            	continue
            	}


				extracted.UUID = types.GenerateUUID(uri, []byte(name))
				extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
				extracted.Tags = map[string]string{
					"unit":            device_units[name],
					"name":            name,
					"prediction_step": fmt.Sprintf("%d", step),
				}
				extracted.IntTags = map[string]int64{
					"prediction_time": int64(_prediction.PredictionTime),
				}
				if err := add(extracted); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
