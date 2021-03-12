package main

import (
	"strconv"
	"thanhgit.com/backend-end-remote-server/main/db"
	_ "thanhgit.com/backend-end-remote-server/main/db"
	_ "thanhgit.com/backend-end-remote-server/main/docs"
	"thanhgit.com/backend-end-remote-server/main/dto"
	_ "thanhgit.com/backend-end-remote-server/main/dto"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func connectDB() {
	instance := db.GetInstance()

	if instance != nil {
		println("Connect MySQL OK")

	} else {
		println("Fail to connect")
	}
}

func business() {

}

func setupGlobalMiddleware(c *gin.Context) *gin.Context {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	return c
}

func main() {

	connectDB()

	business()

	r := gin.New()

	url := ginSwagger.URL("http://localhost:7000/swagger/doc.json") // The url endpoint
	// API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Servers API
	r.GET("/api/servers", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		var servers []db.Server = db.GetAllServers()
		c.JSON(200, gin.H{
			"servers": servers,
			"numbers": len(servers),
		})
	})

	r.POST("/api/servers", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)

		var server db.Server
		if c.BindJSON(&server) != nil {
			c.JSON(500, gin.H{
				"Status": "Error for posting server",
				"Code":   500,
			})
		} else {
			db.CreateServer(server)
			c.JSON(200, gin.H{
				"Status": "Success",
				"Code":   200,
			})
		}
	})

	r.POST("/api/servers/:id_update", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		var server db.Server
		if c.BindJSON(&server) != nil {
			c.JSON(500, gin.H{
				"Status": "Erorr for updating server",
				"Code":   500,
			})
		} else {
			if db.CheckExistServerById(int(server.ID)) {
				db.UpdateServer(server)
				c.JSON(200, gin.H{
					"Status": "Sucess",
					"Code":   200,
				})
			} else {
				c.JSON(500, gin.H{
					"Status": "Not found Server",
					"Code":   500,
				})
			}

		}

	})

	r.DELETE("/api/servers/:id_delete", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)

		_id := c.Param("id_delete")

		id, err := strconv.Atoi(_id)

		if err == nil {
			db.DeleteServer(id)
			c.JSON(200, gin.H{
				"Status": "Success",
				"Code":   200,
			})
		} else {
			c.JSON(500, gin.H{
				"Status": err.Error(),
				"Code":   500,
			})
		}
	})

	// Websites API
	r.GET("/api/websites", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)

		var websites []db.Website = db.GetAllWebsites()
		c.JSON(200, gin.H{
			"websites": websites,
			"numbers":  len(websites),
		})

	})

	r.POST("/api/websites", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		var website db.Website
		if c.BindJSON(&website) != nil {
			c.JSON(500, gin.H{
				"Status": "Error for posting website instance",
				"Code":   500,
			})
		} else {
			db.CreateWebsite(website)
			c.JSON(200, gin.H{
				"Status": "Success",
				"Code":   200,
			})
		}
	})

	r.POST("/api/websites/:id_update", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		var website db.Website
		if c.BindJSON(&website) != nil {
			c.JSON(500, gin.H{
				"Status": "Error for updating website",
				"Code":   500,
			})
		} else {
			if db.CheckExistWebsiteById(int(website.ID)) {
				db.UpdateWebsite(website)
				c.JSON(200, gin.H{
					"Status": "Sucess",
					"Code":   200,
				})
			} else {
				c.JSON(500, gin.H{
					"Status": "Not found Website",
					"Code":   500,
				})
			}
		}
	})

	r.DELETE("/api/websites/:id_delete", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		_id := c.Param("id_delete")
		id, err := strconv.Atoi(_id)
		if err == nil {
			db.DeleteWebsite(id)
			c.JSON(200, gin.H{
				"Status": "Success",
				"Code":   200,
			})
		} else {
			c.JSON(500, gin.H{
				"Status": err.Error(),
				"Code":   500,
			})
		}
	})

	// Get dashboard
	r.POST("/api/dashboards", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		var password dto.PasswordDTO
		if c.BindJSON(&password) != nil {
			websites := db.GetWebsitesByTypeAndPassword("dashboard", "")
			c.JSON(200, gin.H{
				"Websites": websites,
				"Numbers":  len(websites),
				"Type":     "dashboard",
			})
		} else {
			websites := db.GetWebsitesByTypeAndPassword("dashboard", password.Password)
			c.JSON(200, gin.H{
				"Websites": websites,
				"Numbers":  len(websites),
				"Type":     "dashboard",
			})
		}

	})

	r.Run(":7000")
}
