package models

import "time"

// Guild Model
// @Description Model of guild data
type Guild struct {
	GuildId     string     `json:"guild_id"`
	Multiplier  int        `json:"multiplier"`
	PbChannelId NullString `json:"pb_channel_id"`
}

// User Model
// @Description Model of active guild member
type User struct {
	UserId  string `json:"user_id"`
	GuildId string `json:"guild_id"`
	Points  int    `json:"points"`
}

// Users Model
// @Description Model of active guild members
type Users struct {
	Users []User `json:"users,omitempty"`
}

// Time Model
// @Description Model of a fetched time and the team
type Time struct {
	Time     int        `json:"time"`
	BossName string     `json:"boss_name"`
	RunId    int        `json:"run_id"`
	Date     time.Time  `json:"date"`
	Team     []Teammate `json:"team"`
}

type Teammate struct {
	RunId   int    `json:"run_id"`
	GuildId string `json:"guild_id"`
	UserId  string `json:"user_id"`
}

type RSN struct {
	RSN     string `json:"rsn"`
	WomId   string `json:"wom_id"`
	UserId  string `json:"user_id"`
	GuildId string `json:"guild_id"`
}

// Body Model
// @Description HTTP Body model for all responses
type Body struct {
	Content interface{} `json:"content,omitempty"`
}

// Body Model
// @Description HTTP Body model for all responses
type Empty struct{}
