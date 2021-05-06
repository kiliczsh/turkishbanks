package turkishbanks

import (
	"encoding/json"
	"testing"
)

func TestGetBankListAsJson(t *testing.T) {
	list, err := GetBankListAsJson()
	if err != nil {
		t.Errorf("TurkishBanks | Tests | GetBankListAsJson failed: %s", err)
	}

	var banks BanksBranchList
	err = json.Unmarshal(list, &banks)
	if err != nil {
		return
	}

	if len(banks.BankBranches) <= 0 {
		t.Errorf("TurkishBanks | Tests | List is empty.")
	}
}

func TestGetBankList(t *testing.T) {
	list, err := GetBankList()
	if err != nil {
		t.Errorf("TurkishBanks | Tests | GetBankList failed: %s", err)
	}

	if len(list.BankBranches) <= 0 {
		t.Errorf("TurkishBanks | Tests | List is empty.")
	}
}
