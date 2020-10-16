package fifaindex

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

//go:generate easytags $GOFILE

type Player struct {
	Information PlayerInformation `json:"information"`
	BallSkills  PlayerBallSkills  `json:"ball_skills"`
	Defence     PlayerDefence     `json:"defence"`
	Mental      PlayerMental      `json:"mental"`
	Passing     PlayerPassing     `json:"passing"`
	Physical    PlayerPhysical    `json:"physical"`
	Shooting    PlayerShooting    `json:"shooting"`
	Goalkeeper  PlayerGoalkeeper  `json:"goalkeeper"`
}

type PlayerInformation struct {
	OVR           int       `json:"ovr"`
	POT           int       `json:"pot"`
	Height        int       `json:"height"`
	Weight        int       `json:"weight"`
	PrefFoot      string    `json:"pref_foot"`
	BirthDate     time.Time `json:"birth_date"`
	Age           int       `json:"age"`
	PrefPositions []string  `json:"pref_positions"`
	AWR           string    `json:"awr"`
	DWR           string    `json:"dwr"`
	WeakFoot      int       `json:"weak_foot"`
	SkillMoves    int       `json:"skill_moves"`
	Value         float64   `json:"value"`
	Wage          float64   `json:"wage"`
}

type PlayerBallSkills struct {
	BallControl int `json:"ball_control"`
	Dribbling   int `json:"dribbling"`
}

type PlayerPassing struct {
	Crossing  int `json:"crossing"`
	ShortPass int `json:"short_pass"`
	LongPass  int `json:"long_pass"`
}

type PlayerGoalkeeper struct {
	Positioning int `json:"positioning"`
	Diving      int `json:"diving"`
	Handling    int `json:"handling"`
	Kicking     int `json:"kicking"`
	Reflexes    int `json:"reflexes"`
}

type PlayerDefence struct {
	Marking     int `json:"marking"`
	SideTackle  int `json:"side_tackle"`
	StandTackle int `json:"stand_tackle"`
}

type PlayerPhysical struct {
	Acceleration int `json:"acceleration"`
	Stamina      int `json:"stamina"`
	Strength     int `json:"strength"`
	Balance      int `json:"balance"`
	SprintSpeed  int `json:"sprint_speed"`
	Agility      int `json:"agility"`
	Jumping      int `json:"jumping"`
}

type PlayerMental struct {
	Aggression    int `json:"aggression"`
	Reactions     int `json:"reactions"`
	AttPosition   int `json:"att_position"`
	Interceptions int `json:"interceptions"`
	Vision        int `json:"vision"`
	Composure     int `json:"composure"`
}

type PlayerShooting struct {
	Heading   int `json:"heading"`
	ShotPower int `json:"shot_power"`
	Finishing int `json:"finishing"`
	LongShots int `json:"long_shots"`
	Curve     int `json:"curve"`
	FKAcc     int `json:"fk_acc"`
	Penalties int `json:"penalties"`
	Volleys   int `json:"volleys"`
}

var attributes map[string]string

type PlayerNameEmptyError struct {
	msg string
}

func (e PlayerNameEmptyError) Error() string {
	return e.msg
}

type PlayerNotFoundError struct {
	msg string
}

func (e PlayerNotFoundError) Error() string {
	return e.msg
}

type MultiplePlayersFoundError struct {
	msg string
}

func (e MultiplePlayersFoundError) Error() string {
	return e.msg
}

func GetPlayerLink(playerName string) (link string, err error) {
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	c.Limit(&colly.LimitRule{
		DomainGlob: "www.fifaindex.com/*",
		Delay:      time.Second,
	})

	//Prepare search parameters
	if playerName == "" {
		err = PlayerNameEmptyError{"Player name is empty"}
		return
	}
	paraName := "?name=" + strings.Replace(strings.Trim(playerName, " "), " ", "+", -1)

	found := make(map[string]string)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "/player/") && e.Text != "" {
			found[e.Text] = e.Attr("href")
		}
	})

	c.Visit("https://www.fifaindex.com/players/" + paraName)
	c.Wait()

	switch len(found) {
	case 0:
		err = PlayerNotFoundError{"Player " + playerName + " not found"}
	case 1:
		for _, v := range found {
			link = v
		}
	default:
		s := "Found multiple players: "
		for k, v := range found {
			s += k + ": " + v + "; "
		}
		err = MultiplePlayersFoundError{s}
	}

	return
}

// GetPlayer a
func GetPlayer(link, dateID string) (player Player, err error) {
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
	})
	c.Visit("https://www.fifaindex.com" + link + dateID)
	c.Wait()

	player, err = assablePlayer(attributes)

	return
}

func assablePlayer(attributes map[string]string) (p Player, err error) {
	p = Player{}

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
