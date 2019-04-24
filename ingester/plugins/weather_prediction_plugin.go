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
		// proof of concept
		// TODO: finish
		if has_device(msg) {
            var count int
			for _, _prediction := range msg.XBOSIoTDeviceState.WeatherPrediction.Predictions {
				prediction := _prediction.Prediction
				var extracted types.ExtractedTimeseries
				var name string
				time := int64(msg.XBOSIoTDeviceState.Time)
				step := (int64(_prediction.PredictionTime) - time) / 1e9
                fmt.Printf("The step is: %d\n", step)
                fmt.Printf("The prediction time is: %d\n", int64(_prediction.PredictionTime))
                fmt.Printf("The XBOSIoTDeviceState.Time is: %d\n", int64(msg.XBOSIoTDeviceState.Time))
                fmt.Printf("The count is: %d\n", count)
                fmt.Printf("The temperature is: %f\n", float64(prediction.Temperature.Value))
                count++
				extracted.Times = append(extracted.Times, time)
				if prediction.Temperature != nil {
					extracted.Values = append(extracted.Values, float64(prediction.Temperature.Value))
					name = "temperature"
				} else {
					continue
				}
				extracted.UUID = types.GenerateUUID(uri, []byte(name))
				extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
				extracted.Tags = map[string]string{
					"unit":            device_units[name],
					"name":            name,
					"prediction_step": fmt.Sprintf("%d", step),
                    "test_time:" int64(_prediction.PredictionTime),
				}
				extracted.IntTags = map[string]int64{
					"prediction_time": int64(_prediction.PredictionTime),
				}
				if err := add(extracted); err != nil {
                    fmt.Println("Are there any error?")
                    fmt.Println(err)
					return err
				}
                fmt.Println("The final extraced wave message is %d", extracted.IntTags["prediction_time"])
			}
		}
	}
	return nil
}
