package models

type Notification struct {
	ID               string      `json:"id" yaml:"id"`
	Created          string      `json:"created" yaml:"created"`
	Expiry           string      `json:"expiry" yaml:"expiry"`
	Message          string      `json:"message" yaml:"message"`
	Level            string      `json:"level" yaml:"level"`
	ChannelPopup     bool        `json:"channelPopup" yaml:"channelPopup"`
	ChannelEmail     bool        `json:"channelEmail" yaml:"channelEmail"`
	Sticky           bool        `json:"sticky" yaml:"sticky"`
	Tag              string      `json:"tag" yaml:"tag"`
	AdditionalEmails string      `json:"additionalEmails" yaml:"additionalEmails"`
	RecipientsCount  interface{} `json:"recipientsCount" yaml:"recipientsCount"`
	LastSent         string      `json:"lastSent" yaml:"lastSent"`
}

type AltinityNote struct {
	ID       string `json:"id" yaml:"id"`
	Created  string `json:"created" yaml:"created"`
	Severity string `json:"severity" yaml:"severity"`
	Title    string `json:"title" yaml:"title"`
	Message  string `json:"message" yaml:"message"`
}

func (n Notification) TableHeaders() []string {
	return []string{"ID", "LEVEL", "MESSAGE", "CREATED"}
}

func (n Notification) TableRow() []string {
	return []string{
		n.ID,
		n.Level,
		n.Message,
		n.Created,
	}
}
