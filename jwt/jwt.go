package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// https://gist.github.com/c-bata/488c4dbfae807ad113c07195e7388e09

// CertsAPIEndpoint is endpoint of getting Public Key.
var CertsAPIEndpoint = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

// GetCertificate is useful for testing.
var GetCertificate = getCertificate

func getCertificates() (certs map[string]string, err error) {
	res, err := http.Get(CertsAPIEndpoint)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	json.Unmarshal(data, &certs)
	return
}

// GetCertificate returns certificate.
func getCertificate(kid string) (cert []byte, err error) {
	certs, err := getCertificates()
	if err != nil {
		return
	}
	certString := certs[kid]
	cert = []byte(certString)
	err = nil
	return
}

// VerifyJWT to verify the token header, payload and signature.
var VerifyJWT = verifyJWT

// GetCertificateFromToken returns cert from token.
func GetCertificateFromToken(token *jwt.Token) ([]byte, error) {
	// Get kid
	kid, ok := token.Header["kid"]
	if !ok {
		return []byte{}, errors.New("kid not found")
	}
	kidString, ok := kid.(string)
	if !ok {
		return []byte{}, errors.New("kid cast error to string")
	}
	return GetCertificate(kidString)
}

// Verify the token payload.
func verifyPayload(t *jwt.Token, projectID string) (ok bool, uid string) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	// Verify User
	claimsAud, ok := claims["aud"].(string)
	if claimsAud != projectID || !ok {
		return
	}
	// Verify issued at
	iss := "https://securetoken.google.com/" + projectID
	claimsIss, ok := claims["iss"].(string)
	if claimsIss != iss || !ok {
		return
	}
	// sub is uid of user.
	uid, ok = claims["sub"].(string)
	if !ok {
		return
	}
	return
}

func readPublicKey(cert []byte) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode(cert)
	if publicKeyBlock == nil {
		return nil, errors.New("invalid public key data")
	}
	if publicKeyBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("invalid public key type: %s", publicKeyBlock.Type)
	}
	c, err := x509.ParseCertificate(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}
	return publicKey, nil
}

func verifyJWT(t, projectID string) (uid string, ok bool) {
	parsed, _ := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		cert, err := GetCertificateFromToken(t)
		if err != nil {
			return "", err
		}
		publicKey, err := readPublicKey(cert)
		if err != nil {
			return "", err
		}
		return publicKey, nil
	})

	ok = parsed.Valid
	if !ok {
		return
	}
	// Verify header.
	if parsed.Header["alg"] != "RS256" {
		ok = false
		return
	}
	// Verify payload.
	ok, uid = verifyPayload(parsed, projectID)
	return
}
