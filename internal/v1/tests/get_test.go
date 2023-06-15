package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
	mock_company "github.com/tjarkmeyer/companies/companies/internal/v1/repositories/mocks"
)

func Test_GetSuccess(t *testing.T) {
	companyRepoMock := mock_company.NewMockICompaniesRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().GetByID(defaultCompanyID).Times(1).Return(&GetCompanyRepo, nil)
	server := prepareServer(companyRepoMock)
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/companies/" + defaultCompanyID)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	company, err := extractGetResponse(resp)

	assert.NoError(t, err)
	assert.Equal(t, &GetCompanyRepo, company)
}

func Test_GetNotFound(t *testing.T) {
	companyRepoMock := mock_company.NewMockICompaniesRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().GetByID(defaultCompanyID).Times(1).Return(nil, repositories.ErrNotFound)
	server := prepareServer(companyRepoMock)
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/companies/" + defaultCompanyID)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func extractGetResponse(r *http.Response) (res *models.Company, err error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	return res, json.Unmarshal(b, &res)
}
