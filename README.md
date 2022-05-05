
## db setup
``` bash
# nsqlookupd 起動
$ nsqlookupd
# nsqd 起動
$ nsqd --lookupd-tcp-address=localhost:xxxx
# mongo db起動
$ mongod --dbpath ./db
```