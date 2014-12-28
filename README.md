nsq_to_logentries
=================

dump messages from NSQ to Logentries.

Usage
-----

```
nsq_to_logentries -topics="events,alerts" -token="logentries_token" -lookupd="http://localhost:4161"
```

* There's also a docker image - [gbenhaim/nsq_to_logentries](https://registry.hub.docker.com/u/gbenhaim/nsq_to_logentries)
