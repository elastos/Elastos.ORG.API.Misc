Set up client
########################################

.. toctree::
  :maxdepth: 3

.. walk-through:

Walk Through
=================================
The code is written in go . make sure you install go , and version has to be 1.10 or above.

Code
*****
download the code::

    git clone https://github.com/elastos/Elastos.ORG.API.Misc.git

Config
*****
modify config.json::

    {
        // Misc Server Port Config
        "ServerPort":"8080",
        // Misc Database Mysql Db Config
        "Db":{
            "DbDriverName":"mysql",
            "DbDriverSource":"root:87654321@tcp(127.0.0.1:3306)/chain"
        },
        //Local Node config , can be main chain or side chain
        "Ela":{
            "Host":"localhost:21334"
        },
        //CoinMarketCap API Key config
        "Cmc":{
          "ApiKey":["d72b4bf0-00f3-4bab-8dca-45df2b0fb14c","cfa976fc-7642-4946-a0b2-858619fb2802"],
          "Inteval":"30m",
          "NumOfCoin":200
        },
        //Btc Node Config , Can be omit if you don't need it .
        "Btc":{
          "Host":"127.0.0.1:18332",
          "Rpcuser":"clark",
          "Rpcpasswd":"DrGhlxKuqP02m47TnDRAeZqir6Gt5V0secCHajUKW-0=",
          "MinConfirm":1,
          "Net":"testnet3"
        },
        //Visit Key of Cmc Api.
        "VisitKey":"729E2BB0AEEC048FF9DC7996D394889687BF76AFA832F07E011AA5A3BE272310"
    }

Install
*****
install project::

    glide update & glide install
    go build -o misc

Deploy
*****
start your project::

    ./misc

