package http

import (
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/time"
)

var (
	DefaultConfig = Config{
		Address:      "[::]:8080",
		BasePath:     "/",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Middlewares:  &DefaultMiddlewareConfig,
		Swagger:      (func(b bool) *bool { return &b }(true)),
		SwaggerCORS:  (func(b bool) *bool { return &b }(true)),
	}

	DefaultMiddlewareConfig = MiddlewareConfig{
		// XXX: order is important
		Enable: []string{"auth", "logger", "requestid", "log", "pagination"},

		Auth:       &DefaultAuthMiddlewareConfig,
		Limit:      &DefaultLimitMiddlewareConfig,
		Log:        &DefaultLogMiddlewareConfig,
		Logger:     &DefaultLoggerMiddlewareConfig,
		RequestId:  &DefaultRequestIdMiddlewareConfig,
		Pagination: &DefaultPaginationMiddlewareConfig,
	}
)

//

type Config struct {
	Address      string
	BasePath     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Middlewares  *MiddlewareConfig
	Swagger      *bool
	SwaggerCORS  *bool
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Address == "":
			c.Address = DefaultConfig.Address
		case c.BasePath == "":
			c.BasePath = DefaultConfig.BasePath
		case c.ReadTimeout == time.Duration(0):
			c.ReadTimeout = DefaultConfig.ReadTimeout
		case c.WriteTimeout == time.Duration(0):
			c.WriteTimeout = DefaultConfig.WriteTimeout
		case c.Middlewares == nil:
			c.Middlewares = DefaultConfig.Middlewares
		case c.Swagger == nil:
			c.Swagger = DefaultConfig.Swagger
		case c.SwaggerCORS == nil:
			c.SwaggerCORS = DefaultConfig.SwaggerCORS
		default:
			break loop
		}
	}
	c.Middlewares.SetDefaults()
}

func (c Config) Validate() error {
	if c.Address == "" {
		return errors.New("address should not be empty")
	}
	if len(c.BasePath) == 0 || c.BasePath[0] != '/' {
		return errors.New("base path should begin with /")
	}
	if c.ReadTimeout == time.Duration(0) {
		return errors.New("read timeout should not be zero")
	}
	if c.WriteTimeout == time.Duration(0) {
		return errors.New("write timeout should not be zero")
	}
	if c.Middlewares == nil {
		return errors.New("middlewares configuration should be defined")
	}
	if c.Swagger == nil { // only if it is set to nil in config
		return errors.New("you should decide whether or not to use swagger")
	}
	if c.SwaggerCORS == nil { // only if it is set to nil in config
		return errors.New("you should decide whether or not to enable CORS for swagger spec")
	}

	err := c.Middlewares.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate middlewares configuration")
	}
	return nil
}

//

type MiddlewareConfig struct {
	Enable []string

	Auth       *AuthMiddlewareConfig
	Limit      *LimitMiddlewareConfig
	Log        *LogMiddlewareConfig
	Logger     *LoggerMiddlewareConfig
	RequestId  *RequestIdMiddlewareConfig
	Pagination *PaginationMiddlewareConfig
}

func (c *MiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		case c.Auth == nil:
			c.Auth = DefaultMiddlewareConfig.Auth
		case c.Limit == nil:
			c.Limit = DefaultMiddlewareConfig.Limit
		case c.Log == nil:
			c.Log = DefaultMiddlewareConfig.Log
		case c.Logger == nil:
			c.Logger = DefaultMiddlewareConfig.Logger
		case c.RequestId == nil:
			c.RequestId = DefaultMiddlewareConfig.RequestId
		case c.Pagination == nil:
			c.Pagination = DefaultMiddlewareConfig.Pagination
		default:
			break loop
		}
	}
	c.Auth.SetDefaults()
	c.Limit.SetDefaults()
	c.Log.SetDefaults()
	c.Logger.SetDefaults()
	c.RequestId.SetDefaults()
	c.Pagination.SetDefaults()
}

func (c MiddlewareConfig) Validate() error {
	if c.Auth == nil {
		return errors.New("auth middleware configuration should be defined")
	}
	if c.Limit == nil {
		return errors.New("limit middleware configuration should be defined")
	}
	if c.Log == nil {
		return errors.New("log middleware configuration should be defined")
	}
	if c.RequestId == nil {
		return errors.New("requestid middleware configuration should be defined")
	}
	if c.Pagination == nil {
		return errors.New("pagination middleware configuration should be defined")
	}

	err := c.Auth.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate auth middleware configuration")
	}
	err = c.Limit.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate limit middleware configuration")
	}
	err = c.Log.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate log middleware configuration")
	}
	err = c.Logger.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate logger middleware configuration")
	}
	err = c.RequestId.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate requestid middleware configuration")
	}
	err = c.Pagination.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate pagination middleware configuration")
	}

	return nil
}
