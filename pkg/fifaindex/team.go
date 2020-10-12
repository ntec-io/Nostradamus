package fifaindex

type Team struct {
	Name           string
	Rating         float64
	League         League
	Attack         int
	Midfield       int
	Defence        int
	Overall        int
	TransferBudget float64
	RivalTeam      string

	DefensiveStyle TeamDefensiveStyle
	OffensiveStyle TeamOffensiveStyle
	PlayerRoles    TeamPlayerRoles
}

type TeamDefensiveStyle struct {
	Style string
	Width int
	Depth int
}

type TeamOffensiveStyle struct {
	Style        string
	Width        int
	PlayersInBox int
	Corners      int
	FreeKicks    int
}

type TeamPlayerRoles struct {
	Captain       string
	ShortFreeKick string
	LongFreeKick  string
	Penalties     string
	LeftCorner    string
	RightCorner   string
}
