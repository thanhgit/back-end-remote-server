package db

import "github.com/jinzhu/gorm"

type Website struct {
	gorm.Model
	Url string
	Username string
	Password string
	Note string
}

const websites_table_name = "websites"

func CreateWebsite(website Website) {
	if GetInstance().NewRecord(website) {
		GetInstance().Create(&website)
		if GetInstance().NewRecord(website) {
			println("not success")
		} else {
			println("Success")
		}
	}
}

func UpdateWebsite(website Website) {
	instance = GetInstance()
	instance.Model(&website).Updates(website)
}

func UpdateDeleteWebsite(_id uint8) {
	instance = GetInstance()
	instance.Where("id = ?", _id).Delete(&Website{})
}

func DeleteWebsite(_id int) {
	instance = GetInstance()
	instance.Delete(Website{}, "id = ?", _id)
}

func CheckExistWebsiteById(_id int) bool {
	instance = GetInstance()
	var website Website
	website.ID = uint(_id)
	if instance.First(&website).Row() == nil {
		return false
	} else {
		return true
	}
}


func GetAllWebsites() []Website {
	instance = GetInstance();
	if instance != nil {
		var results []Website
		instance.Find(&results)
		return results
	} else {
		return nil
	}
}

func GetWebsitesByTypeAndPassword(_type string, _password string) (resutls []Website) {
	instance = GetInstance()
	if instance != nil {
		instance.Where("note = ?", _type).Where("password = ?", _password).Find(&resutls)
	} else {
		resutls = nil
	}
	return
}
