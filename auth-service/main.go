/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strconv"
)

const (
	privateKeyPath = "/tmp/private.pem"
)

// Sample authentication service returning several HTTP headers in response
func main() {
	http.HandleFunc("/", verifyUser)
	http.HandleFunc("/login", handleLogin)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func generateJwtToken(username string, role string) string {
	key, _ := ioutil.ReadFile(privateKeyPath)
	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM(key)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
		"role":     role,
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func verifyUser(w http.ResponseWriter, r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))

	authCookie, err := r.Cookie("auth")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtToken := generateJwtToken(authCookie.Value, "admin")
	w.Header().Add("X-JWT", jwtToken)

	fmt.Fprint(w, "ok")
	fmt.Println("Authenticated")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	authCookie := &http.Cookie{
		Name:     "auth",
		Value:    strconv.Itoa(rand.Int()),
		Domain:   "kube.local",
		HttpOnly: false,
	}
	http.SetCookie(w, authCookie)

	redirectUrl := r.URL.Query()["rd"][0]
	http.Redirect(w, r, redirectUrl, http.StatusFound)
}
