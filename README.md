# TurkishBanks

[![Go Reference](https://pkg.go.dev/badge/github.com/kiliczsh/turkishbanks.svg)](https://pkg.go.dev/github.com/kiliczsh/turkishbanks)

Go module for Turkish Banks. 
Bank information is taken from [Türkiye Cumhuriyet Merkez Bankası (TCMB)](http://eftemkt.tcmb.gov.tr/bankasubelistesi/bankaSubeTumListe.xml).

## Sample Usage

### Install

```shell
go get github.com/kiliczsh/turkishbanks@v1.0.0
```


### Code
```go
package main

import (
	"fmt"
	"github.com/kiliczsh/turkishbanks"
)

func main() {
	banks, err := TurkishBanks.GetBankList()
	if err != nil {
		return
	}
	for i := 0; i < len(banks.BankBranches); i++ {
		b := banks.BankBranches[i]
		fmt.Printf("%s - %s has %s branches.\n", b.BankCode, b.Bank.BankName, b.BranchSize)
	}
}

```

### Output

```
0001 - TÜRKİYE CUMHURİYET MERKEZ BANKASI has 24 branches.
0004 - İLLER BANKASI A.Ş. has 3 branches.
0010 - TÜRKİYE CUMHURİYETİ ZİRAAT BANKASI A.Ş. has 1784 branches.
0012 - TÜRKİYE HALK BANKASI A.Ş. has 1034 branches.
0014 - TÜRKİYE SINAİ KALKINMA BANKASI A.Ş. has 3 branches.
0015 - TÜRKİYE VAKIFLAR BANKASI T.A.O. has 942 branches.
0016 - TÜRK EXİMBANK has 4 branches.
0017 - TÜRKİYE KALKINMA VE YATIRIM BANKASI A.Ş. has 5 branches.
0029 - BİRLEŞİK FON BANKASI A.Ş. has 6 branches.
0032 - TÜRK EKONOMİ BANKASI A.Ş. has 462 branches.
0046 - AKBANK T.A.Ş. has 711 branches.
0059 - ŞEKERBANK T.A.Ş. has 247 branches.
0062 - TÜRKİYE GARANTİ BANKASI A.Ş. has 938 branches.
0064 - TÜRKİYE İŞ BANKASI A.Ş. has 1213 branches.
0067 - YAPI VE KREDİ BANKASI A.Ş. has 876 branches.
0091 - ARAP TÜRK BANKASI A.Ş. has 13 branches.
0092 - CITIBANK A.Ş. has 8 branches.
0096 - TURKISH BANK A.Ş. has 12 branches.
0098 - JPMORGAN CHASE BANK N.A. has 3 branches.
0099 - ING BANK A.Ş. has 197 branches.
0100 - ADABANK A.Ş. has 5 branches.
0103 - FİBABANKA A.Ş. has 59 branches.
0108 - TURKLAND BANK A.Ş. has 20 branches.
0109 - ICBC TURKEY BANK A.Ş. has 46 branches.
0111 - QNB FİNANSBANK A.Ş. has 548 branches.
0115 - DEUTSCHE BANK A.Ş. has 5 branches.
0116 - PASHA YATIRIM BANKASI A.Ş. has 5 branches.
0121 - STANDARD CHARTERED YATIRIM BANKASI TÜRK A.Ş. has 3 branches.
0122 - SOCIETE GENERALE (SA) has 3 branches.
0123 - HSBC BANK A.Ş. has 90 branches.
0124 - ALTERNATİFBANK A.Ş. has 55 branches.
0125 - BURGAN BANK A.Ş. has 44 branches.
0129 - BANK OF AMERICA YATIRIM BANK A.Ş. has 3 branches.
0132 - İSTANBUL TAKAS VE SAKLAMA BANKASI A.Ş. has 212 branches.
0134 - DENİZBANK A.Ş. has 725 branches.
0135 - ANADOLUBANK A.Ş. has 123 branches.
0137 - RABOBANK A.Ş. has 5 branches.
0138 - DİLER YATIRIM BANKASI A.Ş. has 6 branches.
0139 - GSD YATIRIM BANKASI A.Ş. has 8 branches.
0141 - NUROL YATIRIM BANKASI A.Ş. has 4 branches.
0142 - BANKPOZİTİF KREDİ VE KALKINMA BANKASI A.Ş. has 8 branches.
0143 - AKTİF YATIRIM BANKASI A.Ş. has 23 branches.
0146 - ODEA BANK A.Ş. has 57 branches.
0147 - MUFG Bank Turkey A.Ş. has 5 branches.
0148 - INTESA SANPAOLO S.P.A. has 3 branches.
0149 - BANK OF CHINA TURKEY A.Ş. has 6 branches.
0150 - GOLDEN GLOBAL YATIRIM BANKASI A.Ş. has 3 branches.
0203 - ALBARAKA TÜRK KATILIM BANKASI A.Ş. has 247 branches.
0205 - KUVEYT TÜRK KATILIM BANKASI A.Ş. has 445 branches.
0206 - TÜRKİYE FİNANS KATILIM BANKASI A.Ş. has 328 branches.
0209 - ZİRAAT KATILIM BANKASI A.Ş. has 113 branches.
0210 - VAKIF KATILIM BANKASI A.Ş. has 125 branches.
0211 - TÜRKİYE EMLAK KATILIM BANKASI A.Ş. has 68 branches.
0806 - MERKEZİ KAYIT KURULUŞU A.Ş. has 168 branches.
0807 - POSTA VE TELGRAF TEŞKİLATI A.Ş. has 1026 branches.

```
