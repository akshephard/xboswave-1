from pyxbos.driver import *
from pyxbos import wattnode_pb2
from modbus_driver import Modbus_Driver
import os,sys
import json
import requests
import yaml
import argparse
import time
from inspect import getmembers


class WattnodeDriver(Driver):
    def setup(self, cfg):
        with open('wattnode.yaml') as f:
            # use safe_load instead load for security reasons
            driverConfig = yaml.safe_load(f)
        self.modbus_device = Modbus_Driver("wattnode.yaml")
        self.modbus_device.initialize_modbus()
        self.service_name = cfg['service_name']

    def read(self, requestid=None):
        output = self.modbus_device.get_data()
        print("Do we get here?")
        print(output)
        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                wattnode = wattnode_pb2.WattnodeState(
                    EnergySum  =   types.Double(value=output.get('EnergySum',None)),
                    EnergyPosSUm  =   types.Double(value=output.get('EnergyPosSUm',None)),
                    EnergySumNR  =   types.Double(value=output.get('EnergySumNR',None)),
                    EnergyPosSumNr  =   types.Double(value=output.get('EnergyPosSumNr',None)),
                    PowerSum  =   types.Double(value=output.get('PowerSum',None)),
                    PowerA  =   types.Double(value=output.get('PowerA',None)),
                    PowerB  =   types.Double(value=output.get('PowerB',None)),
                    PowerC  =   types.Double(value=output.get('PowerC',None)),
                    VoltAvgLN  =   types.Double(value=output.get('VoltAvgLN',None)),
                    VoltA  =   types.Double(value=output.get('VoltA',None)),
                    VoltB  =   types.Double(value=output.get('VoltB',None)),
                    VoltC  =   types.Double(value=output.get('VoltC',None)),
                    VoltAvgLL  =   types.Double(value=output.get('VoltAvgLL',None)),
                    VoltAB  =   types.Double(value=output.get('VoltAB',None)),
                    VoltBC  =   types.Double(value=output.get('VoltBC',None)),
                    VoltAC  =   types.Double(value=output.get('VoltAC',None)),
                    Freq  =   types.Double(value=output.get('Freq',None)),
                    EnergyA  =   types.Double(value=output.get('EnergyA',None)),
                    EnergyB  =   types.Double(value=output.get('EnergyB',None)),
                    EnergyC  =   types.Double(value=output.get('EnergyC',None)),
                    EnergyPosA  =   types.Double(value=output.get('EnergyPosA',None)),
                    EnergyPosB  =   types.Double(value=output.get('EnergyPosB',None)),
                    EnergyPosC  =   types.Double(value=output.get('EnergyPosC',None)),
                    EnergyNegSum  =   types.Double(value=output.get('EnergyNegSum',None)),
                    EnergyNegSumNR  =   types.Double(value=output.get('EnergyNegSumNR',None)),
                    EnergyNegA  =   types.Double(value=output.get('EnergyNegA',None)),
                    EnergyNegB  =   types.Double(value=output.get('EnergyNegB',None)),
                    EnergyNegC  =   types.Double(value=output.get('EnergyNegC',None)),
                    EnergyReacSum  =   types.Double(value=output.get('EnergyReacSum',None)),
                    EnergyReacA  =   types.Double(value=output.get('EnergyReacA',None)),
                    EnergyReacB  =   types.Double(value=output.get('EnergyReacB',None)),
                    EnergyReacC  =   types.Double(value=output.get('EnergyReacC',None)),
                    EnergyAppSum  =   types.Double(value=output.get('EnergyAppSum',None)),
                    EnergyAppA  =   types.Double(value=output.get('EnergyAppA',None)),
                    EnergyAppB  =   types.Double(value=output.get('EnergyAppB',None)),
                    EnergyAppC  =   types.Double(value=output.get('EnergyAppC',None)),
                    PowerFactorAvg  =   types.Double(value=output.get('PowerFactorAvg',None)),
                    PowerFactorA  =   types.Double(value=output.get('PowerFactorA',None)),
                    PowerFactorB  =   types.Double(value=output.get('PowerFactorB',None)),
                    PowerFactorC  =   types.Double(value=output.get('PowerFactorC',None)),
                    PowerReacSum  =   types.Double(value=output.get('PowerReacSum',None)),
                    PowerReacA  =   types.Double(value=output.get('PowerReacA',None)),
                    PowerReacB  =   types.Double(value=output.get('PowerReacB',None)),
                    PowerReacC  =   types.Double(value=output.get('PowerReacC',None)),
                    PowerAppSum  =   types.Double(value=output.get('PowerAppSum',None)),
                    PowerAppA  =   types.Double(value=output.get('PowerAppA',None)),
                    PowerAppB  =   types.Double(value=output.get('PowerAppB',None)),
                    PowerAppC  =   types.Double(value=output.get('PowerAppC',None)),
                    CurrentA  =   types.Double(value=output.get('CurrentA',None)),
                    CurrentB  =   types.Double(value=output.get('CurrentB',None)),
                    CurrentC  =   types.Double(value=output.get('CurrentC',None)),
                    Demand  =   types.Double(value=output.get('Demand',None)),
                    DemandMin  =   types.Double(value=output.get('DemandMin',None)),
                    DemandMax  =   types.Double(value=output.get('DemandMax',None)),
                    DemandApp  =   types.Double(value=output.get('DemandApp',None)),
                    DemandA  =   types.Double(value=output.get('DemandA',None)),
                    DemandB  =   types.Double(value=output.get('DemandB',None)),
                    DemandC  =   types.Double(value=output.get('DemandC',None))
                )
            )
        )
        print(self.report(self.service_name, msg))


if __name__ == '__main__':

    # Need to have arguments for config file and entity for systemd
    parser = argparse.ArgumentParser()
    parser.add_argument("config_file", help="config file with api key as well as namespace")
    parser.add_argument("ent_file", help="entity file")
    args = parser.parse_args()
    config_file = args.config_file
    ent_file = args.ent_file

    with open(config_file) as f:
        # use safe_load instead load for security reasons
        driverConfig = yaml.safe_load(f)

    namespace = driverConfig['wavemq']['namespace']
    service_name = driverConfig['xbos']['service_name']
    #driver_cfg = "parker.yaml"
    print(driverConfig)

    xbos_cfg = {
        'wavemq': 'localhost:4516',
        'namespace': namespace,
        'base_resource': 'wattnode',
        'entity': ent_file,
        'id': 'pyxbos-driver-wattnode',
        #'rate': 1800, # half hour
        'rate': 20, # 15 min
        'service_name': service_name
    }
    print(getmembers(iot_pb2))
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    #e = DarkSkyDriver(cfg)
    e = WattnodeDriver(xbos_cfg)
    e.begin()
