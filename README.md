# autohomego
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
