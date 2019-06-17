package services

import (
	"errors"
	"github.com/miguelbemartin/vat/models"
)

type RatesService struct {
}

func NewRatesService() *RatesService {
	return &RatesService{
	}
}

func (s *RatesService) Get(code string) (*models.Rate, error) {
	rates := map[string]models.Rate{
		"ES": {
			Country:      "Spain",
			CountryCode:  "ES",
			SuperReduced: 4.0,
			Reduced:      10.0,
			Standard:     21.0,
		},
		"BG": {
			Country:     "Bulgaria",
			CountryCode: "BG",
			Reduced:     9.0,
			Standard:    20.0,
		},
		"HU": {
			Country:     "Hungary",
			CountryCode: "HU",
			Reduced1:    5.0,
			Reduced2:    18.0,
			Standard:    27.0,
		},
		"LV": {
			Country:     "Latvia",
			CountryCode: "LV",
			Reduced:     12.0,
			Standard:    21.0,
		},
		"PL": {
			Country:     "Poland",
			CountryCode: "PL",
			Reduced1:    5.0,
			Reduced2:    8.0,
			Standard:    23.0,
		},
		"UK": {
			Country:     "United Kingdom",
			CountryCode: "GB",
			Standard:    20.0,
			Reduced:     5.0,
		},
		"CZ": {
			Country:     "Czech Republic",
			CountryCode: "CZ",
			Reduced:     15.0,
			Standard:    21.0,
		},
		"MT": {
			Country:     "Malta",
			CountryCode: "MT",
			Reduced1:    5.0,
			Reduced2:    7.0,
			Standard:    18.0,
		},
		"IT": {
			Country:      "Italy",
			CountryCode:  "IT",
			SuperReduced: 4.0,
			Reduced:      10.0,
			Standard:     22.0,
		},
		"SI": {
			Country:     "Slovenia",
			CountryCode: "SI",
			Reduced:     9.5,
			Standard:    22.0,
		},
		"IE": {
			Country:      "Ireland",
			CountryCode:  "IE",
			SuperReduced: 4.8,
			Reduced1:     9.0,
			Reduced2:     13.5,
			Standard:     23.0,
			Parking:      13.5,
		},
		"SE": {
			Country:     "Sweden",
			CountryCode: "SE",
			Reduced1:    6.0,
			Reduced2:    12.0,
			Standard:    25.0,
		},
		"DK": {
			Country:     "Denmark",
			CountryCode: "DK",
			Standard:    25.0,
		},
		"FI": {
			Country:     "Finland",
			CountryCode: "FI",
			Reduced1:    10.0,
			Reduced2:    14.0,
			Standard:    24.0,
		},
		"CY": {
			Country:     "Cyprus",
			CountryCode: "CY",
			Reduced1:    5.0,
			Reduced2:    9.0,
			Standard:    19.0,
		},
		"LU": {
			Country:      "Luxembourg",
			CountryCode:  "LU",
			SuperReduced: 3.0,
			Reduced1:     8.0,
			Standard:     17.0,
			Parking:      13.0,
		},
		"RO": {
			Country:     "Romania",
			CountryCode: "RO",
			Reduced1:    5.0,
			Reduced2:    9.0,
			Standard:    19.0,
		},
		"EE": {
			Country:     "Estonia",
			CountryCode: "EE",
			Reduced:     9.0,
			Standard:    20.0,
		},
		"EL": {
			Country:     "Greece",
			CountryCode: "GR",
			Reduced1:    6.0,
			Reduced2:    13.5,
			Standard:    24.0,
		},
		"LT": {
			Country:     "Lithuania",
			CountryCode: "LT",
			Reduced1:    5.0,
			Reduced2:    9.0,
			Standard:    21.0,
		},
		"FR": {
			Country:      "France",
			CountryCode:  "FR",
			SuperReduced: 2.1,
			Reduced1:     5.5,
			Reduced2:     10.0,
			Standard:     20.0,
		},
		"HR": {
			Country:     "Croatia",
			CountryCode: "HR",
			Reduced1:    5.0,
			Reduced2:    13.0,
			Standard:    25.0,
		},
		"BE": {
			Country:     "Belgium",
			CountryCode: "BE",
			Reduced1:    6.0,
			Reduced2:    12.0,
			Standard:    21.0,
			Parking:     12.0,
		},
		"NL": {
			Country:     "Netherlands",
			CountryCode: "NL",
			Reduced:     6.0,
			Standard:    21.0,
		},
		"SK": {
			Country:     "Slovakia",
			CountryCode: "SK",
			Reduced:     10.0,
			Standard:    20.0,
		},
		"DE": {
			Country:     "Germany",
			CountryCode: "DE",
			Reduced:     7.0,
			Standard:    19.0,
		},
		"PT": {
			Country:     "Portugal",
			CountryCode: "PT",
			Reduced1:    6.0,
			Reduced2:    13.0,
			Standard:    23.0,
			Parking:     13.0,
		},
		"AT": {
			Country:     "Austria",
			CountryCode: "AT",
			Reduced1:    10.0,
			Reduced2:    13.0,
			Standard:    20.0,
			Parking:     13.0,
		},
	}

	rateCountry, ok := rates[code]
	if !ok {
		return nil, errors.New("rate: code not found")
	}

	return &rateCountry, nil
}
