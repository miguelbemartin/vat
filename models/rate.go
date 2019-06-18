package models

// Rate represents the vat rates per country.
type Rate struct {
	Country      string
	CountryCode  string
	SuperReduced float64
	Reduced      float64
	Reduced1     float64
	Reduced2     float64
	Standard     float64
	Parking      float64
}
