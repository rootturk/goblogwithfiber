package main

import(
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rootturk/goblogwithfiber/blog"
	"github.com/rootturk/goblogwithfiber/database"
	"fmt"
)

func initilazeRoutes(app *fiber.App){
	app.Get("/api/v1/blog",blog.GetBlogs)
	app.Get("/api/v1/blog/:id",blog.GetBlogById)
	app.Post("/api/v1/blog/:id",blog.NewBlog)
	app.Delete("/api/v1/blog/:id",blog.DeleteBlog)

	app.Get("/api/v1/article",blog.GetArticles)
	app.Get("/api/v1/article/:id",blog.GetArticleById)
	app.Post("/api/v1/article/:id",blog.NewArticle)
	app.Delete("/api/v1/article/:id",blog.DeleteArticle)	

	app.Get("/api/v1/user",blog.GetUsers)
	app.Get("/api/v1/user/:id",blog.GetUserById)
	app.Post("/api/v1/user/:id",blog.NewUser)
	app.Delete("/api/v1/user/:id",blog.DeleteUser)	
}

func initilazeDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "fiberblog.db")

	if err != nil {
		panic("Database connection Error.")
	}

	fmt.Println("Database Connection is succeeded")

	database.DBConn.AutoMigrate(&blog.Blog{})
}

func main(){
	app := fiber.New()

	initilazeDatabase()
	initilazeRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}