## Parsing tzsp streams of captured growatt data
On mikrotik:
```
/tool sniffer
set file-limit=10000KiB filter-interface=WAN filter-ip-address=47.91.67.66/32 filter-port=5279 filter-stream=yes streaming-enabled=yes streaming-server=192.168.88.2:9900
```

On destination server (192.168.88.2):
```
go run tzsp-parser.go -verbose -port 9900
```

## parsing raw growatt data (just the tcp packet payload)
```
go run file-parser.go <filename>
```
