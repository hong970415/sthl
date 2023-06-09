package config

import (
	"os"
	"sthl/constants"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	logger             *zap.Logger
	nodeEnv            string
	port               int
	dbDomain           string
	dbPort             int
	dbUser             string
	dbPw               bool
	jwtsecret          bool
	allowOrigin        string
	awsAccessKeyId     bool
	awsSecretAccessKey bool
	awsRegion          string
	s3Path             string
}

// new config by env, return false if not found
func NewConfig(logger *zap.Logger) (*Config, bool) {
	err := godotenv.Load()
	if err != nil {
		logger.Info("fail to godotenv.Load()", zap.Error(err))
	}

	// NODE_ENV
	nodeEnv, exist := os.LookupEnv("NODE_ENV")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "NODE_ENV not set"))
		return nil, false
	}

	// PORT
	_port, exist := os.LookupEnv("PORT")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "PORT not set"))
		return nil, false
	}
	port, err := strconv.Atoi(_port)
	if err != nil {
		logger.Info("fail to NewConfig", zap.String("err", "_port cannot convert to int"))
		return nil, false
	}

	// DB_DOMAIN
	dbDomain, exist := os.LookupEnv("DB_DOMAIN")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "DB_DOMAIN not set"))
		return nil, false
	}
	// DB_USER
	dbUser, exist := os.LookupEnv("DB_USER")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "DB_USER not set"))
		return nil, false
	}
	// DB_PASSWORD
	_, exist = os.LookupEnv("DB_PASSWORD")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "DB_PASSWORD not set"))
		return nil, false
	}
	// DB_PORT
	_dbPort, exist := os.LookupEnv("DB_PORT")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "DB_PORT not set"))
		return nil, false
	}
	dbPort, err := strconv.Atoi(_dbPort)
	if err != nil {
		logger.Info("fail to NewConfig", zap.String("err", "_dbPort cannot convert to int"))
		return nil, false
	}

	// JWT_SECRET
	_, exist = os.LookupEnv("JWT_SECRET")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "JWT_SECRET not set"))
		return nil, false
	}

	// ALLOW_ORIGIN
	allowOrigin, exist := os.LookupEnv("ALLOW_ORIGIN")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "ALLOW_ORIGIN not set"))
		return nil, false
	}

	// AWS_ACCESS_KEY_ID
	_, exist = os.LookupEnv("AWS_ACCESS_KEY_ID")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "AWS_ACCESS_KEY_ID not set"))
		return nil, false
	}
	// AWS_SECRET_ACCESS_KEY
	_, exist = os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "AWS_SECRET_ACCESS_KEY not set"))
		return nil, false
	}
	// AWS_REGION
	awsRegion, exist := os.LookupEnv("AWS_REGION")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "AWS_REGION not set"))
		return nil, false
	}

	// S3_PATH
	s3Path, exist := os.LookupEnv("S3_PATH")
	if !exist {
		logger.Info("fail to NewConfig", zap.String("err", "S3_PATH not set"))
		return nil, false
	}

	val := &Config{
		logger:             logger,
		nodeEnv:            nodeEnv,
		port:               port,
		dbDomain:           dbDomain,
		dbUser:             dbUser,
		dbPw:               true,
		dbPort:             dbPort,
		jwtsecret:          true,
		allowOrigin:        allowOrigin,
		awsAccessKeyId:     true,
		awsSecretAccessKey: true,
		awsRegion:          awsRegion,
		s3Path:             s3Path,
	}
	val.Print()
	return val, true
}
func (c *Config) Print() {
	c.logger.Info("Success to NewConfig",
		zap.String("NODE_ENV", c.nodeEnv),
		zap.Int("PORT", c.port),
		zap.String("DB_DOMAIN", c.dbDomain),
		zap.String("DB_USER", c.dbUser),
		zap.Bool("DB_PASSWORD", c.dbPw),
		zap.Int("DB_PORT", c.dbPort),
		zap.Bool("JWT_SECRET", c.jwtsecret),
		zap.String("ALLOW_ORIGIN", c.allowOrigin),
		zap.Bool("AWS_ACCESS_KEY_ID", c.awsAccessKeyId),
		zap.Bool("AWS_SECRET_ACCESS_KEY", c.awsSecretAccessKey),
		zap.String("AWS_REGION", c.awsRegion),
		zap.String("S3_PATH", c.s3Path),
	)
}

func (c *Config) GetNodeEnv() string {
	return c.nodeEnv
}
func (c *Config) GetPort() int {
	return c.port
}
func (c *Config) GetDbDomain() string {
	return c.dbDomain
}
func (c *Config) GetDbUser() string {
	return c.dbUser
}
func (c *Config) GetDbPw() string {
	dbPw, exist := os.LookupEnv("DB_PASSWORD")
	if !exist {
		c.logger.Error("fail to GetDB_PASSWORD")
		return ""
	}
	return dbPw
}
func (c *Config) GetDbPort() int {
	return c.dbPort
}
func (c *Config) SetDbPort(value int) {
	c.dbPort = value
}
func (c *Config) GetJwtSecret() (string, error) {
	secret, exist := os.LookupEnv("JWT_SECRET")
	if !exist {
		c.logger.Error("fail to GetJwtSecret")
		return "", constants.ErrNotFound
	}
	return secret, nil
}
func (c *Config) GetAllowOrigin() string {
	return c.allowOrigin
}
func (c *Config) GetS3Path() string {
	return c.s3Path
}
func (c *Config) GetAwsAccessKeyId() (string, error) {
	keyId, exist := os.LookupEnv("AWS_ACCESS_KEY_ID")
	if !exist {
		c.logger.Error("fail to GetAwsAccessKeyId")
		return "", constants.ErrNotFound
	}
	return keyId, nil
}
func (c *Config) GetAwsSecretAccessKey() (string, error) {
	keyId, exist := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !exist {
		c.logger.Error("fail to GetAwsSecretAccessKey")
		return "", constants.ErrNotFound
	}
	return keyId, nil
}
func (c *Config) GetAwsRegion() string {
	return c.awsRegion
}
