package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ejilay/draftjs"
	"github.com/urfave/cli"
)

// build command: go build -o draftjs-cli cli/main.go
// run command: ./draftjs-cli '{"entityMap":{},"blocks":[{"key":"4g603","text":"dasdasdasdsadsaывфвыфв","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":22,"style":"BOLD"}],"entityRanges":[],"data":{}}]}'
// arg 1: stringified json
// output: html
func main() {
	var file string

	app := cli.NewApp()
	app.Name = "draftjs-cli"
	app.Usage = "convert draftjs json to html"

	app.Flags = append(app.Flags, cli.StringFlag{
		Name:        "file",
		Usage:       "file json draftjs data",
		Destination: &file,
	})

	app.Action = func(c *cli.Context) error {
		contentState := draftjs.ContentState{}
		if file != "" {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()
			// from stringify json

			if err := json.NewDecoder(f).Decode(&contentState); err != nil {
				return err
			}
		} else {
			draftState := c.Args().First()
			if err := json.Unmarshal([]byte(draftState), &contentState); err != nil {
				return err
			}
		}

		if contentState.Blocks == nil {
			return fmt.Errorf("no blocks found")
		}

		config := draftjs.NewDefaultConfig()
		s := draftjs.Render(&contentState, config)
		fmt.Println("html>>>", s)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
