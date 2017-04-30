# qframe-filter-inventory
Qframe filter which keeps an inventory and can be queried by other plugins.


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

### Fire Up `main.go`

```bash
# 
```
