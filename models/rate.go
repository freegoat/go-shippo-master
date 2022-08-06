package models

// See https://goshippo.com/docs/reference#rates
type Rate struct {
	CommonOutputFields
	Attributes       []string         `json:"attributes,omitempty"`
	AmountLocal      string           `json:"amount_local,omitempty"`
	CurrencyLocal    string           `json:"currency_local,omitempty"`
	Amount           string           `json:"amount,omitempty"`
	Currency         string           `json:"currency,omitempty"`
	Provider         string           `json:"provider,omitempty"`
	ProviderImage75  string           `json:"provider_image_75,omitempty"`
	ProviderImage200 string           `json:"provider_image_200,omitempty"`
	ServiceLevel     *ServiceLevel    `json:"servicelevel,omitempty"`
	Days             int              `json:"days,omitempty"`
	DurationTerms    string           `json:"duration_terms,omitempty"`
	CarrierAccount   string           `json:"carrier_account,omitempty"`
	Zone             string           `json:"zone,omitempty"`
	Messages         []*OutputMessage `json:"messages,omitempty"`
	Test             bool             `json:"test,omitempty"`
}

type ServiceLevel struct {
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"` // https://goshippo.com/docs/reference#servicelevels
	Terms string `json:"terms,omitempty"`
}
