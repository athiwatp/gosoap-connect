package gosoap_connect

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_SoapConnect(t *testing.T) {

	type InterFaceTest struct {
		CountryName string `xml:"countryName"`
	}
	type ResultInterFaceTest struct {
		XMLName                      xml.Name `xml:"Envelope"`
		GetCountryISO2ByNameResponse string   `xml:"Body>GetCountryISO2ByNameResponse>GetCountryISO2ByNameResult"`
	}

	t.Run("Happy - simple call soap", func(t *testing.T) {
		mockInterface := &InterFaceTest{CountryName: "thailand"}
		respSoap, err := SoapCall("POST", "http://wsgeoip.lavasoft.com/ipservice.asmx?op=GetCountryISO2ByName", "http://lavasoft.com/GetCountryISO2ByName", mockInterface)

		assert.NoError(t, err)
		assert.NotNil(t, respSoap)
	})

	t.Run("Happy - simple call with unmarshal", func(t *testing.T) {
		mockInterface := &InterFaceTest{CountryName: "thailand"}
		mockResultInterface := &ResultInterFaceTest{}
		err := SoapCallHandleResponse("POST", "http://wsgeoip.lavasoft.com/ipservice.asmx?op=GetCountryISO2ByName", "http://lavasoft.com/GetCountryISO2ByName", mockInterface, mockResultInterface)
		log.Printf("===> %+v", mockResultInterface.GetCountryISO2ByNameResponse)

		assert.NoError(t, err)
	})

}
