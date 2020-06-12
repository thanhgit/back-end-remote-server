## Back-end-remote-server
This is apart of remote server project. It's reposibility is to provide API for front-end project

## Quickly
nohup go run main/main.go &

## File Structure

```
.
|-- README.md
|-- go.mod
|-- go.sum
|-- main
|   |-- api
|   |   -- api.go
|   |-- db
|   |   |-- dbconnection.go
|   |   -- server.go
|   |-- docs
|   |   -- docs.go
|   |-- go.mod
|   |-- go.sum
|   |-- main.go
|   -- web
|       -- handler.go
-- vendor
    |-- github.com
    |   -- swaggo
    |       |-- files
    |       |-- gin-swagger
    |       |   |-- LICENSE
    |       |   |-- README.md
    |       |   |-- b0x.yml
    |       |   |-- go.mod
    |       |   |-- go.sum
    |       |   -- swagger.go
    |       -- swag
    |           |-- cmd
    |           |   `-- swag
    |           |       -- main.go
    |           -- license
    -- vendor.json
```

## Database
### Using mysql
### Build with docker
```
docker run --name remote-server -e MYSQL_ROOT_PASSWORD=123456 -d mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

### Functionality of application
#### Show dashboard 
#### CRUD for servers, can navigate to sshweb2 or CRT
#### CRUD for dashboards


## Technical Support or Questions

If you have questions or need help integrating the product please "thanh29695@gmail.com" instead of opening an issue.

