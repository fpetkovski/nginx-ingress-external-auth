package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	proxyPort   = 8000
	servicePort = 80
)

const publicKeyPath = "/tmp/public.pem"

// Create a structure to define the proxy functionality.
type Proxy struct{}

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Forward the HTTP request to the destination service.
	res, duration, err := p.forwardRequest(req)

	// Notify the client if there was an error while forwarding the request.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// If the request was forwarded successfully, write the response back to
	// the client.
	p.writeResponse(w, res)

	// Print request and response statistics.
	p.printStats(req, res, duration)
}

func main() {
	// Listen on the predefined proxy port.
	http.ListenAndServe(fmt.Sprintf(":%d", proxyPort), &Proxy{})
}

func (p *Proxy) forwardRequest(req *http.Request) (*http.Response, time.Duration, error) {
	// Prepare the destination endpoint to forward the request to.
	proxyUrl := fmt.Sprintf("http://127.0.0.1:%d%s", servicePort, req.RequestURI)

	// Print the original URL and the proxied request URL.
	fmt.Printf("Original URL: http://%s:%d%s\n", req.Host, servicePort, req.RequestURI)
	fmt.Printf("Proxy URL: %s\n", proxyUrl)

	// Create an HTTP client and a proxy request based on the original request.
	httpClient := http.Client{}
	proxyReq, err := http.NewRequest(req.Method, proxyUrl, req.Body)

	user := decodeJWT(req.Header.Get("X-JWT"))
	proxyReq.Header.Add("UserID", user.Username)
	proxyReq.Header.Add("UserRole", user.Role)

	// Capture the duration while making a request to the destination service.
	start := time.Now()
	res, err := httpClient.Do(proxyReq)
	duration := time.Since(start)

	// Return the response, the request duration, and the error.
	return res, duration, err
}

func (p *Proxy) writeResponse(w http.ResponseWriter, res *http.Response) {
	// Copy all the header values from the response.
	for name, values := range res.Header {
		w.Header()[name] = values
	}

	// Set a special header to notify that the proxy actually serviced the request.
	w.Header().Set("Server", "amazing-proxy")

	// Set the status code returned by the destination service.
	w.WriteHeader(res.StatusCode)

	// Copy the contents from the response body.
	io.Copy(w, res.Body)

	// Finish the request.
	res.Body.Close()
}

func (p *Proxy) printStats(req *http.Request, res *http.Response, duration time.Duration) {
	fmt.Printf("Request Duration: %v\n", duration)
	fmt.Printf("Request Size: %d\n", req.ContentLength)
	fmt.Printf("Response Size: %d\n", res.ContentLength)
	fmt.Printf("Response Status: %d\n\n", res.StatusCode)
}

func decodeJWT(jwtToken string) User {
	publicKey := parsePublicKey()

	parts := strings.Split(jwtToken, ".")
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	if err != nil {
		panic("Invalid JWT token")
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	user := User{
		Username: fmt.Sprintf("%v", claims["username"]),
		Role:     fmt.Sprintf("%v", claims["role"]),
	}

	return user
}

func parsePublicKey() *rsa.PublicKey {
	key, _ := ioutil.ReadFile(publicKeyPath)
	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(key)

	return publicKey
}
