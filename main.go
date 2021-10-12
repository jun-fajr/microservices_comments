package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Comment struct {
    Id uint `json:"id"`
    PostId uint `json:"post_id"`
    Text string `json:"text"`
}

func main() {
    db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/comments_ms"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    db.AutoMigrate(Comment{})

    app := fiber.New()

    app.Use(cors.New())

    app.Get("/api/posts/:id/comments", func(ctx *fiber.Ctx) error {
        var comments []Comment
        db.Find(&comments, "post_id = ?", ctx.Params("id"))
        return ctx.JSON(comments)
    })

    //app.Get("/", func(c *fiber.Ctx) error {
    //    return c.SendString("Hello, World 👋!")
    //})

    app.Post("/api/comments", func(ctx *fiber.Ctx) error {
        var comment Comment
        if err := ctx.BodyParser(&comment); err != nil{
            return err
        }
        db.Create(&comment)
        //ctx.BodyParser(&posts)
        //db.Find(&posts)
        return ctx.JSON(comment)
    })

    app.Listen(":8001")
}
