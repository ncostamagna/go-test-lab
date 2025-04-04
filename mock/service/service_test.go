// service_test.go
package service_test

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/ncostamagna/go-test-lab/mock/service"
)

type MockAPIClient struct {
    mock.Mock
}

func (m *MockAPIClient) FetchData(id string) (string, error) {
    args := m.Called(id)
    return args.String(0), args.Error(1)
}

func TestGetProcessedData_Success(t *testing.T) {
    mockClient := new(MockAPIClient)
    mockClient.On("FetchData", "123").Return("raw-data", nil)

    svc := service.NewDataService(mockClient)
    result, err := svc.GetProcessedData("123")

    assert.NoError(t, err)
    assert.Equal(t, "Processed: raw-data", result)
    mockClient.AssertExpectations(t)
}

func TestGetProcessedData_Error(t *testing.T) {
    mockClient := new(MockAPIClient)
    mockClient.On("FetchData", "123").Return("", errors.New("not found"))

    svc := service.NewDataService(mockClient)
    result, err := svc.GetProcessedData("123")

    assert.Error(t, err)
    assert.Equal(t, "", result)
    mockClient.AssertExpectations(t)
}
