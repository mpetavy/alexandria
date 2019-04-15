package openssl_test

import (
	"testing"

	"github.com/mpetavy/alexandria/openssl"
)

func TestCheckInstalled(t *testing.T) {
	_, err := openssl.CheckInstalled()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertDER2PEM(t *testing.T) {
	derFile := "mp.p7b"
	pemFile := "mp.pem"

	err := openssl.ConvertDER2PEM(derFile, pemFile)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestExportPubkey(t *testing.T) {
	pemFile := "mp.pem"
	pubkeyFile := "mp.pubkey"

	err := openssl.ExportPubkey(pemFile, pubkeyFile)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestVerify(t *testing.T) {
	pemFile := "mp.pem"
	signatureFile := "mp.seckey"
	messageFile := "mp.message"

	err := openssl.Verify(pemFile, signatureFile, messageFile)

	if err != nil {
		t.Error(err.Error())
	}
}
