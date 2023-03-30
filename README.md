# Go/MySQL API boilerplate.

This is an example of how to build a simple API using Go and MySQL.

Inside /db directory, there is a .sql file that could be used to create a database locally.
- You will need MySQL installed on your machine or a running container.

Before starting, you should rename the file `envexample` to `.env` and configure its content.

Then you may run `make server-run` to start the server.

You could test the API using Insomnia or any other client, or using the following `curl` commands on any terminal.

```
curl --request OPTIONS 
  --url http://localhost:8080/echo

curl --request GET \
  --url http://localhost:8080/echo \
  --header 'Content-Type: application/json'

curl --request POST \
  --url http://localhost:8080/echo \
  --header 'Content-Type: application/json' \
  --data Hey

curl --request GET \
  --url http://localhost:8080/testdb

curl --request GET \
  --url http://localhost:8080/testsdb
```
