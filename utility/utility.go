package utility

import "net/http"
import "io/ioutil"
import "bytes"

func Request(url string ,data []byte) []byte {

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	// if resp.Status != 200{
	// 	fmt:Println("status not ok")
	// }
    body, _ := ioutil.ReadAll(resp.Body)
	return body
}

