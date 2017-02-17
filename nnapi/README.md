# ABOUT nnapi
Network Node API, implementing REST like API for network nodes.

# OS / Tools / Framework
* OS : Ubuntu 15.10 (should work on other versions as well)
* flask-restful
  * Python REST framework
* python virtualenv
  * Isolating test environment
* python netifaces
  * libraly to fetch network info from OS (Linux)

# QUICK START

First time:
```shell
$ sudo apt-get install python-virtualenv
$ git clone <nnapi-repo-url>
$ cd nnapi
~/nnapi$ virtualenv venv
~/nnapi$ ls
LICENSE  README.md  venv
~/nnapi$ . venv/bin/activate
(venv) ~/nnapi$ pip install Flask
(venv) ~/nnapi$ pip install flask-restful
(venv) ~/nnapi$ pip install netifaces
```

Running REST server
```shell
$ cd nnapi
~/nnapi$ . venv/bin/activate
(venv) ~/nnapi$ python ./api.py
 * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
 * Restarting with stat
 * Debugger is active!
 * Debugger pin code: 723-124-226
```

Accessing REST server
```shell
# List netdev.
$ curl http://127.0.0.1:5000/netdev/
# Get interface info (address + stats)
$ curl http://127.0.0.1:5000/netdev/<ifname>/
# Get address of an interface
$ curl http://127.0.0.1:5000/netdev/<ifname>/address/
# Get only IPv4 address of an interface
$ curl http://127.0.0.1:5000/netdev/<ifname>/address/ipv4/
# Get only stats of an interface
$ curl http://127.0.0.1:5000/netdev/<ifname>/stats/
```

## Sample output
```shell
$ curl http://127.0.0.1:5000/netdev/
[
    "lo",
    "enp0s3",
    "virbr0",
    "virbr0-nic"
]

$ curl http://127.0.0.1:5000/netdev/enp0s3/
{
    "address": {
        "ipv4": [
            {
                "addr": "10.0.2.15",
                "broadcast": "10.0.2.255",
                "netmask": "255.255.255.0"
            }
        ],
        "ipv6": [
            {
                "addr": "fe80::a00:27ff:feaf:dc60%enp0s3",
                "netmask": "ffff:ffff:ffff:ffff::"
            }
        ],
        "mac": [
            {
                "addr": "08:00:27:af:dc:60",
                "broadcast": "ff:ff:ff:ff:ff:ff"
            }
        ]
    },
    "stats": {
        "rx-bytes": 11324896,
        "rx-compressed": 0,
        "rx-drop": 0,
        "rx-errs": 0,
        "rx-fifo": 0,
        "rx-frame": 0,
        "rx-multicast": 0,
        "rx-packets": 53451,
        "tx-bytes": 8512050,
        "tx-carrier": 0,
        "tx-colls": 0,
        "tx-compressed": 0,
        "tx-drop": 0,
        "tx-errs": 0,
        "tx-fifo": 0,
        "tx-packets": 33368
    }
}

$ curl http://127.0.0.1:5000/netdev/enp0s3/address/
{
    "ipv4": [
        {
            "addr": "10.0.2.15",
            "broadcast": "10.0.2.255",
            "netmask": "255.255.255.0"
        }
    ],
    "ipv6": [
        {
            "addr": "fe80::a00:27ff:feaf:dc60%enp0s3",
            "netmask": "ffff:ffff:ffff:ffff::"
        }
    ],
    "mac": [
        {
            "addr": "08:00:27:af:dc:60",
            "broadcast": "ff:ff:ff:ff:ff:ff"
        }
    ]
}

ebiken@u1510d:~/nnapi/test$ curl http://127.0.0.1:5000/netdev/enp0s3/address/ipv4/
[
    {
        "addr": "10.0.2.15",
        "broadcast": "10.0.2.255",
        "netmask": "255.255.255.0"
    }
]

$ curl http://127.0.0.1:5000/netdev/enp0s3/stats/
{
    "rx-bytes": 11340796,
    "rx-compressed": 0,
    "rx-drop": 0,
    "rx-errs": 0,
    "rx-fifo": 0,
    "rx-frame": 0,
    "rx-multicast": 0,
    "rx-packets": 53635,
    "tx-bytes": 8542150,
    "tx-carrier": 0,
    "tx-colls": 0,
    "tx-compressed": 0,
    "tx-drop": 0,
    "tx-errs": 0,
    "tx-fifo": 0,
    "tx-packets": 33501
}
````

