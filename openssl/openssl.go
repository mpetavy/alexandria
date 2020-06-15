package openssl

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/mpetavy/common"
)

// CheckInstalled checks if OPENSSL is available
func CheckInstalled() (path string, err error) {
	path, err = exec.LookPath("openssl")
	if err != nil {
		return "", err
	}

	return path, nil
}

// ConvertDER2PEM convert a certificate from DER to PEM format
func ConvertDER2PEM(derFile, pemFile string) error {
	path, err := CheckInstalled()
	common.Error(err)

	cmd := exec.Command(path, "pkcs7", "-print_certs", "-inform", "der", "-in", derFile, "-outform", "pem", "-out", pemFile)
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	common.Error(err)

	return nil
}

// ExportPubkey expors the public key from certificate
func ExportPubkey(pemFile, pubkeyFile string) error {
	path, err := CheckInstalled()
	common.Error(err)

	cmd := exec.Command(path, "x509", "-pubkey", "-noout", "-in", pemFile, "-outform", "pem")
	if *common.FlagLogVerbose {
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	common.Error(err)

	pubkey, _ := cmd.Output()

	err = ioutil.WriteFile(pubkeyFile, pubkey, 0644)
	common.Error(err)

	return nil
}

// Verify the messagFile with the signatureFile by the certificate
func Verify(pemFile, signatureFile, messageFile string) error {
	path, err := CheckInstalled()
	common.Error(err)

	cmd := exec.Command(path, "smime", "-verify", "-inform", "der", "-in", signatureFile, "-content", messageFile, "-certfile", pemFile, "-noverify")
	if *common.FlagLogVerbose {
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	common.Error(err)

	return nil
}
