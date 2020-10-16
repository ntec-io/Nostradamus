package fifaindex

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type link string

type FifaDate struct {
	Date string `selector:"div > div > nav > ol > li > div > a"`
}

func GetTimeId(t time.Time) int {
	return 0
}

func GetAllDateIDs() (m map[time.Time]string) {
	m = make(map[time.Time]string)
	for i := 5; i <= 21; i++ {
		for k, v := range getFifaDateIDs(strconv.Itoa(i)) {
			m[k] = v
		}
	}

	return
}

func getFifaDateIDs(fifa string) (m map[time.Time]string) {
	m = make(map[time.Time]string)
	layout := "Jan. 2, 2006"

	attributes = make(map[string]string)
	c := colly.NewCollector()

	c.IgnoreRobotsTxt = true

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "/players/top/fifa") {

			s := strings.ReplaceAll(e.Text, "Sept", "Sep")

			date, err := time.Parse(layout, s)
			if err == nil {
				m[date] = strings.Trim(strings.ReplaceAll(e.Attr("href"), "/players/top/", ""), "/")
			}
		} else if strings.Contains(e.Attr("href"), "/players/top/") {
			s := strings.ReplaceAll(e.Text, "Sept", "Sep")

			date, err := time.Parse(layout, s)

			if err == nil {
				m[date] = "fifa" + fifa
			}
		}
	})

	c.Visit("https://www.fifaindex.com/players/top/fifa" + fifa + "/")
	c.Wait()

	max := 0

	for _, v := range m {
		x, _ := strconv.Atoi(strings.Trim(strings.ReplaceAll(v, "fifa"+fmt.Sprint(fifa)+"_", ""), "/"))
		if max < x {
			max = x
		}
	}

	for i, v := range m {
		if v == "fifa"+fifa {
			m[i] = "fifa" + fifa + "_" + fmt.Sprint(max+1)
			break
		}
	}

	return m
}
