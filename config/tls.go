package config

import (
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
)

func loadCACertPool(caPath string) (*x509.CertPool, error) {
	rootCertPool := x509.NewCertPool()

	pem, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	if ok := rootCertPool.AppendCertsFromPEM([]byte(strings.TrimSpace(string(pem)))); !ok {
		return nil, errors.New("invalid ca cert file provided")
	}

	return rootCertPool, nil
}

func loadCertify(vaultAddr string, vaultPki string, vaultRole string, vaultCn string, vaultToken string) (*certify.Certify, error) {
	url, err := url.Parse(vaultAddr)
	if err != nil {
		return nil, err
	}

	issuer := &vault.Issuer{
		URL:        url,
		Mount:      vaultPki,
		AuthMethod: &vault.RenewingToken{Initial: vaultToken},
		Role:       vaultRole,
		TimeToLive: time.Hour * 24,
	}

	certify := &certify.Certify{
		Issuer:      issuer,
		CommonName:  vaultCn,
		Cache:       certify.NewMemCache(),
		RenewBefore: time.Minute * 10,
	}
	return certify, nil
}
