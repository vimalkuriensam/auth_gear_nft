package enums

type UserKind int

const (
	UserWithToken UserKind = iota
	UserWithoutToken
)
