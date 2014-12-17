// Package ukrpost implements post tracking for http://ukrposhta.ua/
package ukrpost

const (
	baseUrl = `http://services.ukrposhta.com`
	defGuid = `fcc8d9e1-b6f9-438f-9ac8-b67ab44391dd`
)

type Service struct {
	Guid string
	Lang string
}

// New creates new instance of ukrpost client with the given guid (api key).
//
// If guid parameter is empty it will be replaced with default test key.
func New(guid string) *Service {
	if guid == "" {
		guid = defGuid
	}
	return &Service{Guid: guid, Lang: "en"}
}
