package db

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Server struct {
	gorm.Model
	Name string
	Ip string
	Port string
	Password string
}

var serverType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Server",
		Fields:      graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Ip": &graphql.Field{
				Type: graphql.String,
			},
			"Port": &graphql.Field{
				Type:              graphql.String,
			},
			"Password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
	)

const servers_table_name  = "servers"

func CreateServer(server Server) {
	println(server.Port)
	if GetInstance().NewRecord(server) {
		GetInstance().Create(&server)
		if GetInstance().NewRecord(server) {
			println("not success")
		} else {
			println("Success")
		}
	}
}

func UpdateServer(server Server) {
	instance = GetInstance()
	instance.Model(&server).Updates(server)
}

func UpdateDeleteServer(_id uint8) {
	instance = GetInstance()
	instance.Where("id = ?", _id).Delete(&Server{})
}

func DeleteServer(_id int) {
	instance = GetInstance()
	instance.Delete(Server{}, "id = ?", _id)
}

func CheckExistServerById(_id int) bool {
	instance = GetInstance()
	var server Server
	server.ID = uint(_id)
	if instance.First(&server).Row() == nil {
		return false
	} else {
		return true
	}

}

func GetAllServers() []Server {
	instance = GetInstance();
	if instance != nil {
		var results []Server
		instance.Find(&results)
		return results
	} else {
		return nil
	}
}