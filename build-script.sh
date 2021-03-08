#! /bin/sh

export MYSQL_ROOT_PASSWORD=123
export MYSQL_DATABASE=remoteserver
export MYSQL_USER=rsuser
export MYSQL_PASSWORD=123456

change_source_code=$(git status  --porcelain main/ | wc -l)
if [ $change_source_code -gt 0 ]
then
    go build -o remote-server ./main/main.go
    docker build -t thanhgit/back-end-remote-server:latest .
else
    echo "No change source code"
fi

docker-compose up