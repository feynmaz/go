package controllers

import (
	"net/http"

	"github.com/feynmaz/go/gorm/initializers"
	"github.com/feynmaz/go/gorm/models"
	"github.com/labstack/echo/v4"
)

func CreatePost(ctx echo.Context) error {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// Create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.NoContent(http.StatusBadRequest)
		return nil
	}

	// Return it
	return ctx.JSON(http.StatusOK, map[string]any{
		"post": post,
	})
}

func GetPosts(ctx echo.Context) error {
	// Get posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		ctx.NoContent(http.StatusBadRequest)
		return nil
	}

	// respond
	return ctx.JSON(200, map[string]any{
		"posts": posts,
	})
}

func GetPost(ctx echo.Context) error {
	// Get id off url
	id := ctx.Param("id")

	// Get post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		ctx.NoContent(http.StatusBadRequest)
		return nil
	}

	// respond
	return ctx.JSON(200, map[string]any{
		"id":   id,
		"post": post,
	})
}

func UpdatePost(ctx echo.Context) error {
	// Get id off url
	id := ctx.Param("id")

	// Get the data off request body
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)

	// Find post to update
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		ctx.NoContent(http.StatusBadRequest)
		return nil
	}

	// Update it
	result = initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		},
	)
	if result.Error != nil {
		ctx.NoContent(http.StatusBadRequest)
		return nil
	}

	// respond
	return ctx.JSON(200, map[string]any{
		"post": post,
	})
}

func DeletePost(ctx echo.Context) error {
	// Get id off url
	id := ctx.Param("id")

	// Delete post
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		ctx.NoContent(http.StatusBadRequest)
		return nil
	}

	// Respond
	return ctx.JSON(200, map[string]any{
		"deleted": id,
	})
}
