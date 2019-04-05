from pyxbos import *
import os,sys
import json
import requests
import yaml
import argparse


class WeatherCurrentDriver(Driver):
    def setup(self, cfg):
        self.baseurl = cfg['darksky']['url']
        self.apikey = cfg['darksky']['apikey']
        self.coords = cfg['darksky']['coordinates']
        self.url = self.baseurl + self.apikey + '/' + self.coords

    def read(self, requestid=None):
        print("In current driver")
        response = requests.get(self.url)
        json_data = json.loads(response.text)
        if 'currently' not in json_data: return

        logging.info("currently {0}".format(json_data['currently']))
        print(json_data['currently'])
        nearestStormDistance =  json_data['currently'].get('nearestStormDistance',None)
        nearestStormBearing =   json_data['currently'].get('nearestStormBearing',None)
        precipIntensity =       json_data['currently'].get('precipIntensity',None)
        apparentTemperature =   json_data['currently'].get('apparentTemperature',None)
        humidity =              json_data['currently'].get('humidity',None)
        #print(json_data['currently'])
        if humidity is not None:
            humidity *= 100 # change from decimal to percent

        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                weather_station = iot_pb2.WeatherStation(
                    time  =   types.Int64(value=output['time']),
                    icon  =  output['icon'],
                    precipIntensity  =   types.Double(value=output['precipIntensity']),
                    precipIntensityError  =   types.Double(value=output['precipIntensityError']),
                    precipProbability  =   types.Double(value=output['precipProbability']),
                    precipAccumulation  =   types.Double(value=output['precipAccumulation']),
                    precipType  =  output['precipType'],
                    temperature  =   types.Double(value=output['temperature']),
                    apparentTemperature  =   types.Double(value=output['apparentTemperature']),
                    dewPoint  =   types.Double(value=output['dewPoint']),
                    humidity  =   types.Double(value=output['humidity']),
                    pressure  =   types.Double(value=output['pressure']),
                    windSpeed  =   types.Double(value=output['windSpeed']),
                    windGust  =   types.Double(value=output['windGust']),
                    windBearing  =   types.Double(value=output['windBearing']),
                    cloudCover  =   types.Double(value=output['cloudCover']),
                    uvIndex  =   types.Double(value=output['uvIndex']),
                    visibility  =   types.Double(value=output['visibility']),
                    ozone  =   types.Double(value=output['ozone']),
                )
            )
        )
        self.report(self.coords, msg)


if __name__ == '__main__':
    with open('dark_sky.yaml') as f:
        # use safe_load instead load for security reasons
        driverConfig = yaml.safe_load(f)

    namespace = driverConfig['wavemq']['namespace']
    api = driverConfig['dark_sky']['api']
    cfg = {
        'darksky': {
            'apikey': api,
            'url': 'https://api.darksky.net/forecast/',
            'coordinates': '40.5301,-124.0000' # Should be near BLR
        },
        'wavemq': 'localhost:4516',
        'namespace': namespace,
        'base_resource': 'weather_current',
        'entity': 'weather_current.ent',
        'id': 'pyxbos-driver-current-1',
        #'rate': 1800, # half hour
        'rate': 20, # 15 min
    }
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    current_driver = WeatherCurrentDriver(cfg)
    current_driver.begin()
