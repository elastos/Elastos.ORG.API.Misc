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