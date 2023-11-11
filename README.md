On mikrotik:
```
/tool sniffer
set file-limit=10000KiB filter-interface=WAN filter-ip-address=47.91.67.66/32 filter-port=5279 filter-stream=yes streaming-enabled=yes streaming-server=192.168.88.2
```

On destination server (192.168.88.2):
```
go run sniffer -verbose -port 9900
```
