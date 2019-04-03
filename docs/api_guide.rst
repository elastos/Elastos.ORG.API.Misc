Getting started with the MISC API
########################################

.. toctree::
  :maxdepth: 3

Introduction
=============
MISC has a Restful API with URL endpoints corresponding to actions that users can perform with their channels. The endpoints accept and return JSON encoded objects. The API URL path always contains the API version in order to differentiate queries to different API versions. All queries start with: ``/api/<version>/`` where ``<version>`` is an integer representing the current API version.

.. api:

Dpos Vote
=================================
using the following api ,we can get dpos vote statistic.

Get dpos producer vote statistics
------------------------------------------------
producer's vote statistics

.. http:get:: /api/1/dpos/producer/(string:`producer_public_key`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/dpos/producer/03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800 HTTP/1.1
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
                    "Producer_public_key":"03330ee8520088b7f578a9afabaef0c034fa31fe1354cb3a14410894f974132800",
                    "Vote_type":"Delegate",
                    "Txid":"fbc81da6db6db5cb09c76fe405cf238353a8e837dda5acacd137ba43a9da1d02",
                    "N":0,
                    "Value":"9999",
                    "Outputlock":0,
                    "Address":"ENaaqePNBtrZsNbs9uc35CPqTbvn8oaYL9",
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
                "Producer_public_key":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                "Value":"311315.3",
                "Address":"Eb8UHkQ2bJ4Ljux4yBePFdxB5Yp77VYHyt",
                "Rank":1,
                "Ownerpublickey":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                "Nodepublickey":"03c18abb98f6679064bd44121f3b0a3f25dea1a8b8cb0e1b51dc9c26729f07ddc9",
                "Nickname":"我怎么这么好看",
                "Url":"www.douniwan.com",
                "Location":263,
                "Active":false,
                "Votes":"311326.30000000",
                "Netaddress":"8.8.8.8",
                "State":"Activate",
                "Registerheight":232288,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":1,
                "Reward":""
            },
            {
                "Producer_public_key":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                "Value":"309844",
                "Address":"EdhP91WcY2WhyV8N6dCnBxbjAnGd2izrzY",
                "Rank":2,
                "Ownerpublickey":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                "Nodepublickey":"ff",
                "Nickname":"今天真好",
                "Url":"www.helloword.com",
                "Location":44,
                "Active":false,
                "Votes":"309855",
                "Netaddress":"1.2.3.4",
                "State":"Activate",
                "Registerheight":234800,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":2,
                "Reward":""
            },
            {
                "Producer_public_key":"03c78467b91805c95ada2530513069bef1f1f1e7b756861381ab534efa6d94e40a",
                "Value":"218115.55555",
                "Address":"EdfJA92nN9X4T9cKqkvyrunVuBWfF1Mumm",
                "Rank":3,
                "Ownerpublickey":"03c78467b91805c95ada2530513069bef1f1f1e7b756861381ab534efa6d94e40a",
                "Nodepublickey":"fffff3fffffffffffffffbff1affffffec",
                "Nickname":"聪聪2",
                "Url":"1.4.7.9",
                "Location":672,
                "Active":false,
                "Votes":"218126.55555000",
                "Netaddress":"1.12.3.4",
                "State":"Activate",
                "Registerheight":233035,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":3,
                "Reward":""
            },
            {
                "Producer_public_key":"021d59a84d2243111e39e8c2af0a5089127d142d52b18c3e4bf744e0c6f8af44e0",
                "Value":"147210",
                "Address":"ESpTiKXgLcYkzxdD7MuCmL9y9fbWrnH591",
                "Rank":4,
                "Ownerpublickey":"021d59a84d2243111e39e8c2af0a5089127d142d52b18c3e4bf744e0c6f8af44e0",
                "Nodepublickey":"ffff1230ffff",
                "Nickname":"www.12306.cn",
                "Url":"www.12306.cn",
                "Location":244,
                "Active":false,
                "Votes":"147221",
                "Netaddress":"www.12306.cn",
                "State":"Activate",
                "Registerheight":232899,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":4,
                "Reward":""
            },
            {
                "Producer_public_key":"036417ab256114a32bcff38f3e10f0384cfa9238afa41a163017687b3ce1fa17f2",
                "Value":"139859",
                "Address":"ETKVMhhQCjttNAjrbqmkAAYuYshLdaDnjm",
                "Rank":5,
                "Ownerpublickey":"036417ab256114a32bcff38f3e10f0384cfa9238afa41a163017687b3ce1fa17f2",
                "Nodepublickey":"03e5b45b44bb1e2406c55b7dd84b727fad608ba7b7c11a9c5ffbfee60e427bd1da",
                "Nickname":"聪聪3",
                "Url":"225.7.3",
                "Location":672,
                "Active":false,
                "Votes":"139870",
                "Netaddress":"1.1.1.8",
                "State":"Activate",
                "Registerheight":233537,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":5,
                "Reward":""
            },
            {
                "Producer_public_key":"02e578a6f4295765ad3be4cdac9be15de5aedaf1ae76e86539bb54c397e467cd5e",
                "Value":"125884",
                "Address":"EHdSBUH3nxkcAk9evU4HrENzEm8MHirkkN",
                "Rank":6,
                "Ownerpublickey":"02e578a6f4295765ad3be4cdac9be15de5aedaf1ae76e86539bb54c397e467cd5e",
                "Nodepublickey":"fffeffddfffffff2fffffffffbffffffff",
                "Nickname":"亦来云",
                "Url":"www.yilaiyun.com",
                "Location":244,
                "Active":false,
                "Votes":"125895",
                "Netaddress":"www.yilaiyun.com",
                "State":"Activate",
                "Registerheight":233680,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":6,
                "Reward":""
            },
            {
                "Producer_public_key":"02ddd829f3495a2ce76d908c3e6e7d4505e12c4718c5af4b4cbff309cfd3aeab88",
                "Value":"108946",
                "Address":"EevRwpP5GYz5s8fuMboUnhsAQVVKbyJSph",
                "Rank":7,
                "Ownerpublickey":"02ddd829f3495a2ce76d908c3e6e7d4505e12c4718c5af4b4cbff309cfd3aeab88",
                "Nodepublickey":"ffffffffffffffffffffffffffffffffffff",
                "Nickname":"曲率区动",
                "Url":"www.bightbc.com",
                "Location":86,
                "Active":false,
                "Votes":"108957",
                "Netaddress":"EfSkh3e9uaVN5iMdU7oUPYPmyMxrMsrDut",
                "State":"Activate",
                "Registerheight":234283,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":7,
                "Reward":""
            },
            {
                "Producer_public_key":"03c7b1f234d5d16472fcdd24d121e4cd224e1074f558a3eb1a6a146aa91dcf9c0d",
                "Value":"108164",
                "Address":"EQR8f9y2Sd5gFG3LWEeC57qXc2yEnDhgm2",
                "Rank":8,
                "Ownerpublickey":"03c7b1f234d5d16472fcdd24d121e4cd224e1074f558a3eb1a6a146aa91dcf9c0d",
                "Nodepublickey":"350181ff",
                "Nickname":"范冰冰",
                "Url":"1.8.5.8",
                "Location":86,
                "Active":false,
                "Votes":"108175",
                "Netaddress":"HTTP//HUANGBINGBING.COM",
                "State":"Activate",
                "Registerheight":233676,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":8,
                "Reward":""
            },
            {
                "Producer_public_key":"03b688e0124580de452c400e01c628a690527e8742b6fa4645026dbc70155d7c8b",
                "Value":"107841",
                "Address":"EQHz2jPpgW8trYD4ejYgfi4sE4JSTf7m9N",
                "Rank":9,
                "Ownerpublickey":"03b688e0124580de452c400e01c628a690527e8742b6fa4645026dbc70155d7c8b",
                "Nodepublickey":"ffffffffffff",
                "Nickname":"基延一族",
                "Url":"1.4.7.9",
                "Location":672,
                "Active":false,
                "Votes":"107852",
                "Netaddress":"www.vogue.com",
                "State":"Activate",
                "Registerheight":233684,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":9,
                "Reward":""
            },
            {
                "Producer_public_key":"03bc2c2b75009a3a551e98bf206730501ecdf46e71b0405840ff1d5750094bd4ff",
                "Value":"105025",
                "Address":"ENxPtTR7Jn1kxhdTXedF28s3iz6djYfRaS",
                "Rank":10,
                "Ownerpublickey":"03bc2c2b75009a3a551e98bf206730501ecdf46e71b0405840ff1d5750094bd4ff",
                "Nodepublickey":"fffffffd29fffffffafff8fafffffdfffa",
                "Nickname":"乐天居士",
                "Url":"www.baidu.com",
                "Location":376,
                "Active":false,
                "Votes":"105036",
                "Netaddress":"尽快哦孩子",
                "State":"Activate",
                "Registerheight":232892,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":10,
                "Reward":""
            },
            {
                "Producer_public_key":"0230d383546d154d67cfafc6091c0736c0b26a8c7c16e879ef8011d91df976f1fb",
                "Value":"104234",
                "Address":"EMyStHAvvy1VLsLyow8uMRW4kUYLeGXF17",
                "Rank":11,
                "Ownerpublickey":"0230d383546d154d67cfafc6091c0736c0b26a8c7c16e879ef8011d91df976f1fb",
                "Nodepublickey":"fffffffffffefffffffffffffbfcffffff",
                "Nickname":"烽火",
                "Url":"www.ela.com",
                "Location":86,
                "Active":false,
                "Votes":"104245",
                "Netaddress":"www.ela.com",
                "State":"Activate",
                "Registerheight":233612,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":11,
                "Reward":""
            },
            {
                "Producer_public_key":"028fb1a85f6a30a516b9e3516d03267403a3af0c96d73b0284ca0c1165318531ff",
                "Value":"104044",
                "Address":"ESqyiCizgyNNLKdVQhhtxtR5v5eCnkk3Qh",
                "Rank":12,
                "Ownerpublickey":"028fb1a85f6a30a516b9e3516d03267403a3af0c96d73b0284ca0c1165318531ff",
                "Nodepublickey":"ffff9262",
                "Nickname":"链世界",
                "Url":"www.7234.cn",
                "Location":86,
                "Active":false,
                "Votes":"101056",
                "Netaddress":"www.7234.cn",
                "State":"Activate",
                "Registerheight":235373,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":12,
                "Reward":""
            },
            {
                "Producer_public_key":"02db921cfb4bf504c83038212aafe52cc1d0a07eb71a399a0d2162fe0cd4d47720",
                "Value":"99029",
                "Address":"ERbFZNj5bukyRQe5G4gdXnbDqVyxcTNeFT",
                "Rank":13,
                "Ownerpublickey":"02db921cfb4bf504c83038212aafe52cc1d0a07eb71a399a0d2162fe0cd4d47720",
                "Nodepublickey":"1234567890ffdffffffffcffffffffffffff",
                "Nickname":"ios_us01",
                "Url":"www.ios_us01.com",
                "Location":684,
                "Active":false,
                "Votes":"99040",
                "Netaddress":"192.168.1.22:25339",
                "State":"Activate",
                "Registerheight":233672,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":13,
                "Reward":""
            },
            {
                "Producer_public_key":"033fb33f39276b93d3474cf7999887bed16c3211ee7f904399eeead4c480d7d592",
                "Value":"98837",
                "Address":"EXQZMbKMcmVmwv25AYbrzWPhFRSfqKcfKM",
                "Rank":14,
                "Ownerpublickey":"033fb33f39276b93d3474cf7999887bed16c3211ee7f904399eeead4c480d7d592",
                "Nodepublickey":"19fffffe9dfffafffffffffffbcaffffff",
                "Nickname":"晓黎-评财经",
                "Url":"www.pingcj.com",
                "Location":86,
                "Active":false,
                "Votes":"98848",
                "Netaddress":"Ed846C7M9Ax8x1qaftjSR53RZmfSvp8CpN",
                "State":"Activate",
                "Registerheight":235077,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":14,
                "Reward":""
            },
            {
                "Producer_public_key":"030e4b487daf8e14dbd7023e3f6f475d00145a1f1cc87be4b8d58a4291ab0a3b1a",
                "Value":"25952",
                "Address":"EVFSvWoxiyvGLka4V6Wt394LEoUu8mDhk4",
                "Rank":15,
                "Ownerpublickey":"030e4b487daf8e14dbd7023e3f6f475d00145a1f1cc87be4b8d58a4291ab0a3b1a",
                "Nodepublickey":"0241db65a4da2cdcbb648a881ced2a5ed64646ecc3a2cc9a75cec2853de61dbed1",
                "Nickname":"ELASuperNode",
                "Url":"www.ELASuperNode.com",
                "Location":86,
                "Active":false,
                "Votes":"25963",
                "Netaddress":"54.64.220.165",
                "State":"Activate",
                "Registerheight":237877,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":15,
                "Reward":""
            },
            {
                "Producer_public_key":"038796d13f0ed94b2587ba2e13ca99b3cafd4d5cea2b08b2d06b841ed10d177a51",
                "Value":"53",
                "Address":"ENXdHHxBv3doDeYPLk4yWSU91vpM9rfEDr",
                "Rank":16,
                "Ownerpublickey":"038796d13f0ed94b2587ba2e13ca99b3cafd4d5cea2b08b2d06b841ed10d177a51",
                "Nodepublickey":"0365b130cf561f04d98a56e2b3acab957fb90ab168d5f0bf9befae7413481168d0",
                "Nickname":"知春里",
                "Url":"www.1559.cn",
                "Location":1441,
                "Active":false,
                "Votes":"64",
                "Netaddress":"18.106.19.20",
                "State":"Activate",
                "Registerheight":238493,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":16,
                "Reward":""
            },
            {
                "Producer_public_key":"0210694f4ab518037bc2dcc3f5e1a1030e8a36821ab019c10f29d4a894b8034498",
                "Value":"33",
                "Address":"ESwKtu2aYSHHfdWUPdg4b3PtibfaEcJEvT",
                "Rank":17,
                "Ownerpublickey":"0210694f4ab518037bc2dcc3f5e1a1030e8a36821ab019c10f29d4a894b8034498",
                "Nodepublickey":"024babfecea0300971a6f0ad13b27519faff0ef595faf9490dc1f5f4d6e6d7f3fb",
                "Nickname":"adr_us01",
                "Url":"www.adr_us01_9.com",
                "Location":93,
                "Active":false,
                "Votes":"44",
                "Netaddress":"node-regtest-509.eadd.co:26339",
                "State":"Activate",
                "Registerheight":238437,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":17,
                "Reward":""
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
                "Index":18,
                "Reward":""
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
                "Index":19,
                "Reward":""
            },
            {
                "Address":"EQVa3pi9K3AkEfapm6cQHGbanyeZE715TW",
                "Rank":20,
                "Ownerpublickey":"024ce9f03960338bcbd82189f4ca6565ffdfce9f300c48d62f5df67d9591f13388",
                "Nodepublickey":"024ce9f03960338bcbd82189f4ca6565ffdfce9f300c48d62f5df67d9591f13388",
                "Nickname":"ios_us08",
                "Url":"www.ios_us08.com",
                "Location":1268,
                "Active":false,
                "Votes":"0",
                "Netaddress":"172.31.40.70:25339",
                "State":"Activate",
                "Registerheight":245547,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":20,
                "Reward":""
            },
            {
                "Address":"ESpgUHUQYKK1KRb6rEdNMz6tYfaE7pwZ9G",
                "Rank":21,
                "Ownerpublickey":"022fca820ec28b4f69b4896ecc7b62b454a93e93b047df6a7c1599e6dcbeb5c16f",
                "Nodepublickey":"022fca820ec28b4f69b4896ecc7b62b454a93e93b047df6a7c1599e6dcbeb5c16f",
                "Nickname":"adr_us12",
                "Url":"www.adr_us1212.com",
                "Location":374,
                "Active":false,
                "Votes":"0",
                "Netaddress":"172.31.32.36:25339",
                "State":"Activate",
                "Registerheight":252595,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":21,
                "Reward":""
            },
            {
                "Address":"EfRqyhv82oexrn92wSXv8CUQeKr7ryf43D",
                "Rank":22,
                "Ownerpublickey":"02e2011494dc7dd5ae65199ae81e6dc0c4a3aaf1220d15ae50a283f0a19c6ba481",
                "Nodepublickey":"02a6272c0db95090f0263e6fd4aeacbccb5b2a2a68df24bddd7a246954845a1a76",
                "Nickname":"ios_us12",
                "Url":"www.ios_us1212.com",
                "Location":244,
                "Active":false,
                "Votes":"0",
                "Netaddress":"172.89.90.91:25339",
                "State":"Activate",
                "Registerheight":252744,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":22,
                "Reward":""
            },
            {
                "Address":"ES197k9gNSLqAu3iJxBiC58RNnusazToQw",
                "Rank":23,
                "Ownerpublickey":"02067481822504f3066c729938ba574dd1e40c8adc5654daf82aafbe3f9e968545",
                "Nodepublickey":"02067481822504f3066c729938ba574dd1e40c8adc5654daf82aafbe3f9e968545",
                "Nickname":"adr_us11",
                "Url":"www.adr_us11.com",
                "Location":93,
                "Active":false,
                "Votes":"0",
                "Netaddress":"172.168.21.32:25339",
                "State":"Activate",
                "Registerheight":252790,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":23,
                "Reward":""
            },
            {
                "Address":"ETVeRXbrgYnu9EsWYrBUv6ysbuGHkknmCk",
                "Rank":24,
                "Ownerpublickey":"0394b945d853b78f12db0fde92d8cf4b468879d7b9fd2c57e0e309c0a280c4504c",
                "Nodepublickey":"0394b945d853b78f12db0fde92d8cf4b468879d7b9fd2c57e0e309c0a280c4504c",
                "Nickname":"bx",
                "Url":"nxnxj",
                "Location":93,
                "Active":false,
                "Votes":"0",
                "Netaddress":"hdjsj",
                "State":"Activate",
                "Registerheight":252607,
                "Cancelheight":0,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":24,
                "Reward":""
            },
            {
                "Address":"Ec6rtmnSQ5EiVBPL2kKgVk4m2zMsJggVV4",
                "Rank":25,
                "Ownerpublickey":"0363c7058dcb9d5267a01a72468e63a8e21d09a33bbcb9f7a9fd2923703f85f42a",
                "Nodepublickey":"0363c7058dcb9d5267a01a72468e63a8e21d09a33bbcb9f7a9fd2923703f85f42a",
                "Nickname":"test1111",
                "Url":"www.hh.com",
                "Location":93,
                "Active":false,
                "Votes":"0",
                "Netaddress":"www.baidu.com",
                "State":"Activate",
                "Registerheight":243378,
                "Cancelheight":243382,
                "Inactiveheight":0,
                "Illegalheight":0,
                "Index":25,
                "Reward":"140714931"
            }
        ],
        "status":200
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


Transaction History
=================================
using the following api ,we can get address transaction history.

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

DID Info
=================================
using the following api ,we can get did info.

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