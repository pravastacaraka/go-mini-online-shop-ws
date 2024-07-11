package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/usecase"
)

func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &domain.AuthUserRequest{}

		if userID, _ := strconv.Atoi(ctx.Get("X-User-ID", "0")); userID != 0 {
			request.ID = uint64(userID)
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "User id must not 0")
		}

		if token := strings.Split(ctx.Get("Authorization", "NOT_FOUND"), " "); len(token) > 1 {
			request.Token = token[1]
		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "Authentication is not found")
		}

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			return err
		}
		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}

func GetAuthenticatedUserID(ctx *fiber.Ctx) *domain.AuthUserResponse {
	return ctx.Locals("auth").(*domain.AuthUserResponse)
}
