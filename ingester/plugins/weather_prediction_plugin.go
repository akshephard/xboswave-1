package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
    //types "/home/solarplus/xboswave-1/ingester/types"
    //"github.com/akshephard/xboswave-1/ingester/types"
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
            var step int
            step = 1
			for _, _prediction := range msg.XBOSIoTDeviceState.WeatherPrediction.Predictions {
				prediction := _prediction.Prediction
				var extracted types.ExtractedTimeseries
				var name string
				time := int64(msg.XBOSIoTDeviceState.Time)


                // this is subtracting the the current xbos time from each prediction time
				//step := (int64(_prediction.PredictionTime) - time) / 1e9

                // Use for debugging
                fmt.Printf("The xbos time in seconds is: %d\n", time/ 1e9)
                fmt.Printf("The step is: %d\n", step)
                fmt.Printf("The prediction time is: %d\n", (int64(_prediction.PredictionTime) / 1e9))
                fmt.Printf("The XBOSIoTDeviceState.Time is: %d\n", (int64(msg.XBOSIoTDeviceState.Time) / 1e9))
                //fmt.Printf("The count is: %d\n", count)
                fmt.Printf("The temperature is: %f\n", float64(prediction.Temperature.Value))


                //This is the time that is being put into influx as the timestamp
				extracted.Times = append(extracted.Times, time)
            	if prediction.PrecipIntensity != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.PrecipIntensity.Value))
            		name = "precipintensity"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.PrecipIntensityError != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.PrecipIntensityError.Value))
            		name = "precipintensityerror"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.PrecipProbability != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.PrecipProbability.Value))
            		name = "precipprobability"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.Temperature != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.Temperature.Value))
            		name = "temperature"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.ApparentTemperature != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.ApparentTemperature.Value))
            		name = "apparenttemperature"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.DewPoint != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.DewPoint.Value))
            		name = "dewpoint"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.Humidity != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.Humidity.Value))
            		name = "humidity"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.Pressure != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.Pressure.Value))
            		name = "pressure"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.WindSpeed != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.WindSpeed.Value))
            		name = "windspeed"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.WindGust != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.WindGust.Value))
            		name = "windgust"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.WindBearing != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.WindBearing.Value))
            		name = "windbearing"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.CloudCover != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.CloudCover.Value))
            		name = "cloudcover"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.UvIndex != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.UvIndex.Value))
            		name = "uvindex"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.Visibility != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.Visibility.Value))
            		name = "visibility"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
            	if prediction.Ozone != nil {
            		extracted.Values = append(extracted.Values, float64(prediction.Ozone.Value))
            		name = "ozone"
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}

                step++
                //fmt.Println("The final extraced wave message is %d", extracted.IntTags["prediction_time"])
			}
		}
	}
	return nil
}
