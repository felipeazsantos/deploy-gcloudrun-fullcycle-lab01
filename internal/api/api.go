package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/config/getenv"
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/domain"
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/dto"
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/validation"
)

func FindTemperatureByCEP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cep := r.URL.Query().Get("cep")
	weatherApiKey := r.Header.Get("WEATHER_API_KEY")

	if err := weatherRequestValidation(cep, weatherApiKey); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fullAddress, err := getFullAddressByCep(cep)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	weatherInfo, err := getWeatherInfo(fullAddress, weatherApiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	kelvin := weatherInfo.Current.TempC + 273

	weatherResponse := dto.WeatherResponseDTO{
		TempC: weatherInfo.Current.TempC,
		TempF: weatherInfo.Current.TempF,
		TempK: kelvin,
	}

	json.NewEncoder(w).Encode(&weatherResponse)
}

func weatherRequestValidation(cep, weatherApiKey string) error {
	if !validation.IsValidCEP(cep) {
		return errors.New("invalid cep")
	}

	if strings.TrimSpace(weatherApiKey) == "" {
		return errors.New("WEATHER_API_KEY header is mandatory")
	}

	return nil
}

func getFullAddressByCep(cep string) (*domain.Cep, error) {
	urlApi := fmt.Sprintf(getenv.ApiCepUrl, cep)
	response, err := http.Get(urlApi)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *domain.Cep
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func getWeatherInfo(fullAddress *domain.Cep, weatherApiKey string) (*domain.Weather, error) {
	weatherApiUrl := fmt.Sprintf(getenv.ApiWeatherUrl, weatherApiKey, url.QueryEscape(fullAddress.Localidade))
	response, err := http.Get(weatherApiUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *domain.Weather
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
