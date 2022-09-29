package common

import "time"

type TimeInfo struct {
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time `json:"deleteAt"`
}
