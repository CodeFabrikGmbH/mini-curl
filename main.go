package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/alexflint/go-arg"
	"io/ioutil"
	"net/http"
	"os"
)

type args struct {
	Username string `arg:"-u" help:"optional: when set http basic authorization is added to the request header"`
	Password string `arg:"-p" help:"optional: needed for http basic authorization"`
	Body     string `arg:"-b" help:"optional: is send as post body"`
	Url      string `arg:"positional,required" help:"url to request"`
}

func (args) Description() string {
	return "Perform an http call"
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func main() {
	var args args
	_ = arg.MustParse(&args)

	client := http.Client{}

	method := "GET"
	if len(args.Body) != 0 {
		method = "POST"
	}

	req, err := http.NewRequest(method, args.Url, bytes.NewBuffer([]byte(args.Body)))
	handleError(err)

	if len(args.Username) != 0 && len(args.Password) != 0 {
		authString := base64.StdEncoding.EncodeToString([]byte(args.Username + ":" + args.Password))
		req.Header.Add("Authorization", "Basic "+authString)
	}

	response, err := client.Do(req)
	handleError(err)
	defer func() {
		_ = response.Body.Close()
	}()

	all, err := ioutil.ReadAll(response.Body)
	handleError(err)

	fmt.Println(string(all))

	if response.StatusCode == http.StatusOK {
		os.Exit(0)
	} else {
		fmt.Printf("status code: %d", response.StatusCode)
		os.Exit(-1)
	}
}
