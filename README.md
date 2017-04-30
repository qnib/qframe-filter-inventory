# qframe-filter-inventory
Qframe filter which keeps an inventory and can be queried by other plugins.

## Prebuild Docker Image

The inventory filter listens for messages from the [`docker-events` plugin](https://github.com/qnib/qframe-collector-docker-events) and holds an inventory of all containers (of the engine reachable by `/var/run/docker.sock`)
which can be queried by other plugins using `qframe_inventory.InventoryRequest` messages on the `QChan.Data` channel.

An example run can be seen here:

```bash
$ docker run -ti --name qframe-filter-inventory --rm \                                                                                                                                                                                      git:(master↑1|)
             -v /var/run/docker.sock:/var/run/docker.sock qnib/qframe-filter-inventory
> execute CMD 'qframe-filter-inventory'
2017/04/30 20:16:37 [II] Dispatch broadcast for Back, Data and Tick
2017/04/30 20:16:37 [  INFO] inventory >> Start inventory v0.1.0
2017/04/30 20:16:37 [  INFO] docker-events >> Connected to 'moby' / v'17.05.0-ce-rc1'
2017/04/30 20:16:37 [ DEBUG] docker-events >> Already running container /qframe-filter-inventory: SetItem(2b6a3146217f0973bef7edb90ce6c1249acd8be7edc9874d5e184b8736d94e2d)
2017/04/30 20:16:38 [ DEBUG] inventory >> SearcRequest for name TestCnt11493583397
2017/04/30 20:16:38 [ DEBUG] inventory >> SearcRequest for name TestCnt21493583397
2017/04/30 20:16:39 [ DEBUG] docker-events >> Just started container /TestCnt11493583397: SetItem(bf3824dd1e85583e377b2bc90f45cb443eb3fde7a93bd127202360738b842264)
2017/04/30 20:16:39 [ DEBUG] inventory >> #### Received message on Data-channel: TestCnt11493583397: container.start
2017/04/30 20:16:39 [  INFO] inventory >> Received Event: container.start
2017/04/30 20:16:40 [ DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
2017/04/30 20:16:40 [  INFO] inventory >>  SUCCESS > Request: /TestCnt11493583397 (length of PendingPendingRequests: 1)
2017/04/30 20:16:42 [ DEBUG] docker-events >> Just started container /TestCnt21493583397: SetItem(0cd195d9035f241dfc789542d31ed061d405de543f804a1f26fae459b8e84a1d)
2017/04/30 20:16:42 [ DEBUG] inventory >> #### Received message on Data-channel: TestCnt21493583397: container.start
2017/04/30 20:16:42 [  INFO] inventory >> Received Event: container.start
2017/04/30 20:16:42 [ DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
2017/04/30 20:16:42 [  INFO] inventory >>  SUCCESS > Request: /TestCnt21493583397 (length of PendingPendingRequests: 0)
2017/04/30 20:16:42 [ DEBUG] inventory >> PendingRequests has length: 0
```

What is going on here:

- `2017/04/30 20:16:38.*inventory >> SearcRequest for name TestCnt.*` The `main.go` creates two search requests which are submitted to the inventory queue.
  As the containers are not yet started, they are added to the `Inventory.PendingRequests` array.
- When the first container is started at `2017/04/30 20:16:39`, it triggers a message to the `docker-events` collector, which subsequently ends up in the Inventory of `filter-inventory`
- The ticker generates a tick-event at `2017/04/30 20:16:40`, which results in a run of [CheckRequests](https://github.com/qnib/qframe-inventory/blob/master/lib/inventory.go#L81) and results in a response of the requests Back channel.
- As at `2017/04/30 20:16:42` the second container is started the search can be fulfilled and results in the second `SUCCESS > Request: /TestCnt21493583397 (length of PendingPendingRequests: 0)`

After that the PendingRequest list is empty and the command will exit.

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


