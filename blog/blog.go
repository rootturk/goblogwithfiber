package blog

import(
		"github.com/jinzhu/gorm"
		"github.com/gofiber/fiber"
		_ "github.com/jinzhu/gorm/dialects/sqlite"
		"github.com/rootturk/goblogwithfiber/database"
		"time"
)


type Blog struct {
	gorm.Model
	Title string
	Description string
	CreatedBy int32
	CreatedOn time.Time
}

func GetBlogs(c *fiber.Ctx){
	db := database.DBConn

	var blogs []Blog

	db.Find(&blogs)

	c.JSON(blogs)
}

func GetBlogById(c *fiber.Ctx){

	id := c.Params("id")
	db := database.DBConn

	var blog Blog

	db.Find(&blog, id)

	c.JSON(blog)

}

func NewBlog(c *fiber.Ctx){
	db := database.DBConn

	blog := new(Blog)

	if err := c.BodyParser(blog); err !=nil{
		c.Status(503).Send(err)
		return
	}

	db.Create(&blog)
	c.JSON(blog)
}


func DeleteBlog(c *fiber.Ctx){
	id := c.Params("id")

	db := database.DBConn

	var blog Blog

	db.First(&blog, id)

	if blog.Title == "" {
		c.Status(500).Send("Blog not found with Id")
		return
	}

	db.Delete(&blog)

	c.Send("Blog deleted succeeded.")
}


type Article struct {
	gorm.Model
	Title string
	Description string
	MetaTags string
	CreatedBy int32
	CreatedOn time.Time
}

func GetArticles(c *fiber.Ctx){
	db := database.DBConn

	var Articles []Article

	db.Find(&Articles)

	c.JSON(Articles)
}

func GetArticleById(c *fiber.Ctx){

	id := c.Params("id")
	db := database.DBConn

	var Article Article

	db.Find(&Article, id)

	c.JSON(Article)

}

func NewArticle(c *fiber.Ctx){
	db := database.DBConn

	Article := new(Article)

	if err := c.BodyParser(Article); err !=nil{
		c.Status(503).Send(err)
		return
	}

	db.Create(&Article)
	c.JSON(Article)
}


func DeleteArticle(c *fiber.Ctx){
	id := c.Params("id")

	db := database.DBConn

	var Article Article

	db.First(&Article, id)

	if Article.Title == "" {
		c.Status(500).Send("Article not found with Id")
		return
	}

	db.Delete(&Article)

	c.Send("Article deleted succeeded.")
}


type User struct {
	gorm.Model
	Username string
	UserType int32
	Password string
	Email string
	BlogId int32
}

func GetUsers(c *fiber.Ctx){
	db := database.DBConn

	var Users []User

	db.Find(&Users)

	c.JSON(Users)
}

func GetUserById(c *fiber.Ctx){

	id := c.Params("id")
	db := database.DBConn

	var User User

	db.Find(&User, id)

	c.JSON(User)

}

func NewUser(c *fiber.Ctx){
	db := database.DBConn

	User := new(User)

	if err := c.BodyParser(User); err !=nil{
		c.Status(503).Send(err)
		return
	}

	db.Create(&User)
	c.JSON(User)
}


func DeleteUser(c *fiber.Ctx){
	id := c.Params("id")

	db := database.DBConn

	var User User

	db.First(&User, id)

	if User.Username == "" {
		c.Status(500).Send("User not found with Id")
		return
	}

	db.Delete(&User)

	c.Send("User deleted succeeded.")
}