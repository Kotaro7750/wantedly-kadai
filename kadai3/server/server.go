package main

import(
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"

	"server/controller"
)

func main()  {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	
	db, err := sql.Open("postgres", "host=db port=5432 user=kadai dbname=kadai password=kadai sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userctr := &controller.UserCtr{DB: db}

	router.Use(cors.New(config))
	router.GET("/", userctr.HelloWorld)
	router.GET("/users", userctr.GetUserAll)
	router.GET("/users/:id", userctr.GetUserByid)
	router.POST("/users", userctr.InsertUser)
	router.PUT("/users/:id", userctr.UpdateUser)
	router.DELETE("/users/:id", userctr.DeleteUser)
	router.Run(":8080")
}
