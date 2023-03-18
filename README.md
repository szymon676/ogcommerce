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