package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	createPublicKey()
	defer cleanupPublicKey()

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VybmFtZSI6InJvb3QifQ.BlHzcFFxr__XOuA8znXsV-5qKTGZGgNZaP8Uy87TEFrTQrC37v6EuvoOQMusDrI-BOlF4fwS_v--8j84PJV6rEcV0Uxx2o1MGfKI7lQ-ckG5ic2MtGEChJDFh5Xh1UfMHikrD-Uh8VgmltvuYVHWX35PBHHQi-w8yp483zox17_H-A2J_QilQ0MF-17IA2aIFUCOUdu0wboq6KS7tsxr9zKYBldC9sWRKM-KcOdB2NwDMqkKtdZRS0mZYuHnvmOb75V3lhvsGRsHfNHkpE_MsS6ZVEFUzWt2XR4jtTNPUjb62706M4twS4wVKjiKRrLQ-5uwDkb_FX7QSVQtONKTVA"
	user := decodeJWT(token)
	if user.Username != "root" {
		t.Error("Username must be equal to root")
	}

	if user.Role != "admin" {
		t.Error("Role must be equal to admin")
	}
}

func createPublicKey() {
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

	err := ioutil.WriteFile(publicKeyPath, []byte(publicKey), 0644)
	if err != nil {
		panic(err)
	}
}

func cleanupPublicKey() {
	err := os.Remove(publicKeyPath)
	if err != nil {
		panic(err)
	}
}
