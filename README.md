# DictionaryApi

> A simple REST API built for key-value pair get/set operations

* No third-party packages/dependencies

## Requirements

To be able to show the desired features of curl this REST API must match a few
requirements:

* [x] `GET /api/get?Key={key}` returns KeyValue as JSON
* [x] `POST /api/set?Key={key}&Value={value}` saves key-value pair and returns http 200 if successful
* [x] `DELETE /api/flush` deletes all data

### Data Types

A KeyValue object should look like this:
```json
{"key": "somekey","value": "somevalue"}
```

### Persistence

There is no persistence.