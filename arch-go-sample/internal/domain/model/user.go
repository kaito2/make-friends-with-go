package model

import "github.com/rs/xid"

type User struct {
	ID   xid.ID
	Name string
}
