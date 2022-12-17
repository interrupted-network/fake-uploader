# Fake Uploader
Fake uploader to balance network usage

## Download

GitHub [Releases](https://github.com/interrupted-network/fake-uploader/releases)

## Requirements

### Install vnstat

```sh
apt install vnstat
```

### Configurations
example:
```yaml
estimator:
  interfaceName: wlo1
uploader:
  targets:
    - network: tcp
      address: 1.2.3.4:443
coordinator:
  uploadSize:
    min: 1024000
    max: 102400000
  interval: 1s
  concurrent: 10
  txRxMinRatio: 10
  txRxMaxRatio: 15
```
- estimator is for section of estimating network. `interfaceName` would be the name
of your current network interface.  
to get your current network interface name, you can run one of commands bellow:
    - route -n
    - ip link show
    - netstat -i  
    in output you'll see a table as result, Under column `Iface` you'll see the interface name. [More information](https://linuxhint.com/list-network-interfaces-ubuntu/)
- uploader: you can manage uploader in this section.
    - targets: destination and protocols are being configured here.  
    you can search and find many addresses by searching in google! the more addresses you place in this section, you'll have more stable service, and less distrupting others(because of sending fake data to target receivers, i hope they don't mind for that little spam...) 
- coordinator is coordinator between estimator and uploader
    - uploadSize: range of upload size for each connection.  
    the data would be a random size in range of given min & max containing zero bytes.
    - interval: is check time of estimation cycle
    - concurrent: number of parallel requests would be configured by this key.  
    each time that ratio of download/upload is less than the configured data, 
    the service will put {concurrent} requests in queue in the same time.
    - txRxMinRatio: the minimum ratio of `download/upload` total size.  
    application uploader would trigger in any ratio less than this configured value.
    - txRxMaxRatio: the maximum ratio of `download/upload` total size.  
    application uploader would stop in any ratio bigger than this configured value.

---

Seeking Freedom :pray:
