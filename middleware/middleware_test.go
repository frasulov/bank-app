package middleware

import (
	"BankApp/config"
	"BankApp/globals"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func addAuthorization(t *testing.T, request *http.Request, authorizationTypeBearer, username string, duration time.Duration) {
	token, err := globals.TokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationTypeBearer, token)
	request.Header.Set(config.Configuration.AuthorizationHeaderKey, authorizationHeader)
}

func TestProtectMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupHeader   func(t *testing.T, request *http.Request)
		checkResponse func(t *testing.T, response *http.Response)
	}{
		{
			name: "OK",
			setupHeader: func(t *testing.T, request *http.Request) {
				addAuthorization(t, request, "bearer", "frasulov", time.Minute)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, fiber.StatusOK, response.StatusCode)
			},
		},
		{
			name: "NoAuthorizationToken",
			setupHeader: func(t *testing.T, request *http.Request) {
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, fiber.StatusUnauthorized, response.StatusCode)
			},
		},
		{
			name: "UnsupportedAuthorizationType",
			setupHeader: func(t *testing.T, request *http.Request) {
				addAuthorization(t, request, "unsupported", "frasulov", time.Minute)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, fiber.StatusUnauthorized, response.StatusCode)
			},
		},
		{
			name: "InvalidAuthorizationHeader",
			setupHeader: func(t *testing.T, request *http.Request) {
				addAuthorization(t, request, "", "frasulov", time.Minute)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, fiber.StatusUnauthorized, response.StatusCode)
			},
		},
		{
			name: "AuthorizationTokenExpired",
			setupHeader: func(t *testing.T, request *http.Request) {
				addAuthorization(t, request, "bearer", "frasulov", -time.Minute)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, fiber.StatusUnauthorized, response.StatusCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			err := globals.Inject()
			require.NoError(t, err)
			app.Get("/test-middleware", Protect, func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{})
			})

			req := httptest.NewRequest("GET", "/test-middleware", nil)
			tc.setupHeader(t, req)
			resp, err := app.Test(req)
			require.NoError(t, err)
			tc.checkResponse(t, resp)
		})
	}
}
