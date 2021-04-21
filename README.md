# API Documentation

## Authentication

All API must use this authentication

Request :
- Header :
    - X-Api-Key : "your secret api key"

## Create User

Request :
- Method : POST
- Endpoint : `/api/users`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id" : "int",
    "name" : "string",
    "email" : "string,unique",
    "password" : "string",
    "address" : "string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "int",
        "name" : "string",
        "email" : "string,unique",
        "password" : "string, hash encode",
        "address" : "string",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```

## Get User Detail

Request :
- Method : GET
- Endpoint : `/api/users/{id_user}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "integer, unique",
        "email" : "string, unique",
        "name" : "string",
        "password" : "string, hash encode",
        "address" : "string",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```

## Update User

Request :
- Method : PUT
- Endpoint : `/api/users/{id_user}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id" : "int",
    "name" : "string",
    "email" : "string,unique",
    "password" : "string",
    "address" : "string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "int",
        "name" : "string",
        "email" : "string,unique",
        "password" : "string, hash encode",
        "address" : "string",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```

## List User

Request :
- Method : GET
- Endpoint : `/api/users`
- Header :
    - Accept: application/json
- Query Param :
    - size : number,
    - page : number

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "id" : "integer, unique",
            "name" : "string",
        },
        {
            "id" : "integer, unique",
            "name" : "string",
        },    ]
}
```

## Delete User

Request :
- Method : DELETE
- Endpoint : `/api/users/{id_user}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string"
}
```

## Create Product

Request :
- Method : POST
- Endpoint : `/api/products`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id" : "string, unique",
    "category_id" : "int,Foreign Key Category Table",
    "name" : "string",
    "description" : "string",
    "price" : "long",
    "stock" : "integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "string, unique",
        "category_id" : "int,Foreign Key Category Table",
        "name" : "string",
        "description" : "string",
        "price" : "long",
        "stock" : "integer",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```

## Get Product Detail

Request :
- Method : GET
- Endpoint : `/api/products/{id_product}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "string, unique",
        "category_id" : "int,Foreign Key Category Table",
        "name" : "string",
        "description" : "string",
        "price" : "long",
        "stock" : "integer",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```

## Update Product

Request :
- Method : PUT
- Endpoint : `/api/products/{id_product}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "category_id" : "int,Foreign Key Category Table",
    "name" : "string",
    "description" : "string",
    "price" : "long",
    "stock" : "integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "string, unique",
        "category_id" : "int,Foreign Key Category Table",
        "name" : "string",
        "description" : "string",
        "price" : "long",
        "stock" : "integer",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```

## List Product

Request :
- Method : GET
- Endpoint : `/api/products`
- Header :
    - Accept: application/json
- Query Param :
    - category : id_category,
    - size : number,
    - page : number

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "id" : "string, unique",
            "category_id" : "int,Foreign Key Category Table",
            "name" : "string",
            "description" : "string",
            "price" : "long",
            "stock" : "integer",
            "createdAt" : "date",
            "updatedAt" : "date"
        },
        {
            "id" : "string, unique",
            "category_id" : "int,Foreign Key Category Table",
            "name" : "string",
            "description" : "string",
            "price" : "long",
            "stock" : "integer",
            "createdAt" : "date",
            "updatedAt" : "date"
         }
    ]
}
```

## Delete Product

Request :
- Method : DELETE
- Endpoint : `/api/products/{id_product}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string"
}
```

## Create Category

Request :
- Method : POST
- Endpoint : `/api/category`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id" : "int,unique",
    "name" : "string",
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "int,unique",
        "name" : "string",
        "createdAt" : "date",
        "updatedAt" : "date"
     }
}
```
## List Category

Request :
- Method : GET
- Endpoint : `/api/category`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "id" : "integer, unique",
            "name" : "string",
            "createdAt" : "date",
            "updatedAt" : "date"
        },
        {
            "id" : "integer, unique",
            "name" : "string",
            "createdAt" : "date",
            "updatedAt" : "date"
         }
    ]
}
```

## Create Cart || Transaction || Send || Delivered

Request :
- Method : POST
- Endpoint : `/api/transactions`
- Query Param :
    - status : enum(cart,transaction,send,delivered),

- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id"        : "integer",
    "id_user"   : "integer,Foreign Key User Table",
    "date"      : "datetime",
    "status"    : "enum(cart,transaction,send,delivered)",
    "data_product"      : [
        {
            "id_product"    : "integer,Foreign Key Product Table",
            "quantity"      : "integer"
        },
        {
            "id_product"    : "integer,Foreign Key Product Table",
            "quantity"      : "integer"
        },
    ]
}
```
Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id"        : "integer",
        "id_user"   : "integer,Foreign Key User Table",
        "date"      : "datetime",
        "status"    : "enum(cart,transaction,send,delivered)",
        "data_product"      : [
            {
                "id_product"    : "integer,Foreign Key Product Table",
                "quantity"      : "integer"
            },
            {
                "id_product"    : "integer,Foreign Key Product Table",
                "quantity"      : "integer"
            },
        ]
    }
}
```

## List Cart || Transaction || Send || Delivered

Request :
- Method : GET
- Endpoint : `/api/transaction`
- Query Param :
    - status : enum(cart,transaction,send,delivered),
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id"        : "integer",
        "id_user"   : "integer,Foreign Key User Table",
        "date"      : "datetime",
        "status"    : "enum(cart,transaction,send,delivered)",
        "data_product"      : [
            {
                "id_product"    : "integer,Foreign Key Product Table",
                "quantity"      : "integer"
            },
            {
                "id_product"    : "integer,Foreign Key Product Table",
                "quantity"      : "integer"
            },
        ]
    }
}
```