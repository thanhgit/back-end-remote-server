## Remote-server
This is apart of remote server project. It's responsibility is to provide API for front-end project

## Remote server project (https://github.com/thanhgit/remote-server)

## Quickly
nohup go run main/main.go &

## File Structure

```
.
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
    -- vendor.json
```

## Database using mysql
```
$ docker run --name remote-server -e MYSQL_ROOT_PASSWORD=123456 -d mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

## Build go
```bash
go build -o remote-server ./main/main.go
```

## Build docker-compose

```bash
./build-script.sh
```

## [API](http:localhost:7000)
- GET    [/api/servers](localhost:7000/api/servers)              
- POST   /api/servers              
- POST   /api/servers/:id_update   
- DELETE /api/servers/:id_delete   
- GET    [/api/websites](localhost:7000/api)             
- POST   /api/websites             
- POST   /api/websites/:id_update  
- DELETE /api/websites/:id_delete  

## [Swager GUI](http://localhost:7000/swagger/index.html#)

### Functionality of application
* #### Show dashboard 
* #### CRUD for servers, can navigate to sshweb2 or CRT
* #### CRUD for dashboards


## Technical Support or Questions

If you have questions or need help integrating the product please "thanh29695@gmail.com" instead of opening an issue.

