package app

import (
	"encoding/json"
	"errors"
	"strconv"
	"web-endterm/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) CreatePost(c *fiber.Ctx) error {
	post := new(model.Post)

	req := c.Request().Body()
	err := json.Unmarshal(req, &post)
	if err != nil {
		c.Status(400).SendString("error while parsing body: " + err.Error())
	}

	err = s.Service.CreatePost(post)
	if err != nil {
		c.Status(500).SendString("error while storing post: " + err.Error())
	}

	return c.SendString("success!")
}

func (s *Server) GetPostById(c *fiber.Ctx) error {
	idStr := c.Params("id")

	if idStr == "" {
		return c.Status(400).SendString("id not provided")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).SendString("id must be int")
	}

	post, err := s.Service.GetPostById(id)
	if err != nil {
		return c.Status(400).SendString("error getting post: " + err.Error())
	}

	return c.JSON(post)
}

func (s *Server) GetAllPosts(c *fiber.Ctx) error {
	posts, err := s.Service.GetAllPosts()
	if err != nil {
		return c.Status(400).SendString("error getting posts: " + err.Error())
	}

	return c.JSON(posts)
}

func (s *Server) PatchPost(c *fiber.Ctx) error {
	id, err := getId(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("error getting id: " + err.Error())
	}
	
	post := new(model.Post)

	req := c.Request().Body()
	err = json.Unmarshal(req, &post)
	if err != nil {
		c.Status(400).SendString("error while parsing body: " + err.Error())
	}

	res, err := s.Service.PatchPost(id, post)
	if err != nil{
		return c.Status(400).SendString("error patching post: " + err.Error())
	}

	return c.JSON(res)
}

func (s *Server) DeletePost(c * fiber.Ctx) error {
	id, err := getId(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("error getting id: " + err.Error())
	}

	err = s.Service.DeletePost(id)
	if err != nil {
		c.Status(400).SendString("error deleting post: "  + err.Error())
	}

	return c.SendString("success")
}

func getId(str string) (int, error) {
	if str == "" {
		return -1, errors.New("id not provided")
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}

	return id, nil
}