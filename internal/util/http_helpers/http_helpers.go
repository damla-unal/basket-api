package http_helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func GetPositiveIntegerQueryParameter(c *gin.Context, queryParamName string) (*int, error) {
	value, err := GetRequiredQueryParameter(c, queryParamName)
	if err != nil {
		return nil, err
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	if intValue <= 0 {
		return nil, errors.New("Expected a positive integer for parameter " + queryParamName + " but value is <= 0")
	}
	return &intValue, nil
}

func GetRequiredQueryParameter(c *gin.Context, key string) (string, error) {
	value := c.Query(key)
	value = strings.TrimSpace(value)
	if value == "" {
		return "", errors.New(key + " is required and missing from the query parameters")
	}
	return value, nil
}

func GetRequiredPathVariable(c *gin.Context, key string) (*int, error) {
	value := c.Params.ByName(key)
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, errors.New(key + " is required and missing from the path variables")
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}
