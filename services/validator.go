package services

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	serviceURL = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"
)

// ValidatorService handles validation for european vats.
type ValidatorService struct {
}

// NewValidatorService creates a new handler for this service.
func NewValidatorService() *ValidatorService {
	return &ValidatorService{}
}

// Validate will validate the vat number format and the
// existence by using the VIES VAT API (using SOAP).
func (s *ValidatorService) Validate(number string) (bool, error) {
	format, err := s.ValidateNumberFormat(number)
	existence := false

	if format {
		existence, err = s.ValidateNumberExistence(number)
	}

	return format && existence, err
}

// ValidateNumberFormat validates a VAT number by its format.
func (s *ValidatorService) ValidateNumberFormat(n string) (bool, error) {
	patterns := map[string]string{
		"AT": `U[A-Z0-9]{8}`,
		"BE": `(0[0-9]{9}|[0-9]{10})`,
		"BG": `[0-9]{9,10}`,
		"CH": `(?:E(?:-| )[0-9]{3}(?:\.| )[0-9]{3}(?:\.| )[0-9]{3}( MWST)?|E[0-9]{9}(?:MWST)?)`,
		"CY": `[0-9]{8}[A-Z]`,
		"CZ": `[0-9]{8,10}`,
		"DE": `[0-9]{9}`,
		"DK": `[0-9]{8}`,
		"EE": `[0-9]{9}`,
		"EL": `[0-9]{9}`,
		"ES": `[A-Z][0-9]{7}[A-Z]|[0-9]{8}[A-Z]|[A-Z][0-9]{8}`,
		"FI": `[0-9]{8}`,
		"FR": `([A-Z]{2}|[0-9]{2})[0-9]{9}`,
		"GB": `[0-9]{9}|[0-9]{12}|(GD|HA)[0-9]{3}`,
		"HR": `[0-9]{11}`,
		"HU": `[0-9]{8}`,
		"IE": `[A-Z0-9]{7}[A-Z]|[A-Z0-9]{7}[A-W][A-I]`,
		"IT": `[0-9]{11}`,
		"LT": `([0-9]{9}|[0-9]{12})`,
		"LU": `[0-9]{8}`,
		"LV": `[0-9]{11}`,
		"MT": `[0-9]{8}`,
		"NL": `[0-9]{9}B[0-9]{2}`,
		"PL": `[0-9]{10}`,
		"PT": `[0-9]{9}`,
		"RO": `[0-9]{2,10}`,
		"SE": `[0-9]{12}`,
		"SI": `[0-9]{8}`,
		"SK": `[0-9]{10}`,
	}

	if len(n) < 3 {
		return false, nil
	}

	n = strings.ToUpper(n)
	pattern, ok := patterns[n[0:2]]
	if !ok {
		return false, nil
	}

	matched, err := regexp.MatchString(pattern, n[2:])
	return matched, err
}

// ValidateNumberExistence validates a VAT number by its existence using the VIES VAT API (using SOAP)
func (s *ValidatorService) ValidateNumberExistence(n string) (bool, error) {
	r, err := s.lookup(n)
	if err != nil {
		return false, err
	}
	return r.Valid, nil
}

// lookup returns *viesResponse for a VAT number
func (s *ValidatorService) lookup(vatNumber string) (*viesResponse, error) {
	if len(vatNumber) < 3 {
		return nil, errors.New("vat: vat number is invalid")
	}

	e := getEnvelope(vatNumber)
	eb := bytes.NewBufferString(e)

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	res, err := client.Post(serviceURL, "text/xml;charset=UTF-8", eb)
	if err != nil {
		return nil, errors.New("vat: service is unreachable")
	}
	defer res.Body.Close()

	xmlRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// check if response contains "INVALID_INPUT" string
	if bytes.Contains(xmlRes, []byte("INVALID_INPUT")) {
		return nil, errors.New("vat: vat number is invalid")
	}

	var rd struct {
		XMLName xml.Name `xml:"Envelope"`
		Soap    struct {
			XMLName xml.Name `xml:"Body"`
			Soap    struct {
				XMLName     xml.Name `xml:"checkVatResponse"`
				CountryCode string   `xml:"countryCode"`
				VATNumber   string   `xml:"vatNumber"`
				RequestDate string   `xml:"requestDate"` // 2015-03-06+01:00
				Valid       bool     `xml:"valid"`
				Name        string   `xml:"name"`
				Address     string   `xml:"address"`
			}
		}
	}
	if err = xml.Unmarshal(xmlRes, &rd); err != nil {
		return nil, err
	}

	r := &viesResponse{
		CountryCode: rd.Soap.Soap.CountryCode,
		VATNumber:   rd.Soap.Soap.VATNumber,
		RequestDate: rd.Soap.Soap.RequestDate,
		Valid:       rd.Soap.Soap.Valid,
		Name:        rd.Soap.Soap.Name,
		Address:     rd.Soap.Soap.Address,
	}

	return r, nil
}

// getEnvelope parses envelope template
func getEnvelope(n string) string {
	n = strings.ToUpper(n)
	countryCode := n[0:2]
	vatNumber := n[2:]
	const envelopeTemplate = `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	<soapenv:Header/>
	<soapenv:Body>
	  <checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
	    <countryCode>{{.countryCode}}</countryCode>
	    <vatNumber>{{.vatNumber}}</vatNumber>
	  </checkVat>
	</soapenv:Body>
	</soapenv:Envelope>
	`

	e := envelopeTemplate
	e = strings.Replace(e, "{{.countryCode}}", countryCode, 1)
	e = strings.Replace(e, "{{.vatNumber}}", vatNumber, 1)
	return e
}

// handle the response of the service.
type viesResponse struct {
	CountryCode string
	VATNumber   string
	RequestDate string
	Valid       bool
	Name        string
	Address     string
}
