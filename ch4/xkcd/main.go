// Tool main allows to create offline index of XKCD comics strip
// and allow to query it and print URL by provided term
// from command-line
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	fmt.Println("Last comic number is %d", maxNum)
	var memIndex []*XKCD
	memIndex = append(memIndex, comic)
	for i := 2; i <= 5; i++ {
		comic, err = getOneComicJSON(i)
		memIndex = append(memIndex, comic)
	} // go over all comics and create an in-memory slice
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
	return comics, err
}

func main() {
	//	comic, err := getCurrentComicJSON()
	//	fmt.Println(comic, err)
	/* 	comics, err := buildIndex()
	   	if err == nil {
	   		for _, comic := range comics {
	   			fmt.Printf("Title \"%s\", image URL : \"%s\" \n", comic.SafeTitle, comic.Img)
	   		}
	   	}
	*/ // saveXKCDIndex("test.index", comics)
	comics, err := loadXKCDIndex("test.index")
	if err == nil {
		for _, comic := range comics {
			fmt.Printf("Title \"%s\", image URL : \"%s\" \n", comic.SafeTitle, comic.Img)
		}
	}

}