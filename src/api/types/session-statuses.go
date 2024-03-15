package types

type SessionStatus struct {
	ID        int
	Name      string
	IsDeleted bool
	Sessions  []Session
}
