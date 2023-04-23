# Installing 

## Go dependencies 
```
 go get go.mongodb.org/mongo-driver/mongo
 go get go.mongodb.org/mongo-driver/bson
```

# Installing mongodb with Docker
```
docker run --name some-mongo -p 27017:27017 -d mongo
```

# Auth dependencies

```
docker exec -it some-mongo mongo
```

```
db.users.insertOne({
    "id": "user123",
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "mypassword"
})
```

#### write this command in mongo console to create account that will be used to login and acces admin routes.

# API Routes

## Create Product
### Method: POST
#### /products

## Get Products
### Method: GET
#### /products

## Get Product by id
### Method: GET

#### /products/{id}
## Delete Product by id
### Method: DELETE
### /products/{id}

## Update Product
### Method: PUT
#### /products/{id}

## User Login
### Method: POST
#### /login