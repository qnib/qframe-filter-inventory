# qframe-filter-inventory
Qframe filter which keeps an inventory and can be queried by other plugins.

## Prebuild Docker Image

```bash
$ docker run -ti --name qframe-filter-inventory --rm \
             -v /var/run/docker.sock:/var/run/docker.sock qnib/qframe-filter-inventory
> execute CMD 'qframe-filter-inventory'
2017/04/30 19:40:49 [II] Dispatch broadcast for Back, Data and Tick
2017/04/30 19:40:49 [  INFO] inventory >> Start inventory v0.1.0
2017/04/30 19:40:49 [  INFO] docker-events >> Connected to 'moby' / v'17.05.0-ce-rc1'
2017/04/30 19:40:49 [ DEBUG] docker-events >> Already running container /qframe-filter-inventory: SetItem(183fa77bca880e68f345c92e9e913d55d1ca0587ffc78879ca051b91bf764126)
2017/04/30 19:40:50 [ DEBUG] inventory >> SearcRequest for name TestCnt11493581249
2017/04/30 19:40:50 [ DEBUG] inventory >> SearcRequest for name TestCnt21493581249
2017/04/30 19:40:50 [ DEBUG] docker-events >> Just started container /TestCnt11493581249: SetItem(f8008cbf26ed0f47043f0d42cc43482120491012ba5c1f17e1d3ff8de037737b)
2017/04/30 19:40:50 [ DEBUG] inventory >> #### Received message on Data-channel: TestCnt11493581249: container.start
2017/04/30 19:40:50 [  INFO] inventory >> Received Event: container.start
2017/04/30 19:40:51 [ DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
2017/04/30 19:40:51 [  INFO] inventory >>  SUCCESS > Request: /TestCnt11493581249 (length of PendingPendingRequests: 1)
2017/04/30 19:40:53 [ DEBUG] docker-events >> Just started container /TestCnt21493581249: SetItem(075d6c40c13c1e3a28d45e06d4695011432d3997acbceb0b66e039bba1af7544)
2017/04/30 19:40:53 [ DEBUG] inventory >> #### Received message on Data-channel: TestCnt21493581249: container.start
2017/04/30 19:40:53 [  INFO] inventory >> Received Event: container.start
2017/04/30 19:40:54 [ DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
2017/04/30 19:40:54 [  INFO] inventory >>  SUCCESS > Request: /TestCnt21493581249 (length of PendingPendingRequests: 0)
2017/04/30 19:40:54 [ DEBUG] inventory >> PendingRequests has length: 0
```

## Development

### Prepare go-libs
```bash
$ docker run -ti --name qframe-filter-inventory --rm -e SKIP_ENTRYPOINTS=1 \
           -v ${GOPATH}/src/github.com/qnib/qframe-filter-inventory:/usr/local/src/github.com/qnib/qframe-filter-inventory \
           -v ${GOPATH}/src/github.com/qnib/qframe-collector-docker-events/lib:/usr/local/src/github.com/qnib/qframe-collector-docker-events/lib \
           -v ${GOPATH}/src/github.com/qnib/qframe-types:/usr/local/src/github.com/qnib/qframe-types \
           -v ${GOPATH}/src/github.com/qnib/qframe-utils:/usr/local/src/github.com/qnib/qframe-utils \
           -v ${GOPATH}/src/github.com/qnib/qframe-inventory/lib:/usr/local/src/github.com/qnib/qframe-inventory/lib \
           -v /var/run/docker.sock:/var/run/docker.sock \
           -w /usr/local/src/github.com/qnib/qframe-filter-inventory \
            qnib/uplain-golang bash
# govendor update github.com/qnib/qframe-types github.com/qnib/qframe-utils \
                  github.com/qnib/qframe-filter-inventory/lib \
                  github.com/qnib/qframe-inventory/lib \
                  github.com/qnib/qframe-collector-docker-events/lib
# govendor fetch +m 
```


