// Package ukrpost implements post tracking for http://ukrposhta.ua/
package ukrpost

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseUrl = `http://services.ukrposhta.com`
	methUrl = `/barcodestatistic/barcodestatistic.asmx`
	defGuid = `fcc8d9e1-b6f9-438f-9ac8-b67ab44391dd`
)

type TrackInfo struct {
	Barcode         string `xml:"barcode"`
	Code            Int    `xml:"code,omitempty"`
	LastOfficeIndex string `xml:"lastofficeindex,omitempty"`
	LastOffice      string `xml:"lastoffice,omitempty"`
	EventDate       Date   `xml:"eventdate"`
	EventDesc       string `xml:"eventdescription"`
}

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

// Call invokes ukrpost api method
func (s *Service) Call(meth string, args url.Values, out interface{}) error {
	resp, err := http.Get(baseUrl + methUrl + "/" + meth + "?" + args.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return xml.NewDecoder(resp.Body).Decode(out)
}

// Track returns delivery status of postal item with the given barcode
func (s *Service) Track(barcode string) (TrackInfo, error) {
	vals := url.Values{
		"guid":    {s.Guid},
		"barcode": {barcode},
		"culture": {s.Lang},
	}
	var status TrackInfo
	if err := s.Call(`GetBarcodeInfo`, vals, &status); err != nil {
		return TrackInfo{}, err
	}
	status.EventDesc = strings.Trim(status.EventDesc, " \n\r")
	return status, nil
}
