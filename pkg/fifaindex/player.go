package fifaindex

import "time"

type Player struct {
	Information PlayerInformation
	BallSkills  PlayerBallSkills
	Defence     PlayerDefence
	Mental      PlayerMental
	Passing     PlayerPassing
	Physical    PlayerPhysical
	Shooting    PlayerShooting
	Goalkeeper  PlayerGoalkeeper
}

type PlayerInformation struct {
	OVR           int
	POT           int
	Height        int
	Weight        int
	PrefFoot      string
	BirthDate     time.Time
	Age           int
	PrefPositions []string
	AWR           string
	DWR           string
	WeakFoot      int
	SkillMoves    int
	Value         float64
	Wage          float64
}

type PlayerBallSkills struct {
	BallControl int
	Dribbling   int
}

type PlayerPassing struct {
	Crossing  int
	ShortPass int
	LongPass  int
}

type PlayerGoalkeeper struct {
	Positioning int
	Diving      int
	Handling    int
	Kicking     int
	Reflexes    int
}

type PlayerDefence struct {
	Marking     int
	SideTackle  int
	StandTackle int
}

type PlayerPhysical struct {
	Acceleration int
	Stamina      int
	Strength     int
	Balance      int
	SprintSpeed  int
	Agility      int
	Jumping      int
}

type PlayerMental struct {
	Aggression    int
	Reactions     int
	AttPosition   int
	Interceptions int
	Vision        int
	Composure     int
}

type PlayerShooting struct {
	Heading   int
	ShotPower int
	Finishing int
	LongShots int
	Curve     int
	FKAcc     int
	Penalties int
	Volleys   int
}
