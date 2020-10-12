package fifaindex

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func scrape(url string) (data []byte, err error) {
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	c.Limit(&colly.LimitRule{
		DomainGlob: "www.fifaindex.com/*",
		Delay:      time.Second,
	})

	c.OnHTML("div.col-lg-8", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.Visit(url)
	c.Wait()

	fmt.Println(c)

	return
}

func GetPlayer(playerName string) (player Player, err error) {
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	c.Limit(&colly.LimitRule{
		DomainGlob: "www.fifaindex.com/*",
		Delay:      time.Second,
	})

	c.OnHTML("div.col-lg-8", func(e *colly.HTMLElement) {
		fmt.Println(e)

	})

	c.Visit("https://www.fifaindex.com/player/234642/%C3%A9douard-mendy/")
	c.Wait()

	fmt.Println(c)

	return
}
