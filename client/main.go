package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/jcbritobr/cstodo/model"
	"github.com/olekukonko/tablewriter"
)

var (
	listFlag   = flag.Bool("list", false, "Lists all data from kvstore")
	insertFlag = flag.Bool("insert", false, "Inserts data in kvstore. Needs data flag")
	dataFlag   = flag.String("data", `{"message":"default"}`, "Pass data as json to kvstore")
	dudFlag    = flag.Bool("dud", false, "Switches the Done field value. Needs data flag")
)

func sendRequest(protocol, url string, body io.Reader) (string, error) {
	req, err := http.NewRequest(protocol, url, body)
	if err != nil {
		return "", err
	}

	if body != nil {
		req.Header.Add("ApplicationType", "text/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad request: %v\n", resp.StatusCode)
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Corrupted data %v\n", err)
	}

	return string(rbody), nil
}

func insert(data io.Reader) {
	res, err := sendRequest("POST", "http://localhost:8080/insert", data)
	checkPanic(err)
	fmt.Println(res)
}

func list() {
	data, err := sendRequest("GET", "http://localhost:8080/list", nil)
	checkPanic(err)

	var buffer map[string]model.Item
	err = json.Unmarshal([]byte(data), &buffer)
	checkPanic(err)

	var lines [][]string
	for i, data := range buffer {
		done := "false"
		if data.Done {
			done = "true"
		}
		lines = append(lines, []string{i, data.Title, data.Description, done})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Uuid", "Title", "Description", "Done"})
	for _, v := range lines {
		table.Append(v)
	}

	table.Render()
}

func doneUndone(data io.Reader) {
	res, err := sendRequest("POST", "http://localhost:8080/doneundone", data)
	checkPanic(err)
	fmt.Println(res)
}

func main() {
	flag.Parse()

	if *listFlag {
		list()
	}

	if *insertFlag {
		insert(strings.NewReader(*dataFlag))
	}

	if *dudFlag {
		doneUndone(strings.NewReader(*dataFlag))
	}
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
