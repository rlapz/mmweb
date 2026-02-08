package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	APP_NAME = "mmweb"

	CLAIMS_CONTEXT_NAME = "claims_context"

	/* Default */
	DEF_LISTEN_HOST = "127.0.0.1"
	DEF_LISTEN_PORT = "8008"

	DEF_DB_PATH           = "./db.sqlite"
	DEF_DB_POOL_INIT_SIZE = 8

	DEF_LOGIN_EXP = time.Hour * 8
)

var (
	JWT_SIGN_METHOD = jwt.SigningMethodHS256
)

type Config struct {
	ListenHost      string
	ListenPort      string
	JwtSignMethod   *jwt.SigningMethodHMAC
	JwtSignatureKey []byte
	LoginExp        time.Duration

	DbPath         string
	DbPoolInitSize int
}

func Load() (*Config, error) {
	c := new(Config)

	c.ListenHost = os.Getenv("MMWEB_LISTEN_HOST")
	if c.ListenHost == "" {
		c.ListenHost = DEF_LISTEN_HOST
	}

	c.ListenPort = os.Getenv("MMWEB_LISTEN_PORT")
	if c.ListenPort == "" {
		c.ListenPort = DEF_LISTEN_PORT
	}

	c.JwtSignMethod = JWT_SIGN_METHOD

	signKey := os.Getenv("MMWEB_SECRET_KEY")
	if signKey == "" {
		return nil, errors.New("no secret key")
	}
	c.JwtSignatureKey = []byte(signKey)

	c.DbPath = os.Getenv("MMWEB_DB_PATH")
	if c.DbPath == "" {
		c.DbPath = DEF_DB_PATH
	}

	num, err := strconv.ParseInt(os.Getenv("MMWEB_DB_POOL_INIT_SIZE"), 10, 32)
	if err != nil {
		num = DEF_DB_POOL_INIT_SIZE
	}
	c.DbPoolInitSize = int(num)

	exp, err := strconv.ParseInt(os.Getenv("MMWEB_LOGIN_EXPIRE"), 10, 64)
	if err != nil {
		exp = int64(DEF_LOGIN_EXP)
	}
	c.LoginExp = time.Duration(exp)

	return c, nil
}
