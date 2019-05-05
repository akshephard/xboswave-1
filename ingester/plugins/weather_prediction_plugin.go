package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)

func has_device(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.WeatherPrediction != nil
}

type add_fn func(types.ExtractedTimeseries) error


// This contains the mapping of each field's value to the unit
var device_units = map[string]string{
	"time":	"seconds",
	"icon":	"",
	"precipIntensity":	"inches per hour",
	"precipIntensityError":	"",
	"precipProbability":	"",
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

func send_time_series_to_influx(value float64,name string,toInflux types.ExtractedTimeseries,
    pass_add add_fn,prediction_time int64, step int, uri types.SubscriptionURI) error{
	toInflux.Values = append(toInflux.Values, value)

    //This UUID is unique to each field in the message
	toInflux.UUID = types.GenerateUUID(uri, []byte(name))
    //The collection comes from the resource name of the driver
	toInflux.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
    //These are the tags that will be used when the point is written
	toInflux.Tags = map[string]string{
		"unit":            device_units[name],
		"name":            name,
		"prediction_time": fmt.Sprintf("%d", prediction_time / 1e9),
		"prediction_step": fmt.Sprintf("%d", step),
	}
    //This add function is passed in from the ingester and when it is executed
    //a point is written into influx
	if err := pass_add(toInflux); err != nil {
		return err
	}
    return nil
}


func Extract(uri types.SubscriptionURI, msg xbospb.XBOS, add func(types.ExtractedTimeseries) error) error {
	if msg.XBOSIoTDeviceState != nil {

		if has_device(msg) {
            var step int
            step = 1

            //Iterate through each hour of prediction from current to 48 hours from current
			for _, _prediction := range msg.XBOSIoTDeviceState.WeatherPrediction.Predictions {
                //This prediction contains all of the fields that were present in WeatherCurrent message
                //There is one for each hour that is retrieved from the DarkSky API
				prediction := _prediction.Prediction

                //This will contain all the information necessary to send one prediction for one hour out of 0-48
				var extracted types.ExtractedTimeseries
				var name string

                //This is the xbos time
				time := int64(msg.XBOSIoTDeviceState.Time)

                // this is subtracting the the current xbos time from each prediction time
				//step := (int64(_prediction.PredictionTime) - time) / 1e9

                // Use for debugging
                /*
                fmt.Printf("The xbos time in seconds is: %d\n", time/ 1e9)
                fmt.Printf("The step is: %d\n", step)
                fmt.Printf("The prediction time is: %d\n", (int64(_prediction.PredictionTime) / 1e9))
                fmt.Printf("The XBOSIoTDeviceState.Time is: %d\n", (int64(msg.XBOSIoTDeviceState.Time) / 1e9))
                fmt.Printf("The temperature is: %f\n", float64(prediction.Temperature.Value))
                */

                //This is the time that is being put into influx as the timestamp
				extracted.Times = append(extracted.Times, time)

                //func send_time_series_to_influx(value float64,name string,toInflux types.ExtractedTimeseries,
                //    add func(types.ExtractedTimeseries),prediction_time int64, step int64){

                if prediction.PrecipIntensity != nil {
                    send_time_series_to_influx(float64(prediction.PrecipIntensity.Value),
                    name,extracted, add,int64(_prediction.PredictionTime),step, uri)
                }
                /*
            	if prediction.PrecipIntensity != nil {
                    //This will be the value that is put into a field in Influx
            		extracted.Values = append(extracted.Values, float64(prediction.PrecipIntensity.Value))
            		name = "precipintensity"
                    //This UUID is unique to each field in the message
            		extracted.UUID = types.GenerateUUID(uri, []byte(name))
                    //The collection comes from the resource name of the driver
            		extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
                    //These are the tags that will be used when the point is written
            		extracted.Tags = map[string]string{
            			"unit":            device_units[name],
            			"name":            name,
            			"prediction_time": fmt.Sprintf("%d", int64(_prediction.PredictionTime) / 1e9),
            			"prediction_step": fmt.Sprintf("%d", step),
            		}
                    //This add function is passed in from the ingester and when it is executed
                    //a point is written into influx
            		if err := add(extracted); err != nil {
            			fmt.Println("Are there any errors?")
            			fmt.Println(err)
            			return err
            		}
            	}
                */
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
                /*
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
                */
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
