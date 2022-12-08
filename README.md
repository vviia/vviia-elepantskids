### vviia-elepantskids project -> building crud api stocks for hospital

API Information:

POST request from localhost:8080/api/newstock

```
{
    "name":"heart",
    "price":30,
    "company":"yogyakarta international hospital"
}
```
and  the result from postman is 
```
{
    "id": 1,
    "message": "Stock created successfully"
}
```
------
GET all request localhost:8080/api/stock


```[
    {
        "stockid": 1,
        "name": "heart",
        "price": 30,
        "company": "yogyakarta international hospital"
    }
]
```

GET by ID request localhost:8080/api/stock/1

```
{
    "stockid": 1,
    "name": "heart",
    "price": 30,
    "company": "yogyakarta international hospital"
}
```
----
PUT request localhost:8080/api/stock/1

insert this into JSON before request
```
{
    "name":"brain",
    "price":30,
    "company":"yogyakarta international hospital"
}
```

and this is the result :

```
{
    "id": 1,
    "message": "Stock updated successfully. Total rows/record affected 1"
}
```
------
DELETE request localhost:8080/api/deletestock/1

```
{
    "id": 1,
    "message": "Stock updated successfully. Total rows/record affected 1"
}
```

