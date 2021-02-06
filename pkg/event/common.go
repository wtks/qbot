package event

import "time"

type Message struct {
	ID        string     `json:"id"`
	User      User       `json:"user"`
	ChannelID string     `json:"channelId"`
	Text      string     `json:"text"`
	PlainText string     `json:"plainText"`
	Embedded  []Embedded `json:"embedded"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IconID      string `json:"iconId"`
	Bot         bool   `json:"bot"`
}

type Channel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	ParentID  string    `json:"parentId"`
	Creator   User      `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Embedded struct {
	Raw  string `json:"raw"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

type MessageStamp struct {
	UserID    string    `json:"userId"`
	StampID   string    `json:"stampId"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
