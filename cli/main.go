package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ejilay/draftjs"
	"github.com/urfave/cli"
)

// build command: go build -o draftjs-cli cli/main.go
// run command: ./draftjs-cli '{"entityMap":{},"blocks":[{"key":"4g603","text":"dasdasdasdsadsaывфвыфв","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":22,"style":"BOLD"}],"entityRanges":[],"data":{}}]}'
// arg 1: stringified json
// output: html
func main() {
	app := cli.NewApp()
	app.Name = "draftjs-cli"
	app.Usage = "convert draftjs json to html"
	app.Action = func(c *cli.Context) error {
		draftState := c.Args().First()
		contentState := draftjs.ContentState{}
		if err := json.Unmarshal([]byte(draftState), &contentState); err != nil {
			fmt.Println(err)
			return err
		}
		config := draftjs.NewDefaultConfig()
		s := draftjs.Render(&contentState, config)
		fmt.Println(s)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
