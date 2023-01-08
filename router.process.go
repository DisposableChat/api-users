package main

import (
	"fmt"
	"errors"
	"context"
	"github.com/DisposableChat/api-core"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) (string, interface{}, error) {
	id := c.Query("id", "")
	if len(id) < 1 {
		return "", nil, errors.New(core.InvalidQueryError)
	}

	cmd := Redis.Client.Get(context.Background(), fmt.Sprintf("user:%s", id))
	if cmd.Err() != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	result, err := cmd.Result()
	if err != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	if result == "" {
		return "", nil, errors.New(core.UserNotFoundError)
	}

	var user core.User
	err = core.UnParseJSON(result, &user)
	if err != nil {
		return "", nil, errors.New(core.InternalServerError)
	}

	return core.Authorized, user, nil
}