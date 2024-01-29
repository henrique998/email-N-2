package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_When_Endpoint_Returns_Internal_Error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerError_When_Endpoint_Returns_Domain_Error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func Test_HandlerError_When_Endpoint_Returns_Obj_And_Status(t *testing.T) {
	assert := assert.New(t)
	type bodyForTest struct {
		Id int
	}
	objExpected := bodyForTest{Id: 2}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	objReturned := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)

	assert.Equal(http.StatusCreated, res.Code)
	assert.Equal(objExpected, objReturned)
}
