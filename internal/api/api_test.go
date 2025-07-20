package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/config/getenv"
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/domain"
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindTemperatureByCEP_WeatherSuccess(t *testing.T) {
	mockApi := &mocks.MockWeatherApi{
		MockGetFullAddressByCep: func(cep string) (*domain.Cep, error) {
			return &domain.Cep{Localidade: "São Paulo"}, nil
		},
		MockGetWeatherInfo: func(fullAddress *domain.Cep, weatherApiKey string) (*domain.Weather, error) {
			return &domain.Weather{
				Current: domain.Current{TempC: 25, TempF: 77},
			}, nil
		},
	}

	wAPI = mockApi

	req := httptest.NewRequest("GET", "/?cep=12345678", nil)
	req.Header.Set("WEATHER_API_KEY", "xxx")

	rr := httptest.NewRecorder()

	FindTemperatureByCEP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "25")
	assert.Contains(t, rr.Body.String(), "77")
	assert.Contains(t, rr.Body.String(), "298")

}

func TestFindTemperatureByCEP_WeatherCEPError(t *testing.T) {

	req := httptest.NewRequest("GET", "/?cep=xxx", nil)
	rr := httptest.NewRecorder()

	FindTemperatureByCEP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid cep")
}

func TestFindTemperatureByCEP_WeatherApiKeyError(t *testing.T) {

	req := httptest.NewRequest("GET", "/?cep=12345678", nil)
	req.Header.Set("WEATHER_API_KEY", "")

	rr := httptest.NewRecorder()

	FindTemperatureByCEP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "WEATHER_API_KEY header is mandatory")
}

func TestFindTemperatureByCEP_WeatherCepNotFound(t *testing.T) {

	getenv.ApiCepUrl = "http://viacep.com.br/ws/%s/json/"

	req := httptest.NewRequest("GET", "/?cep=12345678", nil)
	req.Header.Set("WEATHER_API_KEY", "xxx")

	rr := httptest.NewRecorder()

	FindTemperatureByCEP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "can not find zipcode")
}

func TestFindTemperatureByCEP_GetFullAddressByCEPError(t *testing.T) {

	getenv.ApiCepUrl = "httpssss://viacep.com.br/ws/%s/json/"

	req := httptest.NewRequest("GET", "/?cep=12345678", nil)
	req.Header.Set("WEATHER_API_KEY", "xxx")

	rr := httptest.NewRecorder()

	FindTemperatureByCEP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
