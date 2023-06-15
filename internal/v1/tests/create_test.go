package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
	mock_company "github.com/tjarkmeyer/companies/companies/internal/v1/repositories/mocks"
)

func Test_CreateSuccess(t *testing.T) {
	companyRepoMock := mock_company.NewMockICompaniesRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().Create(&CreateCompanyRepo).Times(1).Return(nil)
	server := prepareServer(companyRepoMock)
	defer server.Close()

	reqBytes, err := json.Marshal(CreateCompanyRepo)
	assert.NoError(t, err)
	req := bytes.NewReader([]byte(reqBytes))

	resp, err := http.Post(server.URL+"/v1/companies", jsonContentType, req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusCreated)
}

func Test_CreateBadRequest(t *testing.T) {
	companyRepoMock := mock_company.NewMockICompaniesRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().Create(&CreateCompanyRepo).Times(1).Return(repositories.ErrAlreadyExist)
	server := prepareServer(companyRepoMock)
	defer server.Close()

	reqBytes, err := json.Marshal(CreateCompanyRepo)
	assert.NoError(t, err)
	req := bytes.NewReader([]byte(reqBytes))

	resp, err := http.Post(server.URL+"/v1/companies", jsonContentType, req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusConflict)
}
