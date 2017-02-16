package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CensorReader struct {
	rd io.Reader
}

func (c CensorReader) Read(p []byte) (int, error) {
	n, err := c.rd.Read(p)
	s := bytes.Replace(p, []byte("Snowden"), []byte("*******"), -1)
	s = bytes.Replace(s, []byte("https://"), []byte(" http://"), -1)
	copy(p, s)
	return n, err
}

func copyHeader(dst, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dst.Set(key, value)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := client.Do(req)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "unable to carry out request: %v", err)
		return
	}
	defer response.Body.Close()

	fmt.Printf("%v %v %v\n", r.Method, r.URL, response.Status)

	copyHeader(w.Header(), response.Header)

	w.WriteHeader(response.StatusCode)
	ct := response.Header.Get("content-type")
	if strings.Contains(ct, "text/html") {
		io.Copy(w, CensorReader{response.Body})
	} else {
		io.Copy(w, response.Body)
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
