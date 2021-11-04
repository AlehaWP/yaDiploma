package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "encoding/json"
)

type uR struct {
	Log string `json:"login"`
	Pas string `json:"password"`
}

func printResult(body io.Reader, r *http.Response) {
	text, err := io.ReadAll(r.Body)
	if err != nil {
		print(err)
	}
	defer r.Body.Close()
	fmt.Printf("%s\n", r.Header)
	fmt.Printf("%s\n", text)
	fmt.Printf("%d\n", r.StatusCode)
}

func testSign(t string, log string, pas string) {
	u := uR{
		Log: log,
		Pas: pas,
	}
	reqBody, err := json.Marshal(&u)
	if err != nil {
		print(err)
	}
	a := "http://localhost:8080/api/user/" + t
	makeRequest(a, reqBody)
	makeZipRequest(a, reqBody)
}

func makeRequest(a string, b []byte) {
	r, err := http.Post(a, "aplication/json", bytes.NewBuffer(b)) //bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	printResult(r.Body, r)
}

func makeZipRequest(a string, reqBody []byte) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	gz.Write(reqBody)
	gz.Flush()
	gz.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", a, bytes.NewReader(b.Bytes()))
	req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Content-Type", "aplication/json")

	r, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	gzr, err := gzip.NewReader(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	printResult(gzr, r)
}

func noRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func makeGet(url string) {
	client := &http.Client{
		CheckRedirect: noRedirect,
	}
	req, _ := http.NewRequest("GET", url, nil)
	//req, _ := http.NewRequest("GET", "http://localhost:8080/5ad2d1e92b0d271798b22621a54043e6", nil)
	// response, err := http.Get("http://localhost:8080/04f51bcd17361670c1dc6d94cbbd0efe")

	req.AddCookie(&http.Cookie{
		Name:  "UserID",
		Value: "877c34328620072ca769e960d250ac9e",
	})

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	// text,_ := io.ReadAll(response.Header.Get("Location"))
	text := response.Header
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Printf("%s\n", text)
	fmt.Printf("%s\n", body)
	fmt.Printf("%d\n", response.StatusCode)
}

func makeGetPing() {
	client := &http.Client{
		CheckRedirect: noRedirect,
	}
	//req, _ := http.NewRequest("GET", "http://localhost:8080/14afc95e687fa093f0edfa25de0766cd", nil)
	req, _ := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	// response, err := http.Get("http://localhost:8080/04f51bcd17361670c1dc6d94cbbd0efe")

	req.AddCookie(&http.Cookie{
		Name:  "UserID",
		Value: "dc3b5af8713f9d0c1f2dc708e8b2f038",
	})

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	fmt.Printf("%d\n", response.StatusCode)
}

func main() {
	testSign("register", "Aleha", "123123213")
	testSign("register", "Kartoha", "457457457457")
	testSign("login", "Aleha", "123123213")
	testSign("login", "Kartoha", "457457457457")
	// makeGetPing()

	// makePost()
	// makePostApi()
	// makePostZipApi("www.testZip3.ru/ip23123")
	// makePostZipApi("www.testZip4.ru/ip23123")
	// makePostApiUrls()

	// makeGet("http://localhost:8080/14afc95e687fa093f0edfa25de0766cd")
	// makeGet("http://localhost:8080/b57b4c84086c45334df6a07dfbbf8ab9")
	// makeGetUserUrls()

	// makeDelUserUrls()
}
