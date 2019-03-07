
## build & run

Install golang version 1.10 or above , install glide to manage package dependency
```bash
>$ glide update & glide install 

>$ go build -o misc

>$ ./misc

```

change config.json
```
{
  "ServerPort":"8080",                                              //misc restful server port  
  "Db":{
      "DbDriverName"  :"mysql",
      "DbDriverSource":"root:87654321@tcp(127.0.0.1:3306)/chain"    //mysql datasource
  },
  "Ela":{
      "Host":"127.0.0.1:21334",                                     //<nodeIp>:<restful-port>
  },
  "Cmc":{
      "ApiKey":["d72b4bf0-00f3-4bab-8dca-45df2b0fb14c","cfa976fc-7642-4946-a0b2-858619fb2802"],
      "Inteval":"30m",
      "NumOfCoin":2000
  },
  // if you don't have btc node remove this part of configuration
  "Btc":{
      "Host"        :"127.0.0.1:18332",                             //<nodeIp>:<jsonrpc-port>
      "Rpcuser"     :"clark",                                       //rpc user
      "Rpcpasswd"   :"DrGhlxKuqP02m47TnDRAeZqir6Gt5V0secCHajUKW-0=" //rpc password
      "MinConfirm"  :1                                              //minimum required confirmation times
      "Net"         :"testnet3"                                     //mainnet/regtest/testnet3
  },
  "VisitKey":"729E2BB0AEEC048FF9DC7996D394889687BF76AFA832F07E011AA5A3BE272310"
}
```