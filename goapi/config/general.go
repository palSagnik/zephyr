package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var SESSION_SECRET = os.Getenv("SESSION_SECRET")
var SESSION_EXPIRY = 72
var MAIL_LEN = 320
var PASS_LEN = 32
var USERNAME_LEN = 24