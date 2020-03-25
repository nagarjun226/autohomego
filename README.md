# configmgr
Home automation project - the config manager

## Features

- `config.json` will contain the deafault config stored on disk
- Service will read `config.json` every 30 seconds
- Change the device config by changing the json
- Obtain the service config by `GET` call to the service manager

## Future

- update config via http request

## Order of Precedence

- `base` is the default config for all services
- Service specific config having any of the same parameters as the base config will take precedence
- In the future config change posted by put request will take precendence over the service specific config

## Developer notes

- config.json will not be checked in

## Schema

- Every `config.json` has a base config which is a common config across all the services
- Every service will have its own specific config. These will take precedence incase there is a clash
```
    base : 
            attribute: value;
            .
            .
    service1 : 
            attribute: value;
            .
            .
    service2: 
            attribute: value;
            .
            .
    .
    .
    .
```

## API

**Definition**

`GET /getconfig/{service_name <string>}`

**Response**
- `400 Error` On Bad Request
- `500 Error` On Internal server error
- `200 OK` On success
```json
service1 : 
        attribute: value;
        .
        .
```

## ToDOs

- Test Code
- logging
- better error handling



