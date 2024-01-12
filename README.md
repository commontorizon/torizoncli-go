`torizoncli` is a command-line interface for [Torizon Cloud](https://app.torizon.io). It is implemented using the [Torizon Cloud API](https://developer.toradex.com/torizon/torizon-platform/torizon-api/) (currently in beta).
### Design Principles
- Feature parity with the Torizon Cloud API: everything that you can do with the API should be wrapped in a convenient (but perhaps opinionated) in `torizoncli`.
- Composability: all commands that can should also output plain json to stdout, so users can easily pipe it into `jq` or use it in CI/CD automation scenarios.
- Easily discoverable CLI: commands should be clear and do one thing one. We try to be familiar with tools Torizon Cloud users may be accustomed to, such as `dockercli`. Sub-subcommands are used to set context instead of flags and shortcuts are always provided.
- Stateless: a command should never depend on a previous command. 
- Configuration files may be used, but as second-class citizens; environment variables will always be preferred and not used as memory.
### Using it
1. [Follow this documentation](https://developer.toradex.com/torizon/torizon-platform/torizon-api/#how-to-use-torizon-cloud-api) to generate API credentials (client id and client secret) using [Torizon Cloud](https://app.torizon.io) web ui.
2. Export the API credentials as environment variables:
```
$ export TORIZON_API_CLIENT_ID="<copy id from web ui>" 
$ export TORIZON_API_CLIENT_SECRET="<copy secret from web ui>"                
```
### Examples
```
$ torizoncli device list
DeviceUuid                              DeviceName                    DeviceId                      LastSeen                      CreatedAt                     ActivatedAt         DeviceStatus        Notes          
31f7a721-4184-459f-94a2-9919f52d9fc2    AM62                          verdin-am62-15133479-31f7a7   2023-12-18 22:24:26           2023-12-09 19:24:13           2023-12-09 19:24:32 NotSeen             <nil>          
ab5ca75a-cfe7-437a-b18f-c0e517f171a2    Healthy-Gelato                verdin-imx8mp-06817296-ab5ca7 2024-01-09 13:51:37           2024-01-09 13:32:44           2024-01-09 13:33:05 UpdatePending       <nil>          
f924d2ee-fb13-42f8-912b-da00ae42a019    Toasted-Flammkuchen           verdin-imx8mp-06817296-f924d2 2024-01-09 17:45:36           2024-01-09 17:24:30           2024-01-09 17:24:47 Error               <nil>          
6de8ec67-a316-4bfb-bb83-19dbb7f54c8b    Uniform-Whale                 qemux86-64-6de8ec             2024-01-11 13:52:53           2024-01-11 13:37:32           2024-01-11 13:37:41 UpToDate            <nil>          
3ae63ca4-7ba1-4d82-b58f-ea193822e3af    torizon-common-6.5.0          qemux86-64-3ae63c             2023-12-13 20:02:05           2023-12-13 18:53:14           2023-12-13 18:53:24 UpToDate            <nil>  
```

- Show network information for all provisioned devices:
```
$ torizoncli device network
DeviceUuid                              Hostname                      LocalIpV4                     MacAddress                    
31f7a721-4184-459f-94a2-9919f52d9fc2    verdin-am62-15133479          10.42.0.156                   00:14:2d:e6:eb:27             
ab5ca75a-cfe7-437a-b18f-c0e517f171a2    verdin-imx8mp-06817296        10.42.0.137                   00:14:2d:68:06:10             
f924d2ee-fb13-42f8-912b-da00ae42a019    verdin-imx8mp-06817296        10.42.0.137                   00:14:2d:68:06:10             
6de8ec67-a316-4bfb-bb83-19dbb7f54c8b    qemux86-64                    10.0.2.15                     52:54:00:12:34:56             
3ae63ca4-7ba1-4d82-b58f-ea193822e3af    qemux86-64                    10.0.2.15                     52:54:00:12:34:56             
```
- Request information in json and pipe it to jq for filtering
```
$ torizoncli device list --json | jq .                            
{
  "limit": 50,
  "offset": 0,
  "total": 5,
  "values": [
    {
      "activatedAt": "2023-12-09T19:24:32Z",
      "createdAt": "2023-12-09T19:24:13Z",
      "deviceId": "verdin-am62-15133479-31f7a7",
      "deviceName": "AM62",
      "deviceStatus": "NotSeen",
      "deviceUuid": "31f7a721-4184-459f-94a2-9919f52d9fc2",
      "lastSeen": "2023-12-18T22:24:26Z"
    },
    {
      "activatedAt": "2024-01-09T13:33:05Z",
      "createdAt": "2024-01-09T13:32:44Z",
      "deviceId": "verdin-imx8mp-06817296-ab5ca7",
      "deviceName": "Healthy-Gelato",
      "deviceStatus": "UpdatePending",
      "deviceUuid": "ab5ca75a-cfe7-437a-b18f-c0e517f171a2",
      "lastSeen": "2024-01-09T13:51:37Z"
    },
    {
      "activatedAt": "2024-01-09T17:24:47Z",
      "createdAt": "2024-01-09T17:24:30Z",
      "deviceId": "verdin-imx8mp-06817296-f924d2",
      "deviceName": "Toasted-Flammkuchen",
      "deviceStatus": "Error",
      "deviceUuid": "f924d2ee-fb13-42f8-912b-da00ae42a019",
      "lastSeen": "2024-01-09T17:45:36Z"
    },
    {
      "activatedAt": "2024-01-11T13:37:41Z",
      "createdAt": "2024-01-11T13:37:32Z",
      "deviceId": "qemux86-64-6de8ec",
      "deviceName": "Uniform-Whale",
      "deviceStatus": "UpToDate",
      "deviceUuid": "6de8ec67-a316-4bfb-bb83-19dbb7f54c8b",
      "lastSeen": "2024-01-11T13:52:53Z"
    },
    {
      "activatedAt": "2023-12-13T18:53:24Z",
      "createdAt": "2023-12-13T18:53:14Z",
      "deviceId": "qemux86-64-3ae63c",
      "deviceName": "torizon-common-6.5.0",
      "deviceStatus": "UpToDate",
      "deviceUuid": "3ae63ca4-7ba1-4d82-b58f-ea193822e3af",
      "lastSeen": "2023-12-13T20:02:05Z"
    }
  ]
}
```
### Building it locally
The Torizon Cloud API uses the [OpenAPI specification](https://swagger.io/specification/), meaning it's necessary to generate the API definitions. In this step, you have two options:
1. Use the pre-generated code available here: https://github.com/commontorizon/torizon-openapi-go (default option using this repo).
2. Generate the API code from scratch: download the `torizon-openapi.yaml` specification from https://app.torizon.io/api/docs/torizon-openapi.yaml.
	1. Follow the instructions in https://github.com/commontorizon/torizon-openapi-go/commit/667658705df107b61bcf912151a444a42bedb06a to build the definitions.
3. Build
```
go build -v ./..
```
### Contributing
Please open an MR or an issue!
Things we currently need:
- [ ] Better pretty-printing (handling different terminal sizes etc)
- [ ] Tests
- [ ] Wrapping more of the Torizon Cloud API

Also, grep for `FIXME` and you'll find something to contribute!
