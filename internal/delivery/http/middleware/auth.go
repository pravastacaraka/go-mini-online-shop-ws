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
		userID, _ := strconv.Atoi(ctx.Get("X-User-ID", "0"))

		request := &domain.AuthUserRequest{
			ID:    uint64(userID),
			Token: strings.Split(ctx.Get("Authorization", "NOT_FOUND"), " ")[1],
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
