package http

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/rezamokaram/sample-auth/api/pb"
	"github.com/rezamokaram/sample-auth/api/service"
	appCtx "github.com/rezamokaram/sample-auth/pkg/context"

	"github.com/gofiber/fiber/v2"
)

func SendSignInOTP(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		phone := strings.TrimSpace(c.Query("phone"))

		if err := svc.SendSignInOTP(c.UserContext(), phone); err != nil {
			return err
		}

		return nil
	}
}

func SignUp(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {

		svc := svcGetter(c.UserContext())

		var req pb.UserSignUpRequest
		if err := c.BodyParser(&req); err != nil {

			appCtx.GetLogger(c.UserContext()).Error("bad request", "err", err.Error())
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignUp(c.UserContext(), &req)

		if err != nil {
			appCtx.GetLogger(c.UserContext()).Error("service error", "err", err.Error())
			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		appCtx.GetLogger(c.UserContext()).Error("successful")
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "successful",
			"data":    resp,
		})
	}
}

func SignIn(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.UserSignInRequest
		if err := c.BodyParser(&req); err != nil {
			appCtx.GetLogger(c.UserContext()).Error("bad request", "err", err.Error())
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignIn(c.UserContext(), &req)
		if err != nil {
			appCtx.GetLogger(c.UserContext()).Error("service error", "err", err.Error())
			if errors.Is(err, service.ErrUserNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			if errors.Is(err, service.ErrInvalidUserPassword) {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		appCtx.GetLogger(c.UserContext()).Error("successful")
		appCtx.GetLogger(c.UserContext()).Error("successful")
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "successful",
			"data":    resp,
		})
	}
}

func TestHandler(ctx *fiber.Ctx) error {
	logger := appCtx.GetLogger(ctx.UserContext())

	logger.Info("from test handler", "time", time.Now().Format(time.DateTime))

	return nil
}
