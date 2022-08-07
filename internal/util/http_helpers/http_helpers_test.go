package http_helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestGetRequiredQueryParameter(t *testing.T) {
	key := "key"
	value := "value"
	urlMock := url.URL{RawQuery: key + "=" + value}
	request := http.Request{
		URL: &urlMock,
	}
	context := gin.Context{
		Request: &request,
	}
	actual, err := GetRequiredQueryParameter(&context, key)
	require.NoError(t, err)
	require.Equal(t, value, actual)
}

func TestGetPositiveIntegerQueryParameter(t *testing.T) {
	key := "key"
	value := "1"
	urlMock := url.URL{RawQuery: key + "=" + value}
	request := http.Request{
		URL: &urlMock,
	}
	context := gin.Context{
		Request: &request,
	}
	expectedValue, err := strconv.Atoi(value)
	require.NoError(t, err)
	actual, err := GetPositiveIntegerQueryParameter(&context, key)
	require.NoError(t, err)
	require.Equal(t, expectedValue, *actual)
}

func TestGetRequiredUrlParameter(t *testing.T) {
	key := "key"
	value := "2"
	param := gin.Param{Key: key, Value: value}
	context := gin.Context{
		Params: gin.Params{param},
	}
	actual, err := GetRequiredPathVariable(&context, key)
	require.NoError(t, err)
	expectedValue, err := strconv.Atoi(value)
	require.NoError(t, err)
	require.Equal(t, expectedValue, *actual)
}

func TestGetRequiredUrlParameterMissingKey(t *testing.T) {
	context := gin.Context{}
	_, err := GetRequiredPathVariable(&context, "missing")
	require.Error(t, err)
}
