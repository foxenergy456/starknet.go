package rpc

import (
	"context"
	_ "embed"
	"testing"
)

// TestTransactionByHash tests transaction by hash
func TestTransactionByHash(t *testing.T) {
	testConfig := beforeEach(t)
	defer testConfig.client.Close()

	type testSetType struct {
		TxHash                     string
		ExpectedContractAddress    string
		ExpectedEntrypointSelector string
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				TxHash:                     "0xdeadbeef",
				ExpectedContractAddress:    "0xdeadbeef",
				ExpectedEntrypointSelector: "0xdeadbeef",
			},
		},
		"testnet": {
			{
				TxHash:                     "0x705547f8f2f8fdfb10ed533d909f76482bb293c5a32648d476774516a0bebd0",
				ExpectedContractAddress:    "0x315e364b162653e5c7b23efd34f8da27ba9c069b68e3042b7d76ce1df890313",
				ExpectedEntrypointSelector: "0x15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad",
			},
		},
		"mainnet": {
			{
				TxHash:                     "0x5f904b9185d4ed442846ac7e26bc4c60249a2a7f0bb85376c0bc7459665bae6",
				ExpectedContractAddress:    "0x3b4be7def2fc08589348966255e101824928659ebb724855223ff3a8c831efa",
				ExpectedEntrypointSelector: "0x2913ee03e5e3308c41e308bd391ea4faac9b9cb5062c76a6b3ab4f65397e106",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		tx, err := testConfig.client.TransactionByHash(context.Background(), test.TxHash)
		if err != nil {
			t.Fatal(err)
		}
		if tx == nil || tx.TransactionHash != test.TxHash {
			t.Fatal("transaction should exist and match the tx hash")
		}
		if tx.ContractAddress != test.ExpectedContractAddress {
			t.Fatalf("expecting contract %s, got %s", test.ExpectedContractAddress, tx.ContractAddress)
		}
		if tx.EntryPointSelector != test.ExpectedEntrypointSelector {
			t.Fatalf("expecting entrypoint %s, got %s", test.ExpectedEntrypointSelector, tx.EntryPointSelector)
		}
	}
}

// TestTransactionReceipt tests transaction receipt
func TestTransactionReceipt(t *testing.T) {

	testConfig := beforeEach(t)
	defer testConfig.client.Close()

	type testSetType struct {
		TxHash         string
		ExpectedStatus string
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				TxHash:         "0xdeadbeef",
				ExpectedStatus: "ACCEPTED_ON_L1",
			},
		},
		"testnet": {
			{
				TxHash:         "0x705547f8f2f8fdfb10ed533d909f76482bb293c5a32648d476774516a0bebd0",
				ExpectedStatus: "ACCEPTED_ON_L1",
			},
		},
		"mainnet": {
			{
				TxHash:         "0x5f904b9185d4ed442846ac7e26bc4c60249a2a7f0bb85376c0bc7459665bae6",
				ExpectedStatus: "ACCEPTED_ON_L1",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		txReceipt, err := testConfig.client.TransactionReceipt(context.Background(), test.TxHash)
		if err != nil {
			t.Fatal(err)
		}
		if txReceipt == nil || txReceipt.TransactionHash != test.TxHash {
			t.Fatal("transaction should exist and match the tx hash")
		}
		if txReceipt.Status != test.ExpectedStatus {
			t.Fatalf("expecting status %s, got %s", test.ExpectedStatus, txReceipt.Status)
		}
	}
}
