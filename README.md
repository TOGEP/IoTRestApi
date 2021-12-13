# IoTRestApi
HTTP API like MQTT

### usage
```
$ curl -X POST -H 'content-type: application/json' --data '{"topic":"test", "data":25}' http://localhost:8080/temperature/`

$ curl -X GET http://localhost:8080/temperature/test
{"topic":"test","data":25}%
```
