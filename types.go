package ukrpost

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format("02.01.2006")
}
func (d Date) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeToken(xml.CharData([]byte(d.String())))
}
func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := dec.DecodeElement(&s, &start); err != nil {
		return err
	}
	if strings.Trim(s, " ") == "" {
		*d = Date(time.Time{})
		return nil
	}
	t, err := time.Parse("02.01.2006", s)
	*d = Date(t)
	return err
}

type Int int

func (d Int) String() string {
	return fmt.Sprint(d)
}
func (d Int) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeToken(xml.CharData([]byte(d.String())))
}
func (d *Int) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := dec.DecodeElement(&s, &start); err != nil {
		return err
	}
	if strings.Trim(s, " ") == "" {
		*d = Int(0)
		return nil
	}
	t, err := strconv.ParseInt(s, 10, 64)
	*d = Int(t)
	return err
}
