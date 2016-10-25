package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var (
	addCommand = cli.Command{
		Name:  "add",
		Usage: "add some cool",
		Action: func(c *cli.Context) {
			cool := c.Args().First()
			authToken := c.GlobalString("auth")
			endpoint := c.GlobalString("endpoint")
			AddCool(endpoint, authToken, cool)
		},
	}
)

type AddRequest struct {
	AuthToken string `json:"authToken"`
	CoolThing string `json:"coolThing"`
}

type AddResponse struct {
	Success bool `json:"success"`
}

func AddCool(endpoint, auth, cool string) {
	jsonB := []byte(fmt.Sprintf(`{"authToken": "%s", "coolThing": "%s"}`, auth, cool))
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonB))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	//   fmt.Println(err.Error())
	// }
	// fmt.Println(string(b))
	dec := json.NewDecoder(resp.Body)
	var data AddResponse
	err = dec.Decode(&data)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	if data.Success {
		fmt.Printf("Success :)\n")
	} else {
		fmt.Printf("Failure :(\n")
	}
}

func main() {
	app := cli.NewApp()

	app.Author = "wercker"
	app.Email = "pleasemailus@wercker.com"
	app.Name = "cooltime"
	app.Usage = "add cool stuff"
	app.Version = FullVersion()
	app.Commands = []cli.Command{
		addCommand,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "auth"},
		cli.StringFlag{Name: "endpoint"},
	}

	app.Run(os.Args)
}
