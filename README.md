### Spec:
    - Create an http REST API server in GO for uploading data and then upload the following payload into a cloud storage such as s3 or google cloud storage.
    - The payload size limit is 10KB
    - API can still handle/recieve requests from clients without having to wait for the uploading process to complete

### Reqirements:
    - Clean code
    - Api can handle a large volume of requests & high traffic
    - Write unit tests for testing your APIs' performance
    - Explain your design

### Example:

    `curl -X POST http://localhost:8080/user/batch --data "@payload.json" -H "Content-Type: application/json"`

payload.json content: 

   ```"batch": [
    {
      "timestamp": "2021-10-22T10:47:13.749627+07:00",
      "name": "bidRequest",
      "requestId": "sample-id",
      "context": {
        "id": "6e3803cd726ef184e996-5",
        "imp": [
          {
            "id": "b0c813e2-81de-bb22-102f37575a9e",
            "banner": {
              "format": [
                {
                  "w": 300,
                  "h": 250
                }
              ],
              "api": [
                5
              ]
            },
            "secure": 1,
            "ext": {
              "prebid": {},
              "aic_ads": {
                "header_bids": [],
                "container_id": "cdb714a7-8f99-480f",
                "inventory_id": "b0c813e2-818c-490f"
              }
            }
          }
        ],
        "site": {
          "domain": "www.vnexpress.net",
          "page": "www.vnexpress.net/index.html"
        },
        "device": {
          "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
          "osv": "14.2",
          "ifa": "00000000-0000-0000-0000-000000000000"
        },
        "user": {
          "gender": "O"
        },
        "at": 2,
        "tmax": 400,
        "ext": {}
      },
      "writerKey": "cdb714a7-8f99-480f-9bfb-01d9116a12b3@ios",
      "anonymousId": "22c44597-d69f-48a9-8e23-129af884420d"
    },
    {
      "context": {
        "id": "6e3803cd726ef184e996-5",
        "seatbid": [
          {
            "bid": [
              {
                "id": "0_6e3803cd726ef184e996-5_ad80ea94-d049-459d-9a99-a3ee6f3b56fc",
                "impid": "b0c813e2-818c-490f-bb22-102f37575a9e",
                "price": 0.001,
                "adm": "<div id='aic_144472732'><script src='http://localhost:8080/serve-bid/_52XC33XS7i23Va5Yf4Y4ubw7inMoszzBp129U0490dzcoGpexitOna2CKV3mrDg8YBvoqQ9POFAAPHc0goZFsDZ3JkFYC2z63x4AQYuNmnm2ybu71KNcWE_TD5-iLlXQDx7EKMz_HlllpK8OHstSUJa0RhpL9aeUlpm8AU50KVIiI4BN8CDstYlcH1DrI3PSS9NYhO_BFEbIQeZQfe9vLo-5BwM5qKHfYopSaJKdYpUgDT4VFVCSIZ_30mM8gJLyWfZ2If3FTIP2qR3fHakNg-0IF2Om4vdus9u_1oIKenk9ssiY3bR10O_PgWQYsS7Aa-E2LPWmrAVfQ9j8VxAYw==/serve-bid.js?rp={{AUCTION_PRICE}}&gdpr_consent=${GDPR_CONSENT_234}' charset='utf-8'></script></div>",
                "adomain": [
                  "aicactus.io"
                ],
                "crid": "1",
                "ext": {
                  "prebid": {
                    "type": ""
                  }
                }
              }
            ],
            "seat": "corebid"
          }
        ],
        "cur": "USD",
        "ext": {
          "responsetimemillis": {
            "aic_ads": 0
          },
          "tmaxrequest": 400,
          "prebid": {
            "auctiontimestamp": 1634874433747
          }
        }
      },
      "writerKey": "cdb714a7-8f99-480f-9bfb-01d9116a12b3@ios",
      "anonymousId": "22c44597-d69f-48a9-8e23-129af884420d",
      "timestamp": "2021-10-22T10:47:13.749628+07:00",
      "name": "bidResponse",
      "requestId": "6e3803cd726ef184e996-5"
    },
    {
      "writerKey": "cdb714a7-8f99-480f-9bfb-01d9116a12b3@ios",
      "anonymousId": "22c44597-d69f-48a9-8e23-129af884420d",
      "timestamp": "2021-10-22T10:47:13.749628+07:00",
      "name": "bidResponse",
      "requestId": "6e3803cd726ef184e996-5",
      "context": {
        "id": "6e3803cd726ef184e996-5",
        "seatbid": [
          {
            "bid": [
              {
                "id": "0_6e3803cd726ef184e996-5_ad80ea94-d049-459d-9a99-a3ee6f3b56fc",
                "impid": "b0c813e2-818c-490f-bb22-102f37575a9e",
                "price": 0.001,
                "adm": "<div id='aic_144472732'><script src='http://localhost:8080/serve-bid/_52XC33XS7i23Va5Yf4Y4ubw7inMoszzBp129U0490dzcoGpexitOna2CKV3mrDg8YBvoqQ9POFAAPHc0goZFsDZ3JkFYC2z63x4AQYuNmnm2ybu71KNcWE_TD5-iLlXQDx7EKMz_HlllpK8OHstSUJa0RhpL9aeUlpm8AU50KVIiI4BN8CDstYlcH1DrI3PSS9NYhO_BFEbIQeZQfe9vLo-5BwM5qKHfYopSaJKdYpUgDT4VFVCSIZ_30mM8gJLyWfZ2If3FTIP2qR3fHakNg-0IF2Om4vdus9u_1oIKenk9ssiY3bR10O_PgWQYsS7Aa-E2LPWmrAVfQ9j8VxAYw==/serve-bid.js?rp={{AUCTION_PRICE}}&gdpr_consent=${GDPR_CONSENT_234}' charset='utf-8'></script></div>",
                "adomain": [
                  "aicactus.io"
                ],
                "crid": "1",
                "ext": {
                  "prebid": {
                    "type": ""
                  }
                }
              }
            ],
            "seat": "aic_ads"
          }
        ],
        "cur": "USD",
        "ext": {
          "responsetimemillis": {
            "aic_ads": 0
          },
          "tmaxrequest": 400,
          "prebid": {
            "auctiontimestamp": 1634874433747
          }
        }
      }
    }
  ],
  "writerKey": "cdb714a7-8f99-48@ios",
  "typeSDK": "adnetwork",
  "sentAt": "2021-10-22T10:47:13.749626+07:00"
}