# XKCD tool
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.

# TODO 
- Compare the file index amount of comics on-disk with the one on-line,  and update index when there a new one

# JSON 

XKCD explanation :

If you want to fetch comics and metadata automatically,
you can use the JSON interface. The URLs look like this:

http://xkcd.com/info.0.json (current comic)

or:

http://xkcd.com/614/info.0.json (comic #614)

Those files contain, in a plaintext and easily-parsed format: comic titles,
URLs, post dates, transcripts (when available), and other metadata.


# Ideas

- The current comic can provide us a max available comic at the moment :  
.e.g. `https://xkcd.com/info.0.json` :
```json
{ 
   "month":"10",
   "num":2220,
   "link":"",
   "year":"2019",
   "news":"",
   "safe_title":"Imagine Going Back in Time",
   "transcript":"",
   "alt":"I wonder what the trendy adults in 2019 who are too cool for Pokemon will be into. Probably Digimon!",
   "img":"https://imgs.xkcd.com/comics/imagine_going_back_in_time.png",
   "title":"Imagine Going Back in Time",
   "day":"25"
}
```

so..the last one is 2220 

