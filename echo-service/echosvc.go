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
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))
	fmt.Fprintf(w, "UserID: %s, UserRole: %s", r.Header.Get("UserID"), r.Header.Get("UserRole"))
}

// Sample  "echo" service displaying UserID and UserRole HTTP request headers
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
