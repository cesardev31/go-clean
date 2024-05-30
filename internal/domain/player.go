package domain

import "time"

type Player struct {
	Name         string    `json:"name" bindig:"required"`
	Age          int       `json:"age" bindig:"required"`
	CreationTime time.Time `json:"-"`
}
