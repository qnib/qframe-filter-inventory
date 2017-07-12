# qframe-filter-inventory
Qframe filter which keeps an inventory and can be queried by other plugins.

**Depreciated!** Moved to [qframe/cache-inventory](https://github.com/qframe/cache-inventory)

## Prebuild Docker Image

The inventory filter listens for messages from the [`docker-events` plugin](https://github.com/qnib/qframe-collector-docker-events) and holds an inventory of all containers (of the engine reachable by `/var/run/docker.sock`)
which can be queried by other plugins using `qframe_inventory.InventoryRequest` messages on the `QChan.Data` channel.

An example run can be seen here:

```bash
$ docker run -ti --name qframe-filter-inventory --rm \                                                                                                                                                                                      git:(master↑1|)
             -v /var/run/docker.sock:/var/run/docker.sock qnib/qframe-filter-inventory
> execute CMD 'qframe-filter-inventory'
 1	 [II] Dispatch broadcast for Back, Data and Tick
 2	 [ INFO] inventory >> Start inventory v0.1.0
 3	 [ INFO] docker-events >> Connected to 'moby' / v'17.05.0-ce-rc1'
 4	 [DEBUG] docker-events >> Already running container /qframe-filter-inventory: SetItem(01bfa4a2b369ad2dbdc2fa9f8e21906b8ad5af59e045e7787a0d18b9e6e37a52)
>5	 [DEBUG] inventory >> SearcRequest for name TestCnt11493595097
>6	 [DEBUG] inventory >> SearcRequest for name TestCnt21493595097
>7	 [DEBUG] inventory >> SearcRequest for name TestCnt31493595097 via QChan.Data.Send()
 8	 [ INFO] inventory >> Received InventoryRequest
>9	 [DEBUG] docker-events >> Just started container /TestCnt11493595097: SetItem(37e46a55892aab71869e7fa15fdcd03cba5230859c86c9275c99443fbcfa160b)
 10	 [DEBUG] inventory >> #### Received message on Data-channel: TestCnt11493595097: container.start
 11	 [ INFO] inventory >> Received Event: container.start
>12	 [DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
 13	 [ INFO] inventory >>  SUCCESS > Request: /TestCnt11493595097 (length of PendingRequests: 2)
 14	 [DEBUG] docker-events >> Just started container /TestCnt21493595097: SetItem(e9f357c856397744ee471d8ee7918e52a17013374766526c36a64062f5162e78)
 15	 [DEBUG] inventory >> #### Received message on Data-channel: TestCnt21493595097: container.start
 16	 [ INFO] inventory >> Received Event: container.start
 17	 [DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
>18	 [ INFO] inventory >>  SUCCESS > Request: /TestCnt21493595097 (length of PendingRequests: 1)
>19	 [DEBUG] docker-events >> Just started container /TestCnt31493595097: SetItem(710ad677c338a767223efe0dc65be860261f9fed483e716f2623bf40ec986e7e)
 20	 [DEBUG] inventory >> #### Received message on Data-channel: TestCnt31493595097: container.start
 21	 [ INFO] inventory >> Received Event: container.start
 22	 [DEBUG] docker-events >> Container was found in the inventory...
 23	 [DEBUG] inventory >> #### Received message on Data-channel: TestCnt11493595097: container.die
 24	 [ INFO] inventory >> Received Event: container.die
 25	 [DEBUG] inventory >> Ticker came along: p.Inventory.CheckRequests()
>26	 [ INFO] inventory >>  SUCCESS > Request: /TestCnt31493595097 (length of PendingRequests: 0)
 27	 [DEBUG] inventory >> PendingRequests has length: 0
```

What is going on here:

- In line 5-7 `SearcRequest for name TestCnt.*` the `main.go` creates two search requests which are submitted directly to to the inventory queue, the third is send to the Data-channel.
  As the containers are not yet started, they are added to the `Inventory.PendingRequests` array.
- When the first container is started in line 9, it triggers a message to the `docker-events` collector, which subsequently ends up in the Inventory of `filter-inventory`
- The ticker generates a tick-event in line 12, which results in a run of [CheckRequests](https://github.com/qnib/qframe-inventory/blob/master/lib/inventory.go#L81) and results in a response of the requests Back channel.
- In line 14 and 19 the last two containers are started, the search can be fulfilled and results in a count-down to an empty Request-queue `SUCCESS > Request: /TestCnt21493583397 (length of PendingRequests: 0)`

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


