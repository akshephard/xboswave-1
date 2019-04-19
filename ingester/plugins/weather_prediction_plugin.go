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

            	if prediction.Time != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Time.Value))
            		name = "time"
            	} else {
            	continue
            	}
            	if prediction.PrecipIntensity != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.PrecipIntensity.Value))
            		name = "precipintensity"
            	} else {
            	continue
            	}
            	if prediction.PrecipIntensityError != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.PrecipIntensityError.Value))
            		name = "precipintensityerror"
            	} else {
            	continue
            	}
            	if prediction.PrecipProbability != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.PrecipProbability.Value))
            		name = "precipprobability"
            	} else {
            	continue
            	}
            	if prediction.Temperature != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Temperature.Value))
            		name = "temperature"
            	} else {
            	continue
            	}
            	if prediction.ApparentTemperature != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.ApparentTemperature.Value))
            		name = "apparenttemperature"
            	} else {
            	continue
            	}
            	if prediction.DewPoint != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.DewPoint.Value))
            		name = "dewpoint"
            	} else {
            	continue
            	}
            	if prediction.Humidity != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Humidity.Value))
            		name = "humidity"
            	} else {
            	continue
            	}
            	if prediction.Pressure != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Pressure.Value))
            		name = "pressure"
            	} else {
            	continue
            	}
            	if prediction.WindSpeed != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.WindSpeed.Value))
            		name = "windspeed"
            	} else {
            	continue
            	}
            	if prediction.WindGust != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.WindGust.Value))
            		name = "windgust"
            	} else {
            	continue
            	}
            	if prediction.WindBearing != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.WindBearing.Value))
            		name = "windbearing"
            	} else {
            	continue
            	}
            	if prediction.CloudCover != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.CloudCover.Value))
            		name = "cloudcover"
            	} else {
            	continue
            	}
            	if prediction.UvIndex != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.UvIndex.Value))
            		name = "uvindex"
            	} else {
            	continue
            	}
            	if prediction.Visibility != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Visibility.Value))
            		name = "visibility"
            	} else {
            	continue
            	}
            	if prediction.Ozone != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Ozone.Value))
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
