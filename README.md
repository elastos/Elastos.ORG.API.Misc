
## build & run

Install golang version 1.10 or above , install glide to manage package dependency
```bash
>$ glide update & glide install 

>$ go build -o elaChain main.go

>$ ./elaChain

```

change config.json
```
{
  "Node":"http://<nodeIp>:<restful-port>",
  "ServerPort":"misc restful server port",
  "DbDriverName":"mysql",
  "DbDriverSource":"root:12345678@tcp(127.0.0.1:3306)/chain"
}
```