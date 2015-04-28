Golang Web Application with Mongo connection
# Build Image
docker build -t daocloud/go-mongo .


## Below Mongo Connection Env and their default value
* MONGODB_PORT_27017_TCP_ADDR=localhost
* MONGODB_PORT_27017_TCP_PORT=27017
* MONGODB_USERNAME /* leave empty by default */
* MONGODB_PASSWORD /* leave empty by default */
* MONGODB_INSTANCE_NAME=test

# Run Container
docker run --link your_mongo:mongodb -d -p 80:80 daocloud/go-mongo

# That's it
