package model

type CartStatus string

const (
	Saved     CartStatus = "saved"
	Completed CartStatus = "completed"
)

func (c CartStatus) String() string {
	return string(c)
}
