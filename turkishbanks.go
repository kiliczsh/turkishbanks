package turkishbanks

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var TurkishBanksURL = "http://eftemkt.tcmb.gov.tr/bankasubelistesi/bankaSubeTumListe.xml"

type BanksBranchList struct {
	XMLName      xml.Name       `xml:"bankaSubeTumListe"`
	BankBranches []BankBranches `xml:"bankaSubeleri"`
	Date         string         `xml:"tarih,attr"`
	Xmlns        string         `xml:"xmlns,attr"`
}

type BankBranches struct {
	XMLName    xml.Name `xml:"bankaSubeleri"`
	Bank       Bank     `xml:"banka"`
	Branch     []Branch `xml:"sube"`
	BankCode   string   `xml:"bKd,attr"`
	BranchSize string   `xml:"sAdt,attr"`
}

type Bank struct {
	XMLName        xml.Name `xml:"banka"`
	BankCode       string   `xml:"bKd"`
	BankName       string   `xml:"bAd"`
	BankCityName   string   `xml:"bIlAd"`
	Address        string   `xml:"adr"`
	LastUpdateType string   `xml:"sonIslemTuru,attr"`
	LastUpdateTime string   `xml:"sonIslemZamani,attr"`
}

type Branch struct {
	XMLName          xml.Name `xml:"sube"`
	BankCode         string   `xml:"bKd"`
	BranchCode       string   `xml:"sKd"`
	BranchName       string   `xml:"sAd"`
	BranchCityCode   string   `xml:"sIlKd"`
	BranchCityName   string   `xml:"sIlAd"`
	BranchCountyCode string   `xml:"sIlcKd"`
	BranchCountyName string   `xml:"sIlcAd"`
	Address          string   `xml:"adr"`
	PhoneNumber      string   `xml:"tlf"`
	Fax              string   `xml:"fks"`
	Email            string   `xml:"epst"`
	LastUpdateType   string   `xml:"sonIslemTuru,attr"`
	LastUpdateTime   string   `xml:"sonIslemZamani,attr"`
}

func GetBankListAsJson() ([]byte, error) {
	if banks, err := GetBankList(); err != nil {
		log.Printf("TurkishBanks | Failed to get XML: %v", err)
	} else {
		jsonObject, err := json.Marshal(banks)
		if err != nil {
			return nil, fmt.Errorf("TurkishBanks | Failed to parse data as json: %v", err)
		}
		return jsonObject, nil
	}
	return nil, errors.New("TurkishBanks | Method failed.")
}

func GetBankList() (BanksBranchList, error) {
	var list BanksBranchList
	if xmlBytes, err := getXML(TurkishBanksURL); err != nil {
		log.Printf("TurkishBanks | Failed to get XML: %v", err)
	} else {
		err := xml.Unmarshal(xmlBytes, &list)
		if err != nil {
			return BanksBranchList{}, err
		}
	}
	return list, nil
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("TurkishBanks | GET error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("TurkishBanks | Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("TurkishBanks | Read body: %v", err)
	}

	return data, nil
}
