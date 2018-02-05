package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
)

type customer struct {
	Name    string
	Address string
}

func main() {

	c:= customer{Name: "a", Address: "ca"}
	b, err:= json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, _ := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(b))

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
}