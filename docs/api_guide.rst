Getting started with the MISC API
########################################

.. toctree::
  :maxdepth: 3

Introduction
=============
MISC has a Restful API with URL endpoints corresponding to actions that users can perform with their channels. The endpoints accept and return JSON encoded objects. The API URL path always contains the API version in order to differentiate queries to different API versions. All queries start with: ``/api/<version>/`` where ``<version>`` is an integer representing the current API version.

.. api:

Elastos
=================================
using the following api ,we can get dpos vote statistic.

Get dpos producer vote statistics
------------------------------------------------
producer's vote statistics of specific height

.. http:get:: /api/1/dpos/producer/(string:`producer_public_key`)/(int:`height`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/dpos/producer/03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800/9999999 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":[
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"2638f858000dd118015daa7b1ee23c86e1c0738b5e641265d52f6612c527c672",
                    "N":0,
                    "Value":"4999",
                    "Outputlock":0,
                    "Address":"EbeD11dua88L9VQtNmJuEez8aVYX294CML",
                    "Block_time":1551800055,
                    "Height":233745
                },
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"82fce02fb0e835102eb37633e513e78c825a534d46146962391866e25bf8005c",
                    "N":0,
                    "Value":"9999",
                    "Outputlock":0,
                    "Address":"EKmp4dqTSMVW2f2H3x5H2A6vQf7FJV8Frj",
                    "Block_time":1551838308,
                    "Height":234056
                },
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"74f2beb77f15fcc6f36e43533aec254fc17b84edbb7e2b3a625c9ac2867a7435",
                    "N":0,
                    "Value":"123",
                    "Outputlock":0,
                    "Address":"EWHEoukFBK6AyMjuS9ucxhQ2twS7BKQEv8",
                    "Block_time":1551838618,
                    "Height":234058
                },
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"1a71b89c5e6c1b9baf31884f075f5e3ea159d8edfe5d665a2f5182d0c715ff91",
                    "N":0,
                    "Value":"9999",
                    "Outputlock":0,
                    "Address":"EYZt2Xk76NNFEHiihqkyBhyzuw1abcheXF",
                    "Block_time":1551850832,
                    "Height":234161
                },
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"71083736e824c73e4b327a8b958dbbd00aec879768a96963cbdfc5008e1bd393",
                    "N":0,
                    "Value":"0.01111111",
                    "Outputlock":0,
                    "Address":"ELbKQrj8DTYn2gU7KBejcNWb4ix4EAGDmy",
                    "Block_time":1551851053,
                    "Height":234163
                },
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"fbc81da6db6db5cb09c76fe405cf238353a8e837dda5acacd137ba43a9da1d02",
                    "N":0,
                    "Value":"9999",
                    "Outputlock":0,
                    "Address":"ENaaqePNBtrZsNbs9uc35CPqTbvn8oaYL9",
                    "Block_time":1551853616,
                    "Height":234180
                },
                {
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"82529a764fd1bbdd4ae39e9bb791d029ecb3010b7db48a7b5d1edfe8be71f36e",
                    "N":0,
                    "Value":"9999",
                    "Outputlock":0,
                    "Address":"Ea3XHVqFiAjYA4sSCTQSmrWQafGkbxaYxe",
                    "Block_time":1551853616,
                    "Height":234180
                }
            ],
            "status":200
        }

   :statuscode 200:   no error
   :statuscode 400:   bad request
   :statuscode 404:   not found request
   :statuscode 500:   internal error
   :statuscode 10001: process error

Get dpos voter's statistics
------------------------------------------------
voter's statistics

.. http:get:: /api/1/dpos/address/(string:`address`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/dpos/address/ENaaqePNBtrZsNbs9uc35CPqTbvn8oaYL9 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":[
                {
                    "Vote_Header":{
                        "Value":"192.99891960",
                        "Node_num":3,
                        "Txid":"9e840a28faedf6a3d1500bbb2a872fe2f7459d5bc831cdcda2e949437f4a33c5",
                        "Height":268392,
                        "Nodes":[
                            "0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                            "033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800"
                        ],
                        "Block_time":1555847840,
                        "Is_valid":"YES"
                    },
                    "Vote_Body":[
                        {
                            "Producer_public_key":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                            "Value":"310196.0425229799",
                            "Address":"EdhP91WcY2WhyV8N6dCnBxbjAnGd2izrzY",
                            "Rank":3,
                            "Ownerpublickey":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                            "Nodepublickey":"ff",
                            "Nickname":"今天真好",
                            "Url":"www.helloword.com",
                            "Location":44,
                            "Active":false,
                            "Votes":"309844",
                            "Netaddress":"1.2.3.4",
                            "State":"Activate",
                            "Registerheight":234800,
                            "Cancelheight":0,
                            "Inactiveheight":0,
                            "Illegalheight":0,
                            "Index":2,
                            "Reward":"0",
                            "EstRewardPerYear":"46718.30201048"
                        },
                        {
                            "Producer_public_key":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "Value":"311559.3568213799",
                            "Address":"Eb8UHkQ2bJ4Ljux4yBePFdxB5Yp77VYHyt",
                            "Rank":2,
                            "Ownerpublickey":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "Nodepublickey":"03c18abb98f6679064bd44121f3b0a3f25dea1a8b8cb0e1b51dc9c26729f07ddc9",
                            "Nickname":"我怎么这么好看",
                            "Url":"www.douniwan.com",
                            "Location":263,
                            "Active":false,
                            "Votes":"311315.30210000",
                            "Netaddress":"8.8.8.8",
                            "State":"Activate",
                            "Registerheight":232288,
                            "Cancelheight":0,
                            "Inactiveheight":0,
                            "Illegalheight":0,
                            "Index":1,
                            "Reward":"0",
                            "EstRewardPerYear":"46909.52589293"
                        },
                        {
                            "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                            "Value":"357371.9983466",
                            "Address":"EX4eQnSSBG2CuhkSvaJHxrwtxS12Lxwy3M",
                            "Rank":1,
                            "Ownerpublickey":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                            "Nodepublickey":"16fffcff2affd4c7fffdfcffecfffff4ff",
                            "Nickname":"河北节点",
                            "Url":"www.elastos.org",
                            "Location":86,
                            "Active":false,
                            "Votes":"357029.00210000",
                            "Netaddress":"5JdHqndX1NyyTJnnRnAAKNsoJ9qBwcMYtvRduxHyGGdhzHwxPZo",
                            "State":"Activate",
                            "Registerheight":233734,
                            "Cancelheight":0,
                            "Inactiveheight":0,
                            "Illegalheight":0,
                            "Index":0,
                            "Reward":"0",
                            "EstRewardPerYear":"53335.38909126"
                        }
                    ]
                },
                {
                    "Vote_Header":{
                        "Value":"199.99935700",
                        "Node_num":1,
                        "Txid":"5a0d7958ff9677eef0fa7194db788add8722cf91fdaedc28c12acb677a58f8b3",
                        "Height":266138,
                        "Nodes":[
                            "033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd"
                        ],
                        "Block_time":1555574076,
                        "Is_valid":"NO"
                    },
                    "Vote_Body":[
                        {
                            "Producer_public_key":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "Value":"313289.9935201299",
                            "Address":"Eb8UHkQ2bJ4Ljux4yBePFdxB5Yp77VYHyt",
                            "Rank":2,
                            "Ownerpublickey":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "Nodepublickey":"03c18abb98f6679064bd44121f3b0a3f25dea1a8b8cb0e1b51dc9c26729f07ddc9",
                            "Nickname":"我怎么这么好看",
                            "Url":"www.douniwan.com",
                            "Location":263,
                            "Active":false,
                            "Votes":"311315.30210000",
                            "Netaddress":"8.8.8.8",
                            "State":"Activate",
                            "Registerheight":232288,
                            "Cancelheight":0,
                            "Inactiveheight":0,
                            "Illegalheight":0,
                            "Index":1,
                            "Reward":"0",
                            "EstRewardPerYear":"47013.01092436"
                        }
                    ]
                }
            ],
            "status":200
        }


.. http:get:: /api/1/dpos/address/(string:`address`)?pageSize=(int:`pageSize`)&pageNum=(int:`pageNum`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/dpos/address/ENaaqePNBtrZsNbs9uc35CPqTbvn8oaYL9?pageSize=1&pageNum=1 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":[
                {
                    "Vote_Header":{
                        "Value":"199.99935700",
                        "Node_num":1,
                        "Txid":"5a0d7958ff9677eef0fa7194db788add8722cf91fdaedc28c12acb677a58f8b3",
                        "Height":266138,
                        "Nodes":[
                            "033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd"
                        ],
                        "Block_time":1555574076,
                        "Is_valid":"NO"
                    },
                    "Vote_Body":[
                        {
                            "Producer_public_key":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "Value":"313289.9935201299",
                            "Address":"Eb8UHkQ2bJ4Ljux4yBePFdxB5Yp77VYHyt",
                            "Rank":2,
                            "Ownerpublickey":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                            "Nodepublickey":"03c18abb98f6679064bd44121f3b0a3f25dea1a8b8cb0e1b51dc9c26729f07ddc9",
                            "Nickname":"我怎么这么好看",
                            "Url":"www.douniwan.com",
                            "Location":263,
                            "Active":false,
                            "Votes":"311315.30210000",
                            "Netaddress":"8.8.8.8",
                            "State":"Activate",
                            "Registerheight":232288,
                            "Cancelheight":0,
                            "Inactiveheight":0,
                            "Illegalheight":0,
                            "Index":1,
                            "Reward":"0",
                            "EstRewardPerYear":"47013.01092436"
                        }
                    ]
                }
            ],
            "status":200
        }





   :statuscode 200:   no error
   :statuscode 400:   bad request
   :statuscode 404:   not found request
   :statuscode 500:   internal error
   :statuscode 10001: process error

Get producers of specific transactions
-----------------------------------------

.. http:post:: /api/1/dpos/transaction/producer

   **Example request**:

   .. sourcecode:: http

    POST /api/1/dpos/transaction/producer HTTP/1.1
    Host: localhost

      {
          "txid":[
            "59b6b468f75856b7980525ad7a1278e4998959211f57d81755e4248982fd18b8"
          ]
      }

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
        "result":[
            {
                "Producer":[
                    {
                        "Ownerpublickey":"02b28266ff709f4764374c0452e379671e47d66713efb4cce7812b3c9f4a12b2bc",
                        "Nodepublickey":"02b28266ff709f4764374c0452e379671e47d66713efb4cce7812b3c9f4a12b2bc",
                        "Nickname":"DHG(大黄哥)",
                        "Url":"www.eladhg.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"263036.79130980",
                        "Netaddress":"",
                        "State":"Activate",
                        "Registerheight":361360,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":6
                    },
                    {
                        "Ownerpublickey":"025220c50d7ba72c8f5a78972b4d157339d5a02d3ed8639f01dbae6c14de5585cb",
                        "Nodepublickey":"02c29d33e3caf772f153c5d866ee799d5d4ad38d5efe402d3d5fa980ae5fb5f9a1",
                        "Nickname":"greengang",
                        "Url":"www.ptcent.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"239143.67333523",
                        "Netaddress":"",
                        "State":"Activate",
                        "Registerheight":360878,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":14
                    },
                    {
                        "Ownerpublickey":"02f2101d918e95b9df92e58322f7b7d70a134dd0bf441c25758fe8a9a64e712ebd",
                        "Nodepublickey":"02f2101d918e95b9df92e58322f7b7d70a134dd0bf441c25758fe8a9a64e712ebd",
                        "Nickname":"ZDJ",
                        "Url":"www.zhidianjia.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"103658.61704950",
                        "Netaddress":"",
                        "State":"Activate",
                        "Registerheight":360618,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":42
                    },
                    {
                        "Ownerpublickey":"0279d982cda37fa7edc1906ec2f4b3d8da5af2c15723e14f368f3684bb4a1e0889",
                        "Nodepublickey":"0279d982cda37fa7edc1906ec2f4b3d8da5af2c15723e14f368f3684bb4a1e0889",
                        "Nickname":"ELA.SYDNEY",
                        "Url":"www.ela.sydney",
                        "Location":61,
                        "Active":false,
                        "Votes":"46492.26739977",
                        "Netaddress":"",
                        "State":"Activate",
                        "Registerheight":372790,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":53
                    }
                ],
                "Txid":"59b6b468f75856b7980525ad7a1278e4998959211f57d81755e4248982fd18b8"
            }
        ],
        "status":200
    }



Get dpos super node rank list
------------------------------------------------
rank list of producer

    .. http:get:: /api/1/dpos/rank/height/(int:`height`)

       **Example request**:

       .. sourcecode:: http

          GET /api/1/dpos/rank/height/241762 HTTP/1.1
          Host: localhost

       **Example response**:

       .. sourcecode:: http

          HTTP/1.1 200 OK
          Content-Type: application/json

            {
                "result":[
                    {
                        "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                        "Value":"357051",
                        "Address":"EX4eQnSSBG2CuhkSvaJHxrwtxS12Lxwy3M",
                        "Rank":1,
                        "Ownerpublickey":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                        "Nodepublickey":"16fffcff2affd4c7fffdfcffecfffff4ff",
                        "Nickname":"河北节点",
                        "Url":"www.elastos.org",
                        "Location":86,
                        "Active":false,
                        "Votes":"357029",
                        "Netaddress":"5JdHqndX1NyyTJnnRnAAKNsoJ9qBwcMYtvRduxHyGGdhzHwxPZo",
                        "State":"Activate",
                        "Registerheight":233734,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":0,
                        "Reward":"",
                        "EstRewardPerYear":"66741.53520809"
                    },
                    {
                        "Producer_public_key":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                        "Value":"311337.3",
                        "Address":"Eb8UHkQ2bJ4Ljux4yBePFdxB5Yp77VYHyt",
                        "Rank":2,
                        "Ownerpublickey":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                        "Nodepublickey":"03c18abb98f6679064bd44121f3b0a3f25dea1a8b8cb0e1b51dc9c26729f07ddc9",
                        "Nickname":"我怎么这么好看",
                        "Url":"www.douniwan.com",
                        "Location":263,
                        "Active":false,
                        "Votes":"311315.30000000",
                        "Netaddress":"8.8.8.8",
                        "State":"Activate",
                        "Registerheight":232288,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":1,
                        "Reward":"",
                        "EstRewardPerYear":"58196.53038233"
                    },
                    {
                        "Producer_public_key":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                        "Value":"309866",
                        "Address":"EdhP91WcY2WhyV8N6dCnBxbjAnGd2izrzY",
                        "Rank":3,
                        "Ownerpublickey":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                        "Nodepublickey":"ff",
                        "Nickname":"今天真好",
                        "Url":"www.helloword.com",
                        "Location":44,
                        "Active":false,
                        "Votes":"309844",
                        "Netaddress":"1.2.3.4",
                        "State":"Activate",
                        "Registerheight":234800,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":2,
                        "Reward":"",
                        "EstRewardPerYear":"57921.50854861"
                    },
                    {
                        "Producer_public_key":"03c78467b91805c95ada2530513069bef1f1f1e7b756861381ab534efa6d94e40a",
                        "Value":"218140.55555",
                        "Address":"EdfJA92nN9X4T9cKqkvyrunVuBWfF1Mumm",
                        "Rank":4,
                        "Ownerpublickey":"03c78467b91805c95ada2530513069bef1f1f1e7b756861381ab534efa6d94e40a",
                        "Nodepublickey":"fffff3fffffffffffffffbff1affffffec",
                        "Nickname":"聪聪2",
                        "Url":"1.4.7.9",
                        "Location":672,
                        "Active":false,
                        "Votes":"218115.55555000",
                        "Netaddress":"1.12.3.4",
                        "State":"Activate",
                        "Registerheight":233035,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":3,
                        "Reward":"",
                        "EstRewardPerYear":"40775.78712439"
                    },
                    {
                        "Producer_public_key":"021d59a84d2243111e39e8c2af0a5089127d142d52b18c3e4bf744e0c6f8af44e0",
                        "Value":"147232",
                        "Address":"ESpTiKXgLcYkzxdD7MuCmL9y9fbWrnH591",
                        "Rank":5,
                        "Ownerpublickey":"021d59a84d2243111e39e8c2af0a5089127d142d52b18c3e4bf744e0c6f8af44e0",
                        "Nodepublickey":"ffff1230ffff",
                        "Nickname":"www.12306.cn",
                        "Url":"www.12306.cn",
                        "Location":244,
                        "Active":false,
                        "Votes":"147210",
                        "Netaddress":"www.12306.cn",
                        "State":"Activate",
                        "Registerheight":232899,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":4,
                        "Reward":"",
                        "EstRewardPerYear":"27521.24965833"
                    },
                    {
                        "Producer_public_key":"036417ab256114a32bcff38f3e10f0384cfa9238afa41a163017687b3ce1fa17f2",
                        "Value":"139881",
                        "Address":"ETKVMhhQCjttNAjrbqmkAAYuYshLdaDnjm",
                        "Rank":6,
                        "Ownerpublickey":"036417ab256114a32bcff38f3e10f0384cfa9238afa41a163017687b3ce1fa17f2",
                        "Nodepublickey":"03e5b45b44bb1e2406c55b7dd84b727fad608ba7b7c11a9c5ffbfee60e427bd1da",
                        "Nickname":"聪聪3",
                        "Url":"225.7.3",
                        "Location":672,
                        "Active":false,
                        "Votes":"139850",
                        "Netaddress":"1.1.1.8",
                        "State":"Activate",
                        "Registerheight":233537,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":5,
                        "Reward":"",
                        "EstRewardPerYear":"26147.16857380"
                    },
                    {
                        "Producer_public_key":"02e578a6f4295765ad3be4cdac9be15de5aedaf1ae76e86539bb54c397e467cd5e",
                        "Value":"125906",
                        "Address":"EHdSBUH3nxkcAk9evU4HrENzEm8MHirkkN",
                        "Rank":7,
                        "Ownerpublickey":"02e578a6f4295765ad3be4cdac9be15de5aedaf1ae76e86539bb54c397e467cd5e",
                        "Nodepublickey":"fffeffddfffffff2fffffffffbffffffff",
                        "Nickname":"亦来云",
                        "Url":"www.yilaiyun.com",
                        "Location":244,
                        "Active":false,
                        "Votes":"125884",
                        "Netaddress":"www.yilaiyun.com",
                        "State":"Activate",
                        "Registerheight":233680,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":6,
                        "Reward":"",
                        "EstRewardPerYear":"23534.90042574"
                    },
                    {
                        "Producer_public_key":"02ddd829f3495a2ce76d908c3e6e7d4505e12c4718c5af4b4cbff309cfd3aeab88",
                        "Value":"108968",
                        "Address":"EevRwpP5GYz5s8fuMboUnhsAQVVKbyJSph",
                        "Rank":8,
                        "Ownerpublickey":"02ddd829f3495a2ce76d908c3e6e7d4505e12c4718c5af4b4cbff309cfd3aeab88",
                        "Nodepublickey":"ffffffffffffffffffffffffffffffffffff",
                        "Nickname":"曲率区动",
                        "Url":"www.bightbc.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"108946",
                        "Netaddress":"EfSkh3e9uaVN5iMdU7oUPYPmyMxrMsrDut",
                        "State":"Activate",
                        "Registerheight":234283,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":7,
                        "Reward":"",
                        "EstRewardPerYear":"20368.77535297"
                    },
                    {
                        "Producer_public_key":"03c7b1f234d5d16472fcdd24d121e4cd224e1074f558a3eb1a6a146aa91dcf9c0d",
                        "Value":"108186",
                        "Address":"EQR8f9y2Sd5gFG3LWEeC57qXc2yEnDhgm2",
                        "Rank":9,
                        "Ownerpublickey":"03c7b1f234d5d16472fcdd24d121e4cd224e1074f558a3eb1a6a146aa91dcf9c0d",
                        "Nodepublickey":"350181ff",
                        "Nickname":"范冰冰",
                        "Url":"1.8.5.8",
                        "Location":86,
                        "Active":false,
                        "Votes":"108164",
                        "Netaddress":"HTTP//HUANGBINGBING.COM",
                        "State":"Activate",
                        "Registerheight":233676,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":8,
                        "Reward":"",
                        "EstRewardPerYear":"20222.60049131"
                    },
                    {
                        "Producer_public_key":"03b688e0124580de452c400e01c628a690527e8742b6fa4645026dbc70155d7c8b",
                        "Value":"107863",
                        "Address":"EQHz2jPpgW8trYD4ejYgfi4sE4JSTf7m9N",
                        "Rank":10,
                        "Ownerpublickey":"03b688e0124580de452c400e01c628a690527e8742b6fa4645026dbc70155d7c8b",
                        "Nodepublickey":"ffffffffffff",
                        "Nickname":"基延一族",
                        "Url":"1.4.7.9",
                        "Location":672,
                        "Active":false,
                        "Votes":"107841",
                        "Netaddress":"www.vogue.com",
                        "State":"Activate",
                        "Registerheight":233684,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":9,
                        "Reward":"",
                        "EstRewardPerYear":"20162.22391801"
                    },
                    {
                        "Producer_public_key":"03bc2c2b75009a3a551e98bf206730501ecdf46e71b0405840ff1d5750094bd4ff",
                        "Value":"105047",
                        "Address":"ENxPtTR7Jn1kxhdTXedF28s3iz6djYfRaS",
                        "Rank":11,
                        "Ownerpublickey":"03bc2c2b75009a3a551e98bf206730501ecdf46e71b0405840ff1d5750094bd4ff",
                        "Nodepublickey":"fffffffd29fffffffafff8fafffffdfffa",
                        "Nickname":"乐天居士",
                        "Url":"www.baidu.com",
                        "Location":376,
                        "Active":false,
                        "Votes":"105025",
                        "Netaddress":"尽快哦孩子",
                        "State":"Activate",
                        "Registerheight":232892,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":10,
                        "Reward":"",
                        "EstRewardPerYear":"19635.84487651"
                    },
                    {
                        "Producer_public_key":"0230d383546d154d67cfafc6091c0736c0b26a8c7c16e879ef8011d91df976f1fb",
                        "Value":"104256",
                        "Address":"EMyStHAvvy1VLsLyow8uMRW4kUYLeGXF17",
                        "Rank":12,
                        "Ownerpublickey":"0230d383546d154d67cfafc6091c0736c0b26a8c7c16e879ef8011d91df976f1fb",
                        "Nodepublickey":"fffffffffffefffffffffffffbfcffffff",
                        "Nickname":"烽火",
                        "Url":"www.ela.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"104234",
                        "Netaddress":"www.ela.com",
                        "State":"Activate",
                        "Registerheight":233612,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":11,
                        "Reward":"",
                        "EstRewardPerYear":"19487.98769547"
                    },
                    {
                        "Producer_public_key":"028fb1a85f6a30a516b9e3516d03267403a3af0c96d73b0284ca0c1165318531ff",
                        "Value":"104066",
                        "Address":"ESqyiCizgyNNLKdVQhhtxtR5v5eCnkk3Qh",
                        "Rank":13,
                        "Ownerpublickey":"028fb1a85f6a30a516b9e3516d03267403a3af0c96d73b0284ca0c1165318531ff",
                        "Nodepublickey":"ffff9262",
                        "Nickname":"链世界",
                        "Url":"www.7234.cn",
                        "Location":86,
                        "Active":false,
                        "Votes":"101045",
                        "Netaddress":"www.7234.cn",
                        "State":"Activate",
                        "Registerheight":235373,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":12,
                        "Reward":"",
                        "EstRewardPerYear":"19452.47206412"
                    },
                    {
                        "Producer_public_key":"02db921cfb4bf504c83038212aafe52cc1d0a07eb71a399a0d2162fe0cd4d47720",
                        "Value":"99051",
                        "Address":"ERbFZNj5bukyRQe5G4gdXnbDqVyxcTNeFT",
                        "Rank":14,
                        "Ownerpublickey":"02db921cfb4bf504c83038212aafe52cc1d0a07eb71a399a0d2162fe0cd4d47720",
                        "Nodepublickey":"1234567890ffdffffffffcffffffffffffff",
                        "Nickname":"ios_us01",
                        "Url":"www.ios_us01.com",
                        "Location":684,
                        "Active":false,
                        "Votes":"99029",
                        "Netaddress":"192.168.1.22:25339",
                        "State":"Activate",
                        "Registerheight":233672,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":13,
                        "Reward":"",
                        "EstRewardPerYear":"18515.04632082"
                    },
                    {
                        "Producer_public_key":"033fb33f39276b93d3474cf7999887bed16c3211ee7f904399eeead4c480d7d592",
                        "Value":"98859",
                        "Address":"EXQZMbKMcmVmwv25AYbrzWPhFRSfqKcfKM",
                        "Rank":15,
                        "Ownerpublickey":"033fb33f39276b93d3474cf7999887bed16c3211ee7f904399eeead4c480d7d592",
                        "Nodepublickey":"19fffffe9dfffafffffffffffbcaffffff",
                        "Nickname":"晓黎-评财经",
                        "Url":"www.pingcj.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"98837",
                        "Netaddress":"Ed846C7M9Ax8x1qaftjSR53RZmfSvp8CpN",
                        "State":"Activate",
                        "Registerheight":235077,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":14,
                        "Reward":"",
                        "EstRewardPerYear":"18479.15684072"
                    },
                    {
                        "Producer_public_key":"030e4b487daf8e14dbd7023e3f6f475d00145a1f1cc87be4b8d58a4291ab0a3b1a",
                        "Value":"25974",
                        "Address":"EVFSvWoxiyvGLka4V6Wt394LEoUu8mDhk4",
                        "Rank":16,
                        "Ownerpublickey":"030e4b487daf8e14dbd7023e3f6f475d00145a1f1cc87be4b8d58a4291ab0a3b1a",
                        "Nodepublickey":"0241db65a4da2cdcbb648a881ced2a5ed64646ecc3a2cc9a75cec2853de61dbed1",
                        "Nickname":"ELASuperNode",
                        "Url":"www.ELASuperNode.com",
                        "Location":86,
                        "Active":false,
                        "Votes":"25952",
                        "Netaddress":"54.64.220.165",
                        "State":"Activate",
                        "Registerheight":237877,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":15,
                        "Reward":"",
                        "EstRewardPerYear":"4855.17373007"
                    },
                    {
                        "Producer_public_key":"0210694f4ab518037bc2dcc3f5e1a1030e8a36821ab019c10f29d4a894b8034498",
                        "Value":"55",
                        "Address":"ESwKtu2aYSHHfdWUPdg4b3PtibfaEcJEvT",
                        "Rank":17,
                        "Ownerpublickey":"0210694f4ab518037bc2dcc3f5e1a1030e8a36821ab019c10f29d4a894b8034498",
                        "Nodepublickey":"024babfecea0300971a6f0ad13b27519faff0ef595faf9490dc1f5f4d6e6d7f3fb",
                        "Nickname":"adr_us01",
                        "Url":"www.adr_us01_9.com",
                        "Location":93,
                        "Active":false,
                        "Votes":"33",
                        "Netaddress":"node-regtest-509.eadd.co:26339",
                        "State":"Activate",
                        "Registerheight":238437,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":16,
                        "Reward":"",
                        "EstRewardPerYear":"10.28084065"
                    },
                    {
                        "Producer_public_key":"0210cd8407f70b26dbb77039cdce61a526168d04b83885844294038759f57c525c",
                        "Value":"20",
                        "Address":"EdUn345wvDWj3knsYsquEkZsqhRRXYSdnK",
                        "Rank":18,
                        "Ownerpublickey":"0210cd8407f70b26dbb77039cdce61a526168d04b83885844294038759f57c525c",
                        "Nodepublickey":"0210cd8407f70b26dbb77039cdce61a526168d04b83885844294038759f57c525c",
                        "Nickname":"ios_us05",
                        "Url":"www.ios_us05.com",
                        "Location":244,
                        "Active":false,
                        "Votes":"20",
                        "Netaddress":"172.31.40.70:25339",
                        "State":"Activate",
                        "Registerheight":244762,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":17,
                        "Reward":"",
                        "EstRewardPerYear":"3.73848751"
                    },
                    {
                        "Producer_public_key":"03325ce52add7a799a61a305973b3d84aa4f622358ab3eb9f010f1175e2dab6b13",
                        "Value":"20",
                        "Address":"Eb9mkpHC787UGqeqNvXs7j4Thh6fX6rF9D",
                        "Rank":19,
                        "Ownerpublickey":"03325ce52add7a799a61a305973b3d84aa4f622358ab3eb9f010f1175e2dab6b13",
                        "Nodepublickey":"03325ce52add7a799a61a305973b3d84aa4f622358ab3eb9f010f1175e2dab6b13",
                        "Nickname":"ios_us06",
                        "Url":"www.ios_us06.com",
                        "Location":54,
                        "Active":false,
                        "Votes":"20",
                        "Netaddress":"172.31.45.130:25339",
                        "State":"Activate",
                        "Registerheight":244768,
                        "Cancelheight":0,
                        "Inactiveheight":0,
                        "Illegalheight":0,
                        "Index":18,
                        "Reward":"",
                        "EstRewardPerYear":"3.73848751"
                    }
                ],
                "status":200
            }

rank list of producer according to node state

    .. http:get:: /api/1/dpos/rank/height/(int:`height`)?state=(string:`state`)

       **Example request**:

       .. sourcecode:: http

          GET /api/1/dpos/rank/height/241762?state=illegal HTTP/1.1
          Host: localhost

       **Example response**:

       .. sourcecode:: http

          HTTP/1.1 200 OK
          Content-Type: application/json

            {
              "result": [
                {
                  "Producer_public_key": "02f59c9e337d4102d49082fe641b2e8aac5b229583bfa27cfba87790ac0113518b",
                  "Value": "1239723.59908182",
                  "Address": "ESq8nLmYuYaNYqJh5HeDy4ywx1jCRfeKe1",
                  "Rank": 1,
                  "Ownerpublickey": "02f59c9e337d4102d49082fe641b2e8aac5b229583bfa27cfba87790ac0113518b",
                  "Nodepublickey": "03a31e03e7b0e38fdf002a3d8954f1df357b638c63040f84955ee28345ce57e886",
                  "Nickname": "The Houston Supernode",
                  "Url": "https://twitter.com/elastos_houston/",
                  "Location": 1001,
                  "Active": 0,
                  "Votes": "506625.45634569",
                  "Netaddress": "",
                  "State": "Illegal",
                  "Registerheight": 360805,
                  "Cancelheight": 0,
                  "Inactiveheight": 408476,
                  "Illegalheight": 409882,
                  "Index": 81,
                  "Reward": "213.93704808",
                  "EstRewardPerYear": "8043.77456689"
                }
              ],
              "status": 200
            }


Get dpos total vote of specific height
------------------------------------------------
total vote of specific height

.. http:get:: /api/1/dpos/vote/height/(int:`height`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/dpos/vote/height/241762 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
          "result":2468878.85555,
          "status":200
        }


Get did information
------------------------------------------------
check did key's value

.. http:get:: /api/1/did/(string:`did`)/(string:`key`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/did/iYnguKQcpeVyrpN6edamSkky1brvQvCWr6/clark HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":{
                "Did":"iYnguKQcpeVyrpN6edamSkky1brvQvCWr6",
                "Did_status":1,
                "Public_key":"0354E3AE040052CD61A38DF72E189EDCBBB8BA81599DED25019E286F2E013A3726",
                "Property_key":"clark",
                "Property_value":"hello,world",
                "Txid":"c39dd1463678146467ea1b43d8905e75ac34e727a4c36824410089a6682b43c8",
                "Block_time":1543902641,
                "Height":66480
            },
            "status":200
        }

Get address transaction history
------------------------------------------------
check history of Address

.. http:get:: /api/1/history/(string:`addr`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/history/EQoascGFzdQ1rLKfNEavJKTm3hMRhBaXBT HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":{
                "History":[
                    {
                        "Txid":"71eedccf5a45acc17036294b7a1e70c5756dd4c72032cc76d5089c522509096b",
                        "Type":"income",
                        "Value":24,
                        "CreateTime":1544101570,
                        "Height":175548,
                        "Fee":100,
                        "Inputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ],
                        "Outputs":[
                            "EQoascGFzdQ1rLKfNEavJKTm3hMRhBaXBT",
                            "EQoascGFzdQ1rLKfNEavJKTm3hMRhBaXBT",
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ]
                    }
                ],
                "TotalNum":1
            },
            "status":200
        }


.. http:get:: /api/1/history/(string:`addr`)?pageSize=(int:`pageSize`)&pageNum=(int:`pageNum`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/history/EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA?pageSize=4&pageNum=1 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":{
                "History":[
                    {
                        "Txid":"58e11ffffafad7456835dc3eff8b28ec0d0b2aa39855d79bfb3e5c52b4cc308f",
                        "Type":"income",
                        "Value":8990000,
                        "CreateTime":1538114443,
                        "Height":38666,
                        "Fee":10000,
                        "Inputs":[

                        ],
                        "Outputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ]
                    },
                    {
                        "Txid":"b04714ff7076e18efeb9963f99af2f10bd870d39a1774e0821b71640d6b14e52",
                        "Type":"spend",
                        "Value":100,
                        "CreateTime":1538115578,
                        "Height":38667,
                        "Fee":100,
                        "Inputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ],
                        "Outputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA",
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ]
                    },
                    {
                        "Txid":"49eb7f1335d89a9c43446613850d1b763a5f84cad0d1d039ab5726125d2eab47",
                        "Type":"spend",
                        "Value":100,
                        "CreateTime":1538203592,
                        "Height":39046,
                        "Fee":100,
                        "Inputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA",
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ],
                        "Outputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA",
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ]
                    },
                    {
                        "Txid":"d7d8b34944790cf2e1fa47851612ddb4eb7e3f5bd75a4af2227d3786b6f4c122",
                        "Type":"spend",
                        "Value":100,
                        "CreateTime":1539238227,
                        "Height":44755,
                        "Fee":100,
                        "Inputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA",
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ],
                        "Outputs":[
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA",
                            "EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA"
                        ]
                    }
                ],
                "TotalNum":55295
            },
            "status":200
        }

Only show the history count of specific address

.. http:get:: /api/1/history/(string:`addr`)?showCount=true

   **Example request**:

   .. sourcecode:: http

      GET /api/1/history/EQoascGFzdQ1rLKfNEavJKTm3hMRhBaXBT HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":{
                "History":[
                ],
                "TotalNum":1
            },
            "status":200
        }

get transaction history with your desired order

.. http:get:: /api/1/history/(string:`addr`)?order=desc

   **Example request**:

   .. sourcecode:: http

      GET /api/1/history/EM2wjL3jgNHDZtR1e266V269n5WH6sYbCf HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":{
                "History":[
                    {
                        "Address":"EM2wjL3jgNHDZtR1e266V269n5WH6sYbCf",
                        "Txid":"a962022bc4a295ab4683ab4079f03d1e5bdb9bfdf5dac9c4eea003d18af16fbd",
                        "Type":"spend",
                        "Value":50000000000,
                        "CreateTime":1561557063,
                        "Height":409201,
                        "Fee":10000,
                        "Inputs":[
                            "EM2wjL3jgNHDZtR1e266V269n5WH6sYbCf"
                        ],
                        "Outputs":[
                            "EUX2LMtHBV1Ni7nAXPhBdnudrUvddU2Ecv"
                        ],
                        "TxType":"TransferAsset",
                        "Memo":""
                    },
                    {
                        "Address":"EM2wjL3jgNHDZtR1e266V269n5WH6sYbCf",
                        "Txid":"920954e00bd1e1d3f674703c9e31988940c4c326382e13a22323d6e5ea3c4c6c",
                        "Type":"income",
                        "Value":50000000000,
                        "CreateTime":1533090125,
                        "Height":159257,
                        "Fee":0,
                        "Inputs":[
                            "8cTn9JAGXfqGgu8kVUaPBJXrhSjoJR9ymG"
                        ],
                        "Outputs":[
                            "EM2wjL3jgNHDZtR1e266V269n5WH6sYbCf"
                        ],
                        "TxType":"TransferAsset",
                        "Memo":""
                    }
                ],
                "TotalNum":2
            },
            "status":200
        }

Get spending address public key
------------------------------------------------

.. http:get:: /api/1/pubkey/(string:`addr`)

If we can get the public key of this adress.
   **Example request**:

   .. sourcecode:: http

      GET /api/1/pubkey/ELbKQrj8DTYn2gU7KBejcNWb4ix4EAGDmy HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":"02eda087df202cfc8904ec8f933bf20920251b3964b117c984a576c6fd9047073c",
            "status":200
        }

If we can not get the public key of this adress.
   **Example request**:

   .. sourcecode:: http

      GET /api/1/pubkey/EbxU18T3M9ufnrkRY7NLt6sKyckDW4VAsA HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":"Can not find pubkey of this address, please using this address send a transaction first",
            "status":200
        }

checking syncing status
------------------------------------------------

.. http:get:: /api/1/history/checking/sync

Current Synced height

   **Example request**:

   .. sourcecode:: http

      GET /api/1/history/checking/sync HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":100000,
            "status":200
        }

Get cmc price
------------------------------------------------

.. http:get:: /api/1/cmc?limit=3

Current Synced height

   **Example request**:

   .. sourcecode:: http

      GET /api/1/history/checking/sync HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result": [
                {
                    "24h_volume_btc": "1329678.25368338",
                    "24h_volume_cny": "107234082168.77766418",
                    "24h_volume_usd": "15185091926.81439972",
                    "available_supply": "17871650.00000000",
                    "id": "1",
                    "last_updated": "2019-08-12T03:58:32.000Z",
                    "market_cap_btc": "17871650.00000000",
                    "market_cap_cny": "1441288506661.53393555",
                    "market_cap_usd": "204096477762.26141357",
                    "max_supply": "21000000.00000000",
                    "name": "Bitcoin",
                    "num_market_pairs": "7803",
                    "percent_change_1h": "-0.11178000",
                    "percent_change_24h": "0.26175900",
                    "percent_change_7d": "0.04060490",
                    "platform_symbol": "",
                    "platform_token_address": "",
                    "price_btc": "1.00000000",
                    "price_cny": "80646.63904349",
                    "price_usd": "11420.12504510",
                    "rank": "1",
                    "symbol": "BTC",
                    "total_supply": "17871650.00000000"
                },
                {
                    "24h_volume_btc": "542188.65987348",
                    "24h_volume_cny": "43725693146.28816223",
                    "24h_volume_usd": "6191862293.79028988",
                    "available_supply": "107284875.62400000",
                    "id": "1027",
                    "last_updated": "2019-08-12T03:59:22.000Z",
                    "market_cap_btc": "2007866.35769792",
                    "market_cap_cny": "161927673396.82467651",
                    "market_cap_usd": "22930084878.75967789",
                    "max_supply": "0.00000000",
                    "name": "Ethereum",
                    "num_market_pairs": "5529",
                    "percent_change_1h": "0.06530480",
                    "percent_change_24h": "1.50229000",
                    "percent_change_7d": "-6.14715000",
                    "platform_symbol": "",
                    "platform_token_address": "",
                    "price_btc": "0.01871528",
                    "price_cny": "1509.32433351",
                    "price_usd": "213.73082408",
                    "rank": "2",
                    "symbol": "ETH",
                    "total_supply": "107284875.62400000"
                },
                {
                    "24h_volume_btc": "73729.05660548",
                    "24h_volume_cny": "5946000615.07916451",
                    "24h_volume_usd": "841995045.89186597",
                    "available_supply": "42872646068.00000000",
                    "id": "52",
                    "last_updated": "2019-08-12T03:59:05.000Z",
                    "market_cap_btc": "1133452.08639632",
                    "market_cap_cny": "91409101284.69120789",
                    "market_cap_usd": "12944164559.27545929",
                    "max_supply": "100000000000.00000000",
                    "name": "XRP",
                    "num_market_pairs": "439",
                    "percent_change_1h": "-0.00627041",
                    "percent_change_24h": "0.90267900",
                    "percent_change_7d": "-6.21009000",
                    "platform_symbol": "",
                    "platform_token_address": "",
                    "price_btc": "0.00002644",
                    "price_cny": "2.13210776",
                    "price_usd": "0.30192129",
                    "rank": "3",
                    "symbol": "XRP",
                    "total_supply": "99991387607.00000000"
                }
            ],
            "status": 200
        }



Ethereum
=================================
using the following api ,we can get eth statistic.

Get transaction history
------------------------------------------------
transaction history of specific address

.. http:post:: /api/1/eth/history

   **Example request**:

   .. sourcecode:: http

    POST /api/1/eth/history HTTP/1.1
    Host: localhost

      {
            "id":"19",
            "account":"0x4ce978642500aa7b056A0329bfA8622f468211e6"
      }

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
          "message": "OK",
          "result": [
            {
              "blockHash": "0x6be299e7ad6a8c2f2e958e30f3804f9732c744ce5563fe26917b621de78c9d2c",
              "blockNumber": "7000510",
              "from": "0x3f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be",
              "gas": "21000",
              "to": "0xd68c7a32cdb5b1f8a9c540e7781fa370145af27c",
              "gasPrice": "40000000000",
              "hash": "0x6b3ba1970bb2e5148866868f30b522aa0d64940a6f10f95a9a99f5e2c885fc2f",
              "input": "0x",
              "nonce": "2858612",
              "transactionIndex": "18",
              "value": "485139180000000000",
              "gasUsed": "21000",
              "cumulativeGasUsed": "537088",
              "isError": "0",
              "contractAddress": "0x0000000000000000000000000000000000000000",
              "timeStamp": "1546474873",
              "confirmations": "11"
            },
            {
              "blockHash": "0xba0afb0c3cd2049026b9b75216fc4ac16672d8bcce67954e75e536ebffa6a901",
              "blockNumber": "7000517",
              "from": "0xd68c7a32cdb5b1f8a9c540e7781fa370145af27c",
              "gas": "50000",
              "to": "0xf7793d27a1b76cdf14db7c83e82c772cf7c92910",
              "gasPrice": "2000000000",
              "hash": "0x59865d793c58704ea09e23f351b6ee1e327c8c52ecda3f91dcc06b4a003e32d4",
              "input": "0x",
              "nonce": "0",
              "transactionIndex": "136",
              "value": "485039180000000000",
              "gasUsed": "21000",
              "cumulativeGasUsed": "7878416",
              "isError": "0",
              "contractAddress": "0x0000000000000000000000000000000000000000",
              "timeStamp": "1546474993",
              "confirmations": "4"
            }
          ],
          "status": 1
        }

.. http:get:: /api/1/eth/history?address=(string:`addr`)

get eth transaction history

   **Example request**:

   .. sourcecode:: http

      GET /api/1/eth/history?address=0x4ce978642500aa7b056A0329bfA8622f468211e6 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
          "message": "OK",
          "result": [
            {
              "blockHash": "0x6be299e7ad6a8c2f2e958e30f3804f9732c744ce5563fe26917b621de78c9d2c",
              "blockNumber": "7000510",
              "from": "0x3f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be",
              "gas": "21000",
              "to": "0xd68c7a32cdb5b1f8a9c540e7781fa370145af27c",
              "gasPrice": "40000000000",
              "hash": "0x6b3ba1970bb2e5148866868f30b522aa0d64940a6f10f95a9a99f5e2c885fc2f",
              "input": "0x",
              "nonce": "2858612",
              "transactionIndex": "18",
              "value": "485139180000000000",
              "gasUsed": "21000",
              "cumulativeGasUsed": "537088",
              "isError": "0",
              "contractAddress": "0x0000000000000000000000000000000000000000",
              "timeStamp": "1546474873",
              "confirmations": "11"
            },
            {
              "blockHash": "0xba0afb0c3cd2049026b9b75216fc4ac16672d8bcce67954e75e536ebffa6a901",
              "blockNumber": "7000517",
              "from": "0xd68c7a32cdb5b1f8a9c540e7781fa370145af27c",
              "gas": "50000",
              "to": "0xf7793d27a1b76cdf14db7c83e82c772cf7c92910",
              "gasPrice": "2000000000",
              "hash": "0x59865d793c58704ea09e23f351b6ee1e327c8c52ecda3f91dcc06b4a003e32d4",
              "input": "0x",
              "nonce": "0",
              "transactionIndex": "136",
              "value": "485039180000000000",
              "gasUsed": "21000",
              "cumulativeGasUsed": "7878416",
              "isError": "0",
              "contractAddress": "0x0000000000000000000000000000000000000000",
              "timeStamp": "1546474993",
              "confirmations": "4"
            }
          ],
          "status": 1
        }


Wrap ethereum rpc
------------------------------------------------
you can throw request to this api as long as it compatible with ethereum rpc reqeust param

.. http:post:: /api/1/eth/wrap

   **Example request**:

   .. sourcecode:: http

    POST /api/1/eth/wrap HTTP/1.1
    Host: localhost

      {
          "jsonrpc": "2.0",
          "method": "eth_getBalance",
          "params": [
            "0xD109D62aD5bfAcC8869e35B4c0261dE5976d3b69",
            "latest"
          ],
          "id": 5
      }

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
          "id":1,
          "jsonrpc": "2.0",
          "result": "0x0234c8a3397aab58" // 158972490234375000
      }

   **Example request**:

   .. sourcecode:: http

    POST /api/1/eth/wrap HTTP/1.1
    Host: localhost

      {
          "jsonrpc": "2.0",
          "method": "eth_sendRawTransaction",
          "params": [
            "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675..."
          ],
          "id": 1
      }

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
          "id":1,
          "jsonrpc": "2.0",
          "result": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331"
      }


Get token balance
------------------------------------------------

.. http:get:: /api/1/eth/token/balance

get token balance

   **Example request**:

   .. sourcecode:: http

      GET /api/1/eth/token/balance?address=0x289B44672d8499A51130d65d2087A151c4e45966&contractaddress=0xa8cac329f783edac931815c5466e283d48c9d7f7 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "result":"0",
            "status":"1",
            "message":"OK"
        }


Get supported erc20 Currencies
------------------------------------------------


.. http:get:: /api/1/eth/currencies

   **Example request**:

   .. sourcecode:: http

      GET /api/1/eth/currencies HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        [
            {
                "code":"1ST",
                "colors":[
                    "#f15a22",
                    "#f15a22"
                ],
                "contract_address":"0xAf30D2a7E90d7DC361c8C4585e9BB7D2F6f15bc7",
                "contract_info":{

                },
                "currency_id":"ethereum-mainnet:0xAf30D2a7E90d7DC361c8C4585e9BB7D2F6f15bc7",
                "is_supported":true,
                "name":"FirstBlood",
                "sale_address":"",
                "scale":18,
                "type":"erc20"
            }
        ]

Get erc20 transaction logs
------------------------------------------------


.. http:get:: /api/1/eth/getLogs

get erc20 transaction history

   **Example request**:

   .. sourcecode:: http

      GET /api/1/eth/getLogs?fromBlock=0&toBlock=latest&topic0=0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359&topic1=0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1&topic1_2_opr=or&topic2=0x000000000000000000000000829bd824b016326a401d083b33d092293333a830 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
            "message":"OK",
            "result":[
                {
                    "address":"0xc5bbae50781be1669306b9e001eff57a2957b09d",
                    "blockNumber":"0x86ad07",
                    "data":"0x000000000000000000000000000000000000000000000000000000005bfff0e0",
                    "logIndex":"0x1e",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0xb45dff89f9a59fa0e6f91faff6c20fc57e0520706a4bb966c42e9466074f77db",
                    "transactionIndex":"0x23",
                    "gasUsed":"0x920b",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5db684cd"
                },
                {
                    "address":"0xc5bbae50781be1669306b9e001eff57a2957b09d",
                    "blockNumber":"0x86b872",
                    "data":"0x000000000000000000000000000000000000000000000000000000005b3b1a40",
                    "logIndex":"0x3c",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0xc822deb160b99471d8aa140409a1961cbc57750b50c14da6dd646b805689c714",
                    "transactionIndex":"0x51",
                    "gasUsed":"0x920b",
                    "gasPrice":"0x2cb417800",
                    "timeStamp":"0x5db71f7d"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x86c3e1",
                    "data":"0x0000000000000000000000000000000000000000000000b8bf8037a463400000",
                    "logIndex":"0x25",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0x108bcdc5305d647340a9f909640b1933bfd7abcf3e9bc6699bd0fd7fb91f8a17",
                    "transactionIndex":"0x33",
                    "gasUsed":"0x92c8",
                    "gasPrice":"0x306dc4200",
                    "timeStamp":"0x5db7c0a2"
                },
                {
                    "address":"0xc5bbae50781be1669306b9e001eff57a2957b09d",
                    "blockNumber":"0x86f9d6",
                    "data":"0x000000000000000000000000000000000000000000000000000000002909ad20",
                    "logIndex":"0x60",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0xcf27c1152e010cef309fbe0fe64829f0c47c118dc36186c8f76887c68a3f4919",
                    "transactionIndex":"0x4f",
                    "gasUsed":"0x920b",
                    "gasPrice":"0x28fa6ae00",
                    "timeStamp":"0x5dbaadf6"
                },
                {
                    "address":"0xc5bbae50781be1669306b9e001eff57a2957b09d",
                    "blockNumber":"0x86fa0f",
                    "data":"0x000000000000000000000000000000000000000000000000000000002909ad20",
                    "logIndex":"0x16",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0x49d1d42af45cb9a58c2716c24808150843b0d7767b9df47aadd9773687b2cd8c",
                    "transactionIndex":"0x26",
                    "gasUsed":"0x920b",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dbab0ef"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x87007b",
                    "data":"0x00000000000000000000000000000000000000000000020545d9ec5139580000",
                    "logIndex":"0xa8",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0x61d34afd1a726664617b18fb82c22851f64507ab679eac071e4b4aa6ab6aa134",
                    "transactionIndex":"0x82",
                    "gasUsed":"0x9308",
                    "gasPrice":"0x2cb417800",
                    "timeStamp":"0x5dbb093b"
                },
                {
                    "address":"0xc5bbae50781be1669306b9e001eff57a2957b09d",
                    "blockNumber":"0x870172",
                    "data":"0x000000000000000000000000000000000000000000000000000000004a109240",
                    "logIndex":"0x13",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0xb589a8efe116d419d925ee5f3a9a20576582700bac36a634d0959f70ee50107f",
                    "transactionIndex":"0x16",
                    "gasUsed":"0x920b",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dbb17fa"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x870172",
                    "data":"0x00000000000000000000000000000000000000000000020545d9ec5139580000",
                    "logIndex":"0x14",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0x3640df4f46f2182e0df53dc7f23406114cd5bbb105a6d66c6c84a24240da53eb",
                    "transactionIndex":"0x17",
                    "gasUsed":"0x9308",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dbb17fa"
                },
                {
                    "address":"0xb98d4c97425d9908e66e53a6fdf673acca0be986",
                    "blockNumber":"0x8720c0",
                    "data":"0x00000000000000000000000000000000000000000000019a35386850d4dc0000",
                    "logIndex":"0x44",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0x94f3995feca11737e938e00c0c40137a68fb1335992829f38eb21bbfc44019f1",
                    "transactionIndex":"0x85",
                    "gasUsed":"0x923d",
                    "gasPrice":"0x28fa6ae00",
                    "timeStamp":"0x5dbcc694"
                },
                {
                    "address":"0xb98d4c97425d9908e66e53a6fdf673acca0be986",
                    "blockNumber":"0x872172",
                    "data":"0x00000000000000000000000000000000000000000000019a35386850d4dc0000",
                    "logIndex":"0x19",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0x13fd9e243152e06eb7869f75d1f00601762bfa950fc4bf1c816334a42878992c",
                    "transactionIndex":"0x1d",
                    "gasUsed":"0x923d",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dbcd2d0"
                },
                {
                    "address":"0xb98d4c97425d9908e66e53a6fdf673acca0be986",
                    "blockNumber":"0x8732e2",
                    "data":"0x00000000000000000000000000000000000000000000019d5a21cd04c1840000",
                    "logIndex":"0x2",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0x36798214a5e21bb14f6dcba4219251e1bdafffc8832a922c3cfc17ca2eda12fd",
                    "transactionIndex":"0x5",
                    "gasUsed":"0x923d",
                    "gasPrice":"0x28fa6ae00",
                    "timeStamp":"0x5dbdc6b3"
                },
                {
                    "address":"0xb98d4c97425d9908e66e53a6fdf673acca0be986",
                    "blockNumber":"0x8733f1",
                    "data":"0x00000000000000000000000000000000000000000000019d5a21cd04c1840000",
                    "logIndex":"0xd",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0x96ad9a34a0b6fd4eebe6e01754390306f3dad73fb844552b176fb20dcedb7867",
                    "transactionIndex":"0x2f",
                    "gasUsed":"0x923d",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dbdd52a"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x87398f",
                    "data":"0x000000000000000000000000000000000000000000000250addaa67bc1b40000",
                    "logIndex":"0x11",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0x345306d82d306ff45508c7cfff9d4eaf112220cc3d3279a85be828d1ccb656aa",
                    "transactionIndex":"0x15",
                    "gasUsed":"0x9308",
                    "gasPrice":"0x28fa6ae00",
                    "timeStamp":"0x5dbe2237"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x8739fb",
                    "data":"0x000000000000000000000000000000000000000000000250addaa67bc1b40000",
                    "logIndex":"0x3a",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0xdfe5442da981a755231bcb5472cbe93987b94c8feb0995d96385c54d3d0a0539",
                    "transactionIndex":"0x59",
                    "gasUsed":"0x9308",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dbe28f5"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x876ea2",
                    "data":"0x0000000000000000000000000000000000000000000000c5a66a12aa02380000",
                    "logIndex":"0x58",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000505387014d6518d5daff534a14d91650f32c9fd6",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1"
                    ],
                    "transactionHash":"0xc7eb63cbc2d7d9c0217f46ecf4dd2b35dd23fa83b39c20297ad4fa39457c3fc9",
                    "transactionIndex":"0x71",
                    "gasUsed":"0x92c8",
                    "gasPrice":"0x28fa6ae00",
                    "timeStamp":"0x5dc104e9"
                },
                {
                    "address":"0xbf2179859fc6d5bee9bf9158632dc51678a4100e",
                    "blockNumber":"0x876fff",
                    "data":"0x0000000000000000000000000000000000000000000000c5a66a12aa02380000",
                    "logIndex":"0x23",
                    "topics":[
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x000000000000000000000000e418a0e203f36cb843079f6ebf0b367e48774ac1",
                        "0x0000000000000000000000006cc5f688a315f3dc28a7781717a9a798a59fda7b"
                    ],
                    "transactionHash":"0x4bde2fb7087c9ec6e683dae80693986462efbfb34781418ba44784558a003eb9",
                    "transactionIndex":"0x2b",
                    "gasUsed":"0x92c8",
                    "gasPrice":"0x826299e00",
                    "timeStamp":"0x5dc1175c"
                }
            ],
            "status":"1"
        }
