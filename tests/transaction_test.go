package tests

import (
	"testing"
)

func TestTransactions(t *testing.T) {
	notWorking := make(map[string]bool, 100)

	// TODO: all these tests should work! remove them from the array when they work
	snafus := []string{
		"TransactionWithHihghNonce256", // fails due to testing upper bound of 256 bit nonce
	}

	for _, name := range snafus {
		notWorking[name] = true
	}

	var err error
	err = RunTransactionTests("./files/TransactionTests/ttTransactionTest.json",
		notWorking)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrongRLPTransactions(t *testing.T) {
	notWorking := make(map[string]bool, 100)
	var err error
	err = RunTransactionTests("./files/TransactionTests/ttWrongRLPTransaction.json",
		notWorking)
	if err != nil {
		t.Fatal(err)
	}
}

func Test10MBtx(t *testing.T) {
	notWorking := make(map[string]bool, 100)
	var err error
	err = RunTransactionTests("./files/TransactionTests/tt10mbDataField.json",
		notWorking)
	if err != nil {
		t.Fatal(err)
	}
}
