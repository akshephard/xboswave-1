from pyxbos import *
from modbus_driver import Modbus_Driver
import os,sys
import json
import requests
import yaml
import argparse
import time
from inspect import getmembers


class ParkerDriver(Driver):
    def setup(self, cfg):
        with open('parker.yaml') as f:
            # use safe_load instead load for security reasons
            driverConfig = yaml.safe_load(f)
        self.modbus_device = Modbus_Driver(driverConfig)
        self.modbus_device.initialize_modbus()
        self.service_name = cfg['service_name']

    def read(self, requestid=None):
        output = self.modbus_device.get_data()
        HAACP_0_0 = output['HAACP_0_0']
        HAACP_0_1 = output['HAACP_0_1']
        HAACP_0_2 = output['HAACP_0_2']
        compressor_working_hours = output['compressor_working_hours']
        defrost_control = output['defrost_control']
        status_on = output['status_on']
        status_light = output['status_light']
        status_aux_output = output['status_aux_output']
        next_defrost_counter = output['next_defrost_counter']
        digital_input_output_status = output['digital_input_output_status']
        cab_probe = output['cab_probe']
        evap_probe = output['evap_probe']
        aux_probe = output['aux_probe']
        alarm_status = output['alarm_status']
        regulator_flag1 = output['regulator_flag1']
        regulator_flag2 = output['regulator_flag2']
        active_set_point = output['active_set_point']
        time_remaining_next_defrost = output['time_remaining_next_defrost']
        current_defrost_counter = output['current_defrost_counter']
        compressor_delay = output['compressor_delay']
        num_HAACP_alarm_historial = output['num_HAACP_alarm_historial']
        digital_output_flag = output['digital_output_flag']
        SP = output['SP']
        CA1 = output['CA1']
        CA2 = output['CA2']
        CA3 = output['CA3']
        p0 = output['p0']
        p1 = output['p1']
        p2 = output['p2']
        p3 = output['p3']
        p4 = output['p4']
        p5 = output['p5']
        p6 = output['p6']
        r0 = output['r0']
        r1 = output['r1']
        r2 = output['r2']
        r3 = output['r3']
        r4 = output['r4']
        C0 = output['c0']
        C1 = output['c1']
        c2 = output['c2']
        c3 = output['c3']
        c4 = output['c4']
        c5 = output['c5']
        c6 = output['c6']
        c7 = output['c7']
        c8 = output['c8']
        c9 = output['c9']
        c10 = output['c10']
        d0 = output['d0']
        d1 = output['d1']
        d2 = output['d2']
        d3 = output['d3']
        d4 = output['d4']
        d5 = output['d5']
        d6 = output['d6']
        d7 = output['d7']
        d8 = output['d8']
        d9 = output['d9']
        dA = output['dA']
        A0 = output['A0']
        A1 = output['A1']
        A2 = output['A2']
        A3 = output['A3']
        A4 = output['A4']
        A5 = output['A5']
        A6 = output['A6']
        A7 = output['A7']
        A8 = output['A8']
        A9 = output['A9']
        AA = output['AA']
        F0 = output['F0']
        F1 = output['F1']
        F2 = output['F2']
        F3 = output['F3']
        i0 = output['i0']
        i1 = output['i1']
        i2 = output['i2']
        i3 = output['i3']
        i4 = output['i4']
        i5 = output['i5']
        i6 = output['i6']
        i7 = output['i7']
        i8 = output['i8']
        i9 = output['i9']
        u1 = output['u1']
        u2 = output['u2']
        u3 = output['u3']
        u4 = output['u4']
        u5 = output['u5']
        u6 = output['u6']
        u7 = output['u7']
        u8 = output['u8']
        HE1 = output['HE1']
        HE2 = output['HE2']
        hd1 = output['hd1']
        hd2 = output['hd2']
        hd3 = output['hd3']
        hd4 = output['hd4']
        hd5 = output['hd5']
        hd6 = output['hd6']

        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                parker_state = iot_pb2.ParkerState(
                    compressor_working_hours  =   types.Double(value=compressor_working_hours),
                    defrost_control  =   types.Int64(value=defrost_control),
                    r1  =   types.Double(value=r1),
                    r2  =   types.Double(value=r2),
                    r4  =   types.Double(value=r4),
                    C0  =   types.Double(value=C0),
                    C1  =   types.Double(value=C1),
                    d0  =   types.Int64(value=d0),
                    d3  =   types.Double(value=d3),
                    d5  =   types.Double(value=d5),
                    d7  =   types.Double(value=d7),
                    d8  =   types.Int64(value=d8),
                    A0  =   types.Int64(value=A0),
                    A1  =   types.Double(value=A1),
                    A2  =   types.Int64(value=A2),
                    A3  =   types.Int64(value=A3),
                    A4  =   types.Double(value=A4),
                    A5  =   types.Int64(value=A5),
                    A6  =   types.Double(value=A6),
                    A7  =   types.Double(value=A7),
                    A8  =   types.Double(value=A8),
                    A9  =   types.Double(value=A9),
                    F0  =   types.Int64(value=F0),
                    F1  =   types.Double(value=F1),
                    F2  =   types.Int64(value=F2),
                    F3  =   types.Double(value=F3),
                    Hd1  =   types.string(value=Hd1),
                    Hd2  =   types.string(value=Hd2),
                    Hd3  =   types.string(value=Hd3),
                    Hd4  =   types.string(value=Hd4),
                    Hd5  =   types.string(value=Hd5),
                    Hd6  =   types.string(value=Hd6)
                )
            )
        )
        self.report(self.service_name, msg)


if __name__ == '__main__':
    with open('parker.yaml') as f:
        # use safe_load instead load for security reasons
        driverConfig = yaml.safe_load(f)

    namespace = driverConfig['wavemq']['namespace']
    service_name = driverConfig['xbos']['service_name']
    #driver_cfg = "parker.yaml"
    print(driverConfig)

    xbos_cfg = {
        'wavemq': 'localhost:4516',
        'namespace': namespace,
        'base_resource': 'test/darksky',
        'entity': 'parker.ent',
        'id': 'pyxbos-driver-darksky-1',
        #'rate': 1800, # half hour
        'rate': 900, # 15 min
        'service_name': service_name
    }
    print(getmembers(iot_pb2))
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    #e = DarkSkyDriver(cfg)
    e = ParkerDriver(xbos_cfg)
    e.begin()
