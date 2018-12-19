
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
  "Db":{
      "DbDriverName":"mysql",
      "DbDriverSource":"root:87654321@tcp(127.0.0.1:3306)/chain" //mysql datasource
  },
  "Ela":{
      "Host":"127.0.0.1:21334",                                  //<nodeIp>:<restful-port>
      "ServerPort":"8080"                                        //misc restful server port
  },
  // if you don't have btc node remove this part of configuration
  "Btc":{
      "Host":"127.0.0.1:18332",                                  //<nodeIp>:<jsonrpc-port>
      "Rpcuser":"clark",                                         //rpc user
      "Rpcpasswd":"DrGhlxKuqP02m47TnDRAeZqir6Gt5V0secCHajUKW-0=" //rpc password
  }
}
```