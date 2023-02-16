package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/urfave/cli/v2"
)

func main() {

	translate := Translation{
		From: "",
		To:   "",
		Word: "",
		Res:  "",
	}

	app := &cli.App{
		Name:      "i18next-helper",
		Usage:     "json builder for i18next",
		UsageText: "from/to/word --from <lang> --to <lang> --word <word>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases:     []string{"f"},
				Name:        "from",
				Destination: &translate.From,
				Required:    true,
			},
			&cli.StringFlag{
				Aliases:     []string{"t"},
				Name:        "to",
				Destination: &translate.To,
				Required:    true,
			},
			&cli.StringFlag{
				Aliases:     []string{"w"},
				Name:        "word",
				Destination: &translate.Word,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			translate.translate()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	//todo parse existing json structure
	//todo add available flags to cli help flag
	file, _ := json.MarshalIndent(translate.Res, "", " ")
	_ = ioutil.WriteFile("output/test.json", file, 0644)
}

type Translation struct {
	From string
	To   string
	Word string
	Res  string
}

func (translation *Translation) translate() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://translate.google.com/?sl=`+translation.From+`&tl=`+translation.To+`&text=`+translation.Word+`&op=translate`),
		chromedp.Text(`jCAhz ChMk0b`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	translation.Res = res
	log.Println(translation.Res)
}
