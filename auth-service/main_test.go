package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const publicKeyPath  = "/tmp/public.pem"

func TestGenerateJwtToken(t *testing.T) {
	createRSAPair()
	defer cleanupRSAPair()

	jwtToken := generateJwtToken("root", "admin")
	validateToken(t, jwtToken)
}

func validateToken(t *testing.T, jwtToken string) {
	publicKey := parsePublicKey()

	parts := strings.Split(jwtToken, ".")
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	if err != nil {
		t.Error("The token cannot be validated by the public key")
	}
}

func parsePublicKey() *rsa.PublicKey {
	key, _ := ioutil.ReadFile(publicKeyPath)
	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(key)

	return publicKey
}

func createRSAPair() {
	privateKey :=
		`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAr3FOsr3tXoWj/Csj+YJLiBBneUjIWUjV12nchoLteL4xWSnJ
9rkrkx9gF88HqMZXoFxKrT2tj1jN51KAqDnG3gy2kBpYz4vy2Rq8dQiTlYhNHcgQ
bpIG2sOTTHbnjHTqxugXSVoMTyyJy2G0ZupkxiU720XldIXgQICA0sFzqwurTeyQ
XVbYmLao4rTVonlPIQKvJgCw8YBWOOdXTZjQrnWJRU+uNKJsq/RQJLzjzvKQ+tPV
7rVARCP6WjwXWk2kP3W/sgj10vNe/mwibh4DhAK02ZFo2OFgI8xm4KXCW9UOctPW
ZexqaV7UBKLUKWhsio3CfRdwkjDRceVU3ggrBQIDAQABAoIBABWxxTGFGt0dWXHN
Q92OpxhkLudogJ3Jy+efR426kvLjUebcrAS+UQ0YTCGlkCTmod9KilFx1wxqxstc
AFzNCDJdEBYxq9m+TIFcNQGj1dwfxqlwG9gQZpiWpphF+8v6iS2fdYG0iOEheMiV
hGFVirCV2hj7Q6xtAZX4TfXBxjPYQRntJDgJYdbileH3fnLxSQuxdw3DL+AYP2Bw
eWmHs3RVJ6L0q6NEwoPLrpKtZU5l0vhDRhDMjrv6ICTYjczRqwH+lewBw3GpXppP
LZGreL/FKgIVx9sNO2FunlEAuB4z26aNUuhUhTMBvz+HHUyWhuggrnsg063oWotG
ElK0GAECgYEA3LAPFCFrdnpqwrkj3ibBPGgSqxg/hdcXpeMZbB3EFeur0vOwwri2
ywty16qE4IUeIJaIwkB1gUrL5Os5j91ShTV5lUpGclty6dBHoREEtzUapCSVvXtL
eWbJ0S4vcxb2z1lMciI6nkJky5PfZCVWCoUcaP6KBraiBqClvYOGUcECgYEAy4Pk
HU1uvFQ81ppcIt+Mu6GbxmpuGY/q2oRO/1c4JESp/1mSVBwhBjCcLv1AmwzToz9z
0RaJja2Ao/9JthIZc9XgHfC9urQAUrHgh2br/FkUvVW3YAfBKkFelLK+zakR/pTS
p4+G6fzOfquymxz4Pi2jLaeDmGWqUuqAHAEYokUCgYB9PimAiirblUPLeJijdakK
qCGYGe3K/jO8gDKoSghDTHk6AfSZvYx3lOq6/FnmtYVQhz3byAsnshQeuWP0gm1X
je6PTBTIx59ilEJiZS8g7jFNYDney/8cSbpVTXm+PhUZvZsF1ukfcZyUcDpCMnIv
DDYAXBxnDPTNABSvhdoPwQKBgB7KHmncfCNb7zRceBICli0Q3xteoLeXUWWr3LO+
w0yhYsKyD3RQKSLhmc92Gx8aCq7I+8GnUjowBKVLCyDTjiw7MEP3Vwz3DJF2Pcze
Yld0NrIKVMrfgXbeGuwOOtWsfX9xjokxKq0dxTPe0A+ti1UE3IocrMkSoHkY5zbV
Z5+tAoGAcZWSrPR3Xzb92wuDJ05kTqvg2/rr8WI5acZQLinVvOYeHUOsC60BpTgq
+fugPct7OuVZ8DcspfCe0WO2nnH00mKM82a7EKdcGCUrZci9ZwUAfyw8m9BED8w+
ffKfKwJNcQHOi+I1lwGnutfN4eukpnx9gRBWXpAWrlZF53TOkfY=
-----END RSA PRIVATE KEY-----
`

	err := ioutil.WriteFile(privateKeyPath, []byte(privateKey), 0644)
	if err != nil {
		panic(err)
	}

	publicKey :=
`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr3FOsr3tXoWj/Csj+YJL
iBBneUjIWUjV12nchoLteL4xWSnJ9rkrkx9gF88HqMZXoFxKrT2tj1jN51KAqDnG
3gy2kBpYz4vy2Rq8dQiTlYhNHcgQbpIG2sOTTHbnjHTqxugXSVoMTyyJy2G0Zupk
xiU720XldIXgQICA0sFzqwurTeyQXVbYmLao4rTVonlPIQKvJgCw8YBWOOdXTZjQ
rnWJRU+uNKJsq/RQJLzjzvKQ+tPV7rVARCP6WjwXWk2kP3W/sgj10vNe/mwibh4D
hAK02ZFo2OFgI8xm4KXCW9UOctPWZexqaV7UBKLUKWhsio3CfRdwkjDRceVU3ggr
BQIDAQAB
-----END PUBLIC KEY-----
`

	err = ioutil.WriteFile(publicKeyPath, []byte(publicKey), 0644)
	if err != nil {
		panic(err)
	}
}

func cleanupRSAPair() {
	err := os.Remove(privateKeyPath)
	if err != nil {
		panic(err)
	}
	err = os.Remove(publicKeyPath)
	if err != nil {
		panic(err)
	}
}
