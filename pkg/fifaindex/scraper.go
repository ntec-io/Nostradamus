package fifaindex

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

var attributes map[string]string

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
	attributes = make(map[string]string)
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	c.Limit(&colly.LimitRule{
		DomainGlob: "www.fifaindex.com/*",
		Delay:      time.Second,
	})

	c.OnHTML("div.card-body", func(e *colly.HTMLElement) {

		for _, v := range strings.SplitAfter(e.Text, "\n") {
			t := strings.Trim(v, " ")
			sc := strings.Split(t, " ")
			var name string
			for _, sub := range sc[:len(sc)-1] {
				name += sub + " "
			}
			name = strings.Trim(name, " ")
			attributes[name] = strings.Trim(sc[len(sc)-1], "\n")
		}

		//fmt.Println(e)

	})

	c.Visit("https://www.fifaindex.com/player/20801/cristiano-ronaldo/fifa20/")
	c.Wait()

	fmt.Println(c)
	fmt.Println(attributes)

	fmt.Println(AssablePlayer(attributes))

	return
}

func AssablePlayer(attributes map[string]string) (p *Player, err error) {
	p = new(Player)

	// Ball Skills
	p.BallSkills.BallControl, err = strconv.Atoi(attributes["Ball Control"])
	if err != nil {
		return
	}
	p.BallSkills.Dribbling, err = strconv.Atoi(attributes["Dribbling"])
	if err != nil {
		return
	}

	// Goalkeeper
	p.Goalkeeper.Diving, err = strconv.Atoi(attributes["GK Diving"])
	if err != nil {
		return
	}
	p.Goalkeeper.Positioning, err = strconv.Atoi(attributes["GK Positioning"])
	if err != nil {
		return
	}
	p.Goalkeeper.Handling, err = strconv.Atoi(attributes["GK Handling"])
	if err != nil {
		return
	}
	p.Goalkeeper.Kicking, err = strconv.Atoi(attributes["GK Kicking"])
	if err != nil {
		return
	}
	p.Goalkeeper.Reflexes, err = strconv.Atoi(attributes["GK Reflexes"])
	if err != nil {
		return
	}

	// Defence
	p.Defence.Marking, err = strconv.Atoi(attributes["Marking"])
	if err != nil {
		return
	}
	p.Defence.SideTackle, err = strconv.Atoi(attributes["Slide Tackle"])
	if err != nil {
		return
	}
	p.Defence.StandTackle, err = strconv.Atoi(attributes["Stand Tackle"])
	if err != nil {
		return
	}

	// Mental
	p.Mental.Aggression, err = strconv.Atoi(attributes["Aggression"])
	if err != nil {
		return
	}
	p.Mental.Reactions, err = strconv.Atoi(attributes["Reactions"])
	if err != nil {
		return
	}
	p.Mental.AttPosition, err = strconv.Atoi(attributes["Att. Position"])
	if err != nil {
		return
	}
	p.Mental.Interceptions, err = strconv.Atoi(attributes["Interceptions"])
	if err != nil {
		return
	}
	p.Mental.Vision, err = strconv.Atoi(attributes["Vision"])
	if err != nil {
		return
	}
	p.Mental.Composure, err = strconv.Atoi(attributes["Composure"])
	if err != nil {
		return
	}

	// Passing
	p.Passing.Crossing, err = strconv.Atoi(attributes["Crossing"])
	if err != nil {
		return
	}
	p.Passing.LongPass, err = strconv.Atoi(attributes["Long Pass"])
	if err != nil {
		return
	}
	p.Passing.ShortPass, err = strconv.Atoi(attributes["Short Pass"])
	if err != nil {
		return
	}

	// Physical
	p.Physical.Acceleration, err = strconv.Atoi(attributes["Acceleration"])
	if err != nil {
		return
	}
	p.Physical.Stamina, err = strconv.Atoi(attributes["Stamina"])
	if err != nil {
		return
	}
	p.Physical.Strength, err = strconv.Atoi(attributes["Strength"])
	if err != nil {
		return
	}
	p.Physical.Balance, err = strconv.Atoi(attributes["Balance"])
	if err != nil {
		return
	}
	p.Physical.SprintSpeed, err = strconv.Atoi(attributes["Sprint Speed"])
	if err != nil {
		return
	}
	p.Physical.Agility, err = strconv.Atoi(attributes["Agility"])
	if err != nil {
		return
	}
	p.Physical.Jumping, err = strconv.Atoi(attributes["Jumping"])
	if err != nil {
		return
	}

	// Shooting
	p.Shooting.Heading, err = strconv.Atoi(attributes["Heading"])
	if err != nil {
		return
	}
	p.Shooting.ShotPower, err = strconv.Atoi(attributes["Shot Power"])
	if err != nil {
		return
	}
	p.Shooting.Finishing, err = strconv.Atoi(attributes["Finishing"])
	if err != nil {
		return
	}
	p.Shooting.LongShots, err = strconv.Atoi(attributes["Long Shots"])
	if err != nil {
		return
	}
	p.Shooting.Curve, err = strconv.Atoi(attributes["Curve"])
	if err != nil {
		return
	}
	p.Shooting.FKAcc, err = strconv.Atoi(attributes["FK Acc."])
	if err != nil {
		return
	}
	p.Shooting.Penalties, err = strconv.Atoi(attributes["Penalties"])
	if err != nil {
		return
	}
	p.Shooting.Volleys, err = strconv.Atoi(attributes["Volleys"])
	if err != nil {
		return
	}

	return
}

func checkString(s string) (string, bool) {
	if len(s) == 0 {
		return s, false
	}
	return s, true
}
