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

Get dpos reward of specific height
------------------------------------------------
voter reward

.. http:get:: /api/1/dpos/reward/height/(int:`height`)

   **Example request**:

   .. sourcecode:: http

      GET /api/1/dpos/reward/height/241762 HTTP/1.1
      Host: localhost

   **Example response**:

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

        {
          "result":[
              {
                  "Address":"8ZNizBf4KhhPjeJRGpox6rPcHE5Np6tFx3",
                  "Value":150729930,
                  "CreateTime":1552781309,
                  "Height":241762
              },
              {
                  "Address":"EaA6CXrtvG1CpethsoodvdTwpeCtWprZzH",
                  "Value":150729929,
                  "CreateTime":1552781309,
                  "Height":241762
              },
              {
                  "Address":"EXr8pYTR5Z56Ni9Vg9r5UnNg2MCQ6enm93",
                  "Value":150729928,
                  "CreateTime":1552781309,
                  "Height":241762
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
                  "Value":"356976",
                  "Rank":1
              },
              {
                  "Producer_public_key":"033c495238ca2b6bb8b7f5ae172363caea9a55cf245ffb3272d078126b1fe3e7cd",
                  "Value":"311262.3",
                  "Rank":2
              },
              {
                  "Producer_public_key":"0337e6eaabfab6321d109d48e135190560898d42a1d871bfe8fecc67f4c3992250",
                  "Value":"309791",
                  "Rank":3
              },
              {
                  "Producer_public_key":"03c78467b91805c95ada2530513069bef1f1f1e7b756861381ab534efa6d94e40a",
                  "Value":"218062.55555",
                  "Rank":4
              },
              {
                  "Producer_public_key":"021d59a84d2243111e39e8c2af0a5089127d142d52b18c3e4bf744e0c6f8af44e0",
                  "Value":"147157",
                  "Rank":5
              },
              {
                  "Producer_public_key":"036417ab256114a32bcff38f3e10f0384cfa9238afa41a163017687b3ce1fa17f2",
                  "Value":"139506",
                  "Rank":6
              },
              {
                  "Producer_public_key":"02e578a6f4295765ad3be4cdac9be15de5aedaf1ae76e86539bb54c397e467cd5e",
                  "Value":"125831",
                  "Rank":7
              },
              {
                  "Producer_public_key":"02ddd829f3495a2ce76d908c3e6e7d4505e12c4718c5af4b4cbff309cfd3aeab88",
                  "Value":"108893",
                  "Rank":8
              },
              {
                  "Producer_public_key":"03c7b1f234d5d16472fcdd24d121e4cd224e1074f558a3eb1a6a146aa91dcf9c0d",
                  "Value":"108111",
                  "Rank":9
              },
              {
                  "Producer_public_key":"03b688e0124580de452c400e01c628a690527e8742b6fa4645026dbc70155d7c8b",
                  "Value":"107788",
                  "Rank":10
              },
              {
                  "Producer_public_key":"03bc2c2b75009a3a551e98bf206730501ecdf46e71b0405840ff1d5750094bd4ff",
                  "Value":"104972",
                  "Rank":11
              },
              {
                  "Producer_public_key":"0230d383546d154d67cfafc6091c0736c0b26a8c7c16e879ef8011d91df976f1fb",
                  "Value":"104181",
                  "Rank":12
              },
              {
                  "Producer_public_key":"028fb1a85f6a30a516b9e3516d03267403a3af0c96d73b0284ca0c1165318531ff",
                  "Value":"99993",
                  "Rank":13
              },
              {
                  "Producer_public_key":"02db921cfb4bf504c83038212aafe52cc1d0a07eb71a399a0d2162fe0cd4d47720",
                  "Value":"98996",
                  "Rank":14
              },
              {
                  "Producer_public_key":"033fb33f39276b93d3474cf7999887bed16c3211ee7f904399eeead4c480d7d592",
                  "Value":"98784",
                  "Rank":15
              },
              {
                  "Producer_public_key":"030e4b487daf8e14dbd7023e3f6f475d00145a1f1cc87be4b8d58a4291ab0a3b1a",
                  "Value":"25899",
                  "Rank":16
              },
              {
                  "Producer_public_key":"0234048d3ee92a7d34fbe3da22bc69583b1785e8f6684c9f4f11804c518cb4e53d",
                  "Value":"678",
                  "Rank":17
              },
              {
                  "Producer_public_key":"02771568d40c1b20f3cbc2f4de327d3f61ae1a97a3e4a014838d267c818f2f999e",
                  "Value":"666",
                  "Rank":18
              },
              {
                  "Producer_public_key":"0203c80103bb094b5870f6b99b0bc6ab857fa87bab1896fc845108bba7aafbfe3c",
                  "Value":"666",
                  "Rank":19
              },
              {
                  "Producer_public_key":"0210694f4ab518037bc2dcc3f5e1a1030e8a36821ab019c10f29d4a894b8034498",
                  "Value":"666",
                  "Rank":20
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