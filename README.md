# Agrak

## How to execute?

First in the Service folder is the product service, which is communicated via gRPC. To test your endpoints, you must load the .proto file in postman and create a gRPC request.
The steps are:
* To lift the service by:  *go run .\api\main.go* 
* Load .proto file in postman
* Next try by postaman with next request (Create product)
` {
    "brand": "New Balance",
    "images": [
        "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/perritos-de-instagram-1616062929.jpg?crop=0.492xw:0.908xh;0.508xw,0.0917xh&resize=640:*"
    ],
    "name": "500 Zapatilla mujer",
    "price": 6123.4543,
    "principalImage": "https://home.ripley.cl/store/Attachment/WOP/D311/2000392630104/2000392630104_2.jpg",
    "size": "35",
    "sku": "FAL-840627987"
} `


### But the idea is to test the HTTP communication.
To test http communication

The steps are:

* Open composite folder
* Open service folder
* To lift the product service by:  *go run .\api\main.go*  with port :8080
* To lift the product service composite by *go run .\api\main.go* with port :8081
* In postman create a new request
* To create a product for example is the next request

`
METHOD: POST
`
`
URL:localhost:8081/product/
`

`
BODY: {
    "brand": "New Balance",
    "images": [
        "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/perritos-de-instagram-1616062929.jpg?crop=0.492xw:0.908xh;0.508xw,0.0917xh&resize=640:*"
    ],
    "name": "500 Zapatilla mujer",
    "price": 6123.4543,
    "principalImage": "https://home.ripley.cl/store/Attachment/WOP/D311/2000392630104/2000392630104_2.jpg",
    "size": "35",
    "sku": "FAL-840627987"}
`
