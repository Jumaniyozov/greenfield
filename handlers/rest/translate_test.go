package rest_test

import (
	"encoding/json"
	"github.com/jumaniyozov/greenfield/handlers/rest"
	"github.com/jumaniyozov/greenfield/translation"
	"net/http"
	"net/http/httptest"
	"testing"
)

//type stubbedService struct{}
//
//func (s *stubbedService) Translate(word string, language string) string {
//	if word == "foo" {
//		return "bar"
//	}
//	return ""
//}

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	//	{
	//	Endpoint:            "/translate/foo?language=GerMan",
	//		StatusCode:          200,
	//		ExpectedLanguage:    "german",
	//		ExpectedTranslation: "bar",
	//	},
	//}
	//
	//h := rest.NewTranslateHandler(&stubbedService{})
	//handler := http.HandlerFunc(h.TranslateHandler)

	underTest := rest.NewTranslateHandler(translation.NewStaticService())
	handler := http.HandlerFunc(underTest.TranslateHandler)
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)
		handler.ServeHTTP(rr, req)
		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`,
				test.StatusCode, rr.Code)
		}
		var resp rest.Resp
		_ = json.Unmarshal(rr.Body.Bytes(), &resp)
		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s`,
				test.ExpectedLanguage, resp.Language)
		}
		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received %s`,
				test.ExpectedTranslation, resp.Translation)
		}
	}
}
