package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"
	"time"

	"github.com/NamelessOne91/go-tls/pkg/key"
)

func CreateCACert(ca *CACert, keyFilePath, caCertFilePath string) error {
	template := &x509.Certificate{
		SerialNumber: ca.Serial,
		Subject: pkix.Name{
			Country:            removeEmptyString([]string{ca.Subject.Country}),
			Organization:       removeEmptyString([]string{ca.Subject.Organization}),
			OrganizationalUnit: removeEmptyString([]string{ca.Subject.OrganizationalUnit}),
			Locality:           removeEmptyString([]string{ca.Subject.Locality}),
			Province:           removeEmptyString([]string{ca.Subject.Province}),
			StreetAddress:      removeEmptyString([]string{ca.Subject.StreetAddress}),
			PostalCode:         removeEmptyString([]string{ca.Subject.PostalCode}),
			CommonName:         ca.Subject.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(ca.ValidForYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	keyBytes, certBytes, err := createCert(template, nil, nil)
	if err != nil {
		return err
	}

	if err := os.WriteFile(keyFilePath, keyBytes, 0600); err != nil {
		return err
	}
	if err := os.WriteFile(caCertFilePath, certBytes, 0644); err != nil {
		return err
	}

	return nil
}

func CreateCert(cert *Cert, caKey []byte, caCert []byte, keyFilePath, certFilePath string) error {
	template := &x509.Certificate{
		SerialNumber: cert.Serial,
		Subject: pkix.Name{
			Country:            removeEmptyString([]string{cert.Subject.Country}),
			Organization:       removeEmptyString([]string{cert.Subject.Organization}),
			OrganizationalUnit: removeEmptyString([]string{cert.Subject.OrganizationalUnit}),
			Locality:           removeEmptyString([]string{cert.Subject.Locality}),
			Province:           removeEmptyString([]string{cert.Subject.Province}),
			StreetAddress:      removeEmptyString([]string{cert.Subject.StreetAddress}),
			PostalCode:         removeEmptyString([]string{cert.Subject.PostalCode}),
			CommonName:         cert.Subject.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(cert.ValidForYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	_, _, err := createCert(template, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func createCert(template *x509.Certificate, caKey *rsa.PrivateKey, caCert *x509.Certificate) ([]byte, []byte, error) {
	var (
		derBytes []byte
		certOut  bytes.Buffer
		keyOut   bytes.Buffer
	)

	pk, err := key.CreateRSAPrivateKey(4096)
	if err != nil {
		return nil, nil, err
	}
	if template.IsCA {
		derBytes, err = x509.CreateCertificate(rand.Reader, template, template, &pk.PublicKey, pk)
		if err != nil {
			return nil, nil, err
		}
	} else {

	}

	if err := pem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, nil, err
	}
	if err := pem.Encode(&keyOut, key.RSAPrivateKeyToPem(pk)); err != nil {
		return nil, nil, err
	}

	return keyOut.Bytes(), certOut.Bytes(), nil
}

func removeEmptyString(input []string) []string {
	if len(input) == 1 && input[0] == "" {
		return []string{}
	}
	return input
}