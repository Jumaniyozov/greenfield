package translation_test

import (
	"encoding/json"
	"github.com/jumaniyozov/greenfield/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloClientSuite(t *testing.T) {
	suite.Run(t, new(HelloClientSuite))
}

type HelloClientSuite struct {
	suite.Suite
	mockServerService *MockService
	server            *httptest.Server
	underTest         translation.HelloClient
}
type MockService struct {
	mock.Mock
}

func (m *MockService) Translate(word, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *HelloClientSuite) SetupSuite() {

	suite.mockServerService = new(MockService)

	handler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		defer func() {
			err := r.Body.Close()
			if err != nil {
				http.Error(w, "error", 500)
			}
		}()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		word := m["word"].(string)
		language := m["language"].(string)
		resp, err := suite.mockServerService.Translate(word, language)
		if err != nil {
			http.Error(w, "error", 500)
		}
		if resp == "" {
			http.Error(w, "missing", 404)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, resp)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	suite.server = httptest.NewServer(mux)
	suite.underTest = translation.NewHelloClient(suite.server.URL)
}
func (suite *HelloClientSuite) TearDownSuite() {
	suite.server.Close()
}

//
//func (suite *HelloClientSuite) TestCall() {
//	// Arrange
//	suite.mockServerService.On("Translate", "foo", "bar").Return(`{"translation":"baz"}`, nil)
//	// Act
//	resp, err := suite.mockServerService.Translate("foo", "bar")
//	// Assert
//	suite.NoError(err)
//	suite.Equal(resp, "baz")
//}
//func (suite *HelloClientSuite) TestCall_APIError() {
//	// Arrange
//	suite.mockServerService.On("Translate", "foo", "bar").Return("", errors.New("this is a test"))
//	// Act
//	resp, err := suite.mockServerService.Translate("foo", "bar")
//	// Assert
//	suite.EqualError(err, "error in api")
//	suite.Equal(resp, "")
//}
//func (suite *HelloClientSuite) TestCall_InvalidJSON() {
//	// Arrange
//	suite.mockServerService.On("Translate", "foo", "bar").Return(`invalid json`, nil)
//	// Act
//	resp, err := suite.mockServerService.Translate("foo", "bar")
//	// Assert
//	suite.EqualError(err, "unable to decode message")
//	suite.Equal(resp, "")
//}
