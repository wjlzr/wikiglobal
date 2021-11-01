package apple

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/k0kubun/pp"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wiki_global/src/config"
)

type Result struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	Error        string `json:"error"`
}

var (
	// replace your configs here
	secret      = "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgTUzMIAkt0AL0n5hi\nmqfhhCDRy8XECYHb+nnpbFY6NUCgCgYIKoZIzj0DAQehRANCAAQ9FEpypYjzGeJ4\nG1Hka/S97GtuAy0rE1dHIJ3guz9a1sCncVWaMjs5rwdw+m6YWNF1kK+RYiAZSl51\n+eYoAvMi\n-----END PRIVATE KEY-----"
	redirectUrl = ""
	BundleID    = "com.wiki.Global"
)

// create client_secret
func GetAppleSecret() string {
	token := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": config.Conf().Apple.KeyId,
		},
		Claims: jwt.MapClaims{
			"iss": config.Conf().Apple.TeamId,
			"iat": time.Now().Unix(),
			// constraint: exp - iat <= 180 days
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"aud": "https://appleid.apple.com",
			"sub": BundleID,
		},
		Method: jwt.SigningMethodES256,
	}

	ecdsaKey, _ := AuthKeyFromBytes([]byte(secret))
	ss, _ := token.SignedString(ecdsaKey)
	return ss
}

// create private key for jwt sign
func AuthKeyFromBytes(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		pp.Println("token: AuthKey must be a valid .p8 PEM file")
		return nil, errors.New("token: AuthKey must be a valid .p8 PEM file")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		pp.Println("token: AuthKey must be of type ecdsa.PrivateKey")
		return nil, errors.New("token: AuthKey must be of type ecdsa.PrivateKey")
	}

	return pkey, nil
}

// do http request
func HttpRequest(method, addr string, params map[string]string) ([]byte, int, error) {
	form := url.Values{}
	for k, v := range params {
		form.Set(k, v)
	}

	var request *http.Request
	var err error
	if request, err = http.NewRequest(method, addr, strings.NewReader(form.Encode())); err != nil {
		return nil, 0, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); nil != err {
		return nil, 0, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	return data, response.StatusCode, nil
}

func Verification(code string) (bool, interface{}, error) {

	data, status, err := HttpRequest("POST", "https://appleid.apple.com/auth/token", map[string]string{
		"client_id":     BundleID,
		"client_secret": GetAppleSecret(),
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  redirectUrl,
	})

	if err != nil {
		return false, "", errors.New("500")
	}
	var result Result
	fmt.Printf("结果：%s\n", data)
	pp.Println(status)
	pp.Println(err)
	_ = json.Unmarshal(data, &result)
	if status != 200 {
		return false, "", errors.New("500")
	}

	res, err := ParseToken(result.IdToken)
	if err != nil {
		return false, "", err
	}
	return true, res["sub"], nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			pp.Println("Unexpected signing method")
			return nil, errors.New("500")
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		pp.Println("解析成功")
		pp.Println(claims)
		return claims, nil
	}
	pp.Println("解析失败")
	return nil, err
}
