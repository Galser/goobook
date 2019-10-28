// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a table of GitHub issues matching the search terms.
// and in some time criteria
// - less than a month old,
// - less than a year old
// - and more than a year old

// added color-coding

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"mycode/ch4/github"

	"github.com/gookit/color"
)

//!+
func main() {
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	red := color.FgRed.Render
	yellow := color.FgYellow.Render
	green := color.FgGreen.Render
	defColor := green
	fmt.Printf("%s  (>1 year), \n%s (>1 month) \n%s - fresh one\n", red("=== very old issue"), yellow("=== slightly old"), green("=== fresh"))

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.Before(yearAgo) {
			defColor = green
		} else {
			if item.CreatedAt.Before(monthAgo) {
				defColor = yellow
			} else {
				defColor = red
			}
		}
		fmt.Printf("%s #%-5d %9.9s %.55s\n",
			defColor("==="), item.Number, item.User.Login, item.Title)
	}
}

//!-

/*
//!+textoutput
go run ch4/issues/main.go repo:golang/go is:open json decoder
42 issues:
#33416   bserdar encoding/json: This CL adds Decoder.InternKeys
#34647 babolivie encoding/json: fix byte counter increments when using d
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#32779       rsc proposal: encoding/json: memoize strings during decode?
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#28923     mvdan encoding/json: speed up the decoding scanner
#11046     kurin encoding/json: Decoder internally buffers full input
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
#30301     zelch encoding/xml: option to treat unknown fields as an erro
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#31701    lr1980 encoding/json: second decode after error impossible
#31789  mgritter encoding/json: provide a way to limit recursion depth
#16212 josharian encoding/json: do all reflect work before decoding
#33714    flimzy proposal: encoding/json: Opt-in for true streaming supp
#22480     blixt proposal: encoding/json: add omitnil option
#28189     adnsv encoding/json: confusing errors when unmarshaling custo
#5901        rsc encoding/json: allow override type marshaling
#7872  extempora encoding/json: Encoder internally buffers full output
#14750 cyberphon encoding/json: parser ignores the case of member names
#28143    arp242 proposal: encoding/json: add "readonly" tag
#30701 LouAdrien encoding/json: ignore tag "-" not working on embedded s
#22752  buyology proposal: encoding/json: add access to the underlying d
#27179  lavalamp encoding/json: no way to preserve the order of map keys
#20754       rsc encoding/xml: unmarshal only processes first XML elemen
#20528  jvshahid net/http: connection reuse does not work happily with n
#22816 ganelon13 encoding/json: include field name in unmarshal error me
#21823  243083df encoding/xml: very low performance in xml parser

//!-textoutput
*/
