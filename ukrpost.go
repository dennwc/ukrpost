// Package ukrpost implements post tracking and index search for http://ukrposhta.ua/
package ukrpost

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

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

// callAPI invokes ukrpost api method on a specific endpoint
func (s *Service) callAPI(endpoint, meth string, args url.Values, out interface{}) error {
	resp, err := http.Get(baseUrl + endpoint + "/" + meth + "?" + args.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return xml.NewDecoder(resp.Body).Decode(out)
}
