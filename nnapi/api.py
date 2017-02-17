from flask import Flask
from flask_restful import Resource, Api, abort, reqparse, fields
import os
import netifaces

app = Flask(__name__)
api = Api(app)

class LinuxNetDev:
    """ Class to access NetDev info on Linux. """
    def __init__(self):
        self.name = ""

    def abort_if_netdev_id_doesnt_exist(self, netdev_id):
        if netdev_id not in netifaces.interfaces():
            abort(404, message="netdev_id doesn't exist")
    
    def get_stats(self, netdev_id):
        with open('/proc/net/dev') as self.fd:
            self.lines = self.fd.readlines()
            _, self.rx_cols, self.tx_cols = self.lines[1].rstrip().split('|')
            self.rx_col = self.rx_cols.split()
            self.tx_col = self.tx_cols.split()
            for self.line in self.lines[2:]:
                if self.line.find(netdev_id) < 0: continue
                self.iface, self.data = self.line.split(':')
                self.stats = self.data.split()
        self.i = 0
        self.stats_pairs = {}
        for self.col_name in self.rx_col:
            self.stats_pairs["rx-" + self.col_name] = int(self.stats[self.i])
            self.i = self.i + 1
        for self.col_name in self.tx_col:
            self.stats_pairs["tx-" + self.col_name] = int(self.stats[self.i])
            self.i = self.i + 1
        return self.stats_pairs
    
    def get_addr_list(self, netdev_id):
        self.ifdata = netifaces.ifaddresses(netdev_id)
        self.addr_list = {}
        self.addr_list['mac'] = self.ifdata.get(netifaces.AF_LINK)
        self.addr_list['ipv4'] = self.ifdata.get(netifaces.AF_INET)
        self.addr_list['ipv6'] = self.ifdata.get(netifaces.AF_INET6)
        return self.addr_list

class RootDir(Resource):
    def get(self):
        return {'hello': 'world'}

class Date(Resource):
    def get(self):
        f = os.popen('date +"%Y/%m/%d"')
        now = f.read()
        return {'Date': now}

class NetDevList(Resource):
    def get(self):
        return netifaces.interfaces()

class NetDev(Resource):
    def get(self, netdev_id):
        l = LinuxNetDev()
        l.abort_if_netdev_id_doesnt_exist(netdev_id)
        #main
        return_fields = {}
        return_fields['address'] = l.get_addr_list(netdev_id)
        return_fields['stats'] = l.get_stats(netdev_id)
        return return_fields

class NetDevAddressList(Resource):
    def get(self, netdev_id):
        l = LinuxNetDev()
        l.abort_if_netdev_id_doesnt_exist(netdev_id)
        return l.get_addr_list(netdev_id)
        
class NetDevAddress(Resource):
    def get(self, netdev_id, addr_type):
        l = LinuxNetDev()
        l.abort_if_netdev_id_doesnt_exist(netdev_id)
        addr_list = l.get_addr_list(netdev_id)
        
        if addr_type not in addr_list:
            abort(404, message="addr_type doesn't exist")
        return addr_list[addr_type]

class NetDevStats(Resource):
    def get(self, netdev_id):
        l = LinuxNetDev()
        l.abort_if_netdev_id_doesnt_exist(netdev_id)
        
        return l.get_stats(netdev_id)

api.add_resource(RootDir, '/')
api.add_resource(Date, '/date')
api.add_resource(NetDevList, '/netdev/')
api.add_resource(NetDev, '/netdev/<netdev_id>/')
api.add_resource(NetDevAddressList, '/netdev/<netdev_id>/address/')
api.add_resource(NetDevAddress, '/netdev/<netdev_id>/address/<addr_type>/')
api.add_resource(NetDevStats, '/netdev/<netdev_id>/stats/')

if __name__ == '__main__':
    app.run(debug=True)

