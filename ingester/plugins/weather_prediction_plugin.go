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
                var extract_map map[string]types.ExtractedTimeseries
                //extract_map = make(map[string]types.ExtractedTimeseries)
                extracted_slice := make([]types.ExtractedTimeseries, 3)
                //var extracted_slice []types.ExtractedTimeseries
                var name_list []string
				time := int64(msg.XBOSIoTDeviceState.Time)
				step := (int64(_prediction.PredictionTime) - time) / 1e9
				extracted.Times = append(extracted.Times, time)


            	if prediction.Visibility != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Visibility.Value))
                        //extracted_slice = append(extracted_slice, extracted.Values)
                        extract_map["visibility"] = extracted.Values
                        name_list = append(name_list, "visibility")
            		name = "visibility"
            	} else {
            	continue
            	}
                for i, current_name := range name_list {
                    fmt.Println(i, s)
    				extracted_slice[i].UUID = types.GenerateUUID(uri, []byte(name_list[i]))
    				extracted_slice[i].Collection = fmt.Sprintf("xbos/%s", uri.Resource)
    				extracted_slice[i].Tags = map[string]string{
    					"unit":            device_units[current_name],
    					"name":            current_name,
    					"prediction_step": fmt.Sprintf("%d", step),
    				}
    				extracted_slice[i].IntTags = map[string]int64{
    					"prediction_time": int64(_prediction.PredictionTime),
    				}
    				if err := add(extracted_slice[i]); err != nil {
    					return err
    				}
                }
                /*
                extracted_slice = append(extracted_slice, extracted)

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


            	if prediction.Ozone != nil {
            			extracted.Values = append(extracted.Values, float64(prediction.Ozone.Value))
            		name = "ozone"
            	} else {
            	continue
            	}
                extracted_slice = append(extracted_slice, extracted)
                //Add the extracted values into some type of array and then iterate through loop
                //You need to keep track of the name in some kind of array as well so they can be done in an order

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
                */
			}
		}
	}
	return nil
}
