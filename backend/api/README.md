# code generation

Open API spec code generated with oap-codegen project

## Generate for client
```shell script
oapi-codegen -package client -o client.get.go -generate "types,server,spec" cerberus-client.yaml
```

All *.gen.go files should not be modified
