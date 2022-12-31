package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Config {
	appName := os.Getenv("APP_NAME")
	appVersion := os.Getenv("APP_VERSION")
	serverAllowedMethods := strings.Split(os.Getenv("SERVER_ALLOWED_METHODS"), ",")
	serverReadTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		AppName:        fmt.Sprintf("%s v%s", appName, appVersion),
		RequestMethods: serverAllowedMethods,
		ReadTimeout:    time.Second * time.Duration(serverReadTimeout),
	}
}
