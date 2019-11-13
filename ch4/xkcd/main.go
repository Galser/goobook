// Tool main allows to create offline index of XKCD comics strip
// and allow to query it and print URL by provided term
// from command-line
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/gookit/color"
)

/*

XKCD explanation :

If you want to fetch comics and metadata automatically,
you can use the JSON interface. The URLs look like this:

http://xkcd.com/info.0.json (current comic)

or:

http://xkcd.com/614/info.0.json (comic #614)

Those files contain, in a plaintext and easily-parsed format: comic titles,
URLs, post dates, transcripts (when available), and other metadata.



Example of one JSON

{
   "month":"4",
   "num":571,
   "link":"",
   "year":"2009",
   "news":"",
   "safe_title":"Can't Sleep",
   "transcript":"[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}",
   "alt":"If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.",
   "img":"https://imgs.xkcd.com/comics/cant_sleep.png",
   "title":"Can't Sleep",
   "day":"20"
}
*/
const BaseURL = "http://xkcd.com/"
const InfoJSON = "info.0.json"
const indexFileName = "xkcd.index"

type XKCD struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Day        string
}

/* type XKCDIndex struct {
    Comics []XKCD `json:"xkcd"`
}
*/

func getCurrentComicJSON() (*XKCD, error) {
	resp, err := http.Get(BaseURL + InfoJSON)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Getting current comic failed: %s", resp.Status)
	}

	var result XKCD
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// Gets one JSON for the comic specified by "Num"
// Note : XKCD starts comic numbering from 1
func getOneComicJSON(Num int) (*XKCD, error) {
	resp, err := http.Get(fmt.Sprintf("%s%d/%s", BaseURL, Num, InfoJSON))
	//fmt.Println(BaseURL + string(Num) + "/" + InfoJSON)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Getting comic %d failed: %s", Num, resp.Status)
	}

	var result XKCD
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func buildIndex() ([]*XKCD, error) {
	// get current comic
	comic, err := getCurrentComicJSON()
	if err != nil {
		return nil, err
	}
	maxNum := comic.Num
	fmt.Println("Last comic number is : ", maxNum)
	fmt.Print("Building index (one # - is 5% ) --> [")
	var memIndex []*XKCD
	memIndex = append(memIndex, comic)
	progress := 0
	for i := 2; i <= maxNum; i++ {
		if int(math.Ceil(float64(i*20/maxNum))) > progress {
			progress++
			fmt.Print("#")
		}
		comic, err = getOneComicJSON(i)
		memIndex = append(memIndex, comic)
	} // go over all comics and create an in-memory slice
	fmt.Println("]")
	return memIndex, nil
}

//  saveXKCDIndex - reads from site all possible
// JSONs and saving them into file
func saveXKCDIndex(filename string, Comics []*XKCD) error {
	comicsJSON, err := json.Marshal(Comics)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, comicsJSON, 0644)
	return err
}

func loadXKCDIndex(filename string) ([]*XKCD, error) {
	comicsJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var comics []*XKCD
	err = json.Unmarshal(comicsJSON, &comics)
	if err == nil {
		fmt.Println("Successfully loaded index from disk")
	}
	return comics, err
}

func searchTerm(term string, Comics []*XKCD) {
	yellow := color.FgYellow.Render
	for _, comic := range Comics {
		if comic != nil && comic.SafeTitle != "" && comic.Img != "" && comic.Transcript != "" {
			if strings.Contains(comic.Transcript, term) {
				s := strings.ReplaceAll(comic.Transcript, term, yellow(term))
				fmt.Printf("Title \"%s\", \n \" %s \" \nimage URL : \"%s\" \n", comic.SafeTitle, s, comic.Img)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "TERM")
		fmt.Println(" -- will search for TERM in offline XKCD index")
		return
	}
	comicsInIndexNow := 0
	comics, err := loadXKCDIndex(indexFileName)
	if err != nil {
		fmt.Printf("Index file could not be opened ('%s'), going to create one\n", indexFileName)
		comics, err = buildIndex()
		_ = saveXKCDIndex(indexFileName, comics)
		// we don't care about the error above for now
	} else {
		comicsInIndexNow = len(comics)
		lastcomic, err := getCurrentComicJSON()
		if err != nil {
			fmt.Println("Failed to read the current/last comic")
		} else {
			maxNum := lastcomic.Num
			if maxNum > comicsInIndexNow {
				fmt.Printf("Index need to be updated...")
			}
		}
	}
	comicsInIndexNow = len(comics)
	fmt.Println("Number of comics in index right now : ", comicsInIndexNow)
	fmt.Println("Looking for '", os.Args[1], "' in offline index \n")
	searchTerm(os.Args[1], comics)

}

/*

Output example (it comes with colors ):

go run main.go Cisco
Successfully loaded index from disk
Index need to be updated...Number of comics in index right now :  2224
Looking for ' Cisco ' in offline index

Title "Missed Connections",
 " ((The page is set up like the missed connections area of Craigslist, with a list of messages from an individual to a person they weren't able to communicate with at the time.))
Personals > Missed Connections

You: Clinging to hood of your stolen wienermobile, trying to reach into engine to unstick throttle
Me: Screaming, diving out of the way

You: Vaguely human silhouette
Me: At bottom of wishing well with harpoon gun

You: Confused UDP packet
Me: Cisco router in 45.170
16 block

You: Baddest fuckin' Juggalo at Violent J's party
Me: Nancy Pelosi (D-Ca)

You: Getting married to me
Me: Also getting married, but distracted by my phone

You: Cute boy on corner of 4th & Main, 5'11, 169lbs, social security number 078-05-1120, pockets contained $2.09 in change, keys, and a condom. Retinal scan attached
Me: Driving street view van

You: George Herman "Babe" Ruth
Me: Fellow Time Lord. Saw your tardis on third moon of <<Sentence cuts off, partially obscured by bottom of panel>>

{{Title text: The Street View van isn't going to find out anything Google won't already know from reading my email.}} "
image URL : "https://imgs.xkcd.com/comics/missed_connections.png"
*/
