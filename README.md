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
This route is used to create a new product. It requires authentication using JWT (JSON Web Token) middleware. The request should contain the necessary data for creating a product in the request body. Upon successful creation, the server will return the newly created product object.

## Get Products
### Method: GET

#### /products
This route is used to retrieve all products. No authentication is required for this route. The server will return an array of product objects containing the details of all products available in the system.

## Get Product by id
### Method: GET

#### /products/{id}
This route is used to retrieve a single product by its ID. No authentication is required for this route. The server will return the product object containing the details of the requested product.

## Delete Product by id
### Method: DELETE

### /products/{id}
This route is used to delete a product by its ID. Authentication using JWT middleware is required to access this route. Upon successful deletion, the server will return a response indicating that the product has been deleted.

## Update Product
### Method: PUT

#### /products/{id}
This route is used to update a product by its ID. Authentication using JWT middleware is required to access this route. The request should contain the updated data for the product in the request body. Upon successful update, the server will return the updated product object.

## User Login
### Method: POST

#### /login
This route is used to authenticate a user and generate a JWT token for further authentication. The request should contain the user's credentials in the request body. Upon successful authentication, the server will return a JWT token that can be used for authentication in subsequent requests to routes that require authentication.