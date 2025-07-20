package mocks

import (
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/domain"
)

type MockWeatherApi struct {
	MockGetFullAddressByCep func(cep string) (*domain.Cep, error)
	MockGetWeatherInfo      func(fullAddress *domain.Cep, weatherApiKey string) (*domain.Weather, error)
}

func (m *MockWeatherApi) GetFullAddressByCep(cep string) (*domain.Cep, error) {
	return m.MockGetFullAddressByCep(cep)
}

func (m *MockWeatherApi) GetWeatherInfo(fullAddress *domain.Cep, weatherApiKey string) (*domain.Weather, error) {
	return m.MockGetWeatherInfo(fullAddress, weatherApiKey)
}
