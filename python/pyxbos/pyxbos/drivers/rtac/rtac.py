from pyxbos.driver import *
from modbus_driver import Modbus_Driver
import os,sys
import json
import requests
import yaml
import argparse
import time
from inspect import getmembers


class rtacDriver(Driver):
    def setup(self, cfg):
        with open('rtac.yaml') as f:
            # use safe_load instead load for security reasons
            driverConfig = yaml.safe_load(f)
        self.modbus_device = Modbus_Driver(driverConfig)
        self.modbus_device.initialize_modbus()
        self.service_name = cfg['service_name']

    def read(self, requestid=None):
        output = self.modbus_device.get_data()

        print(output)
        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                rtac_state = parker_pb2.RtacState(
                    island_state  =   types.Bool(value=output.get('island_state',None)),
                    island_type  =   types.Bool(value=output.get('island_type',None)),
                    bess_availability  =   types.Bool(value=output.get('bess_availability',None)),
                    pge_state  =   types.Bool(value=output.get('pge_state',None)),
                    pcc_breaker_state  =   types.Bool(value=output.get('pcc_breaker_state',None)),
                    bess_pv_breaker_state  =   types.Bool(value=output.get('bess_pv_breaker_state',None)),
                    heartbeat  =   types.Int64(value=output.get('heartbeat',None)),
                    real_power_setpoint  =   types.Double(value=output.get('real_power_setpoint',None)),
                    reactive_power_setpoint  =   types.Double(value=output.get('reactive_power_setpoint',None)),
                    target_real_power  =   types.Double(value=output.get('target_real_power',None)),
                    target_reactive_power  =   types.Double(value=output.get('target_reactive_power',None)),
                    battery_total_capacity  =   types.Double(value=output.get('battery_total_capacity',None)),
                    battery_current_stored_energy  =   types.Double(value=output.get('battery_current_stored_energy',None)),
                    total_actual_real_power  =   types.Double(value=output.get('total_actual_real_power',None)),
                    total_actual_reactive_power  =   types.Double(value=output.get('total_actual_reactive_power',None)),
                    total_actual_apparent_power  =   types.Double(value=output.get('total_actual_apparent_power',None)),
                    active_power_output_limit  =   types.Double(value=output.get('active_power_output_limit',None)),
                    current_power_production  =   types.Double(value=output.get('current_power_production',None)),
                    ac_current_phase_a  =   types.Double(value=output.get('ac_current_phase_a',None)),
                    ac_current_phase_b  =   types.Double(value=output.get('ac_current_phase_b',None)),
                    ac_current_phase_c  =   types.Double(value=output.get('ac_current_phase_c',None)),
                    ac_voltage_ab  =   types.Double(value=output.get('ac_voltage_ab',None)),
                    ac_voltage_bc  =   types.Double(value=output.get('ac_voltage_bc',None)),
                    ac_voltage_ca  =   types.Double(value=output.get('ac_voltage_ca',None)),
                    ac_frequency  =   types.Double(value=output.get('ac_frequency',None)),
                    fault_condition  =   types.Int64(value=output.get('fault_condition',None)),
                    pge_voltage  =   types.Double(value=output.get('pge_voltage',None)),
                    pge_frequency  =   types.Double(value=output.get('pge_frequency',None)),
                )
            )
        )
        print(self.report(self.service_name, msg))

if __name__ == '__main__':
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
        'base_resource': 'rtac',
        'entity': ent_file,
        'id': 'pyxbos-driver-rtac-1',
        #'rate': 1800, # half hour
        'rate': 20, # 15 min
        'service_name': service_name
    }
    print(getmembers(iot_pb2))
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    #e = DarkSkyDriver(cfg)
    e = rtacDriver(xbos_cfg)
    e.begin()
