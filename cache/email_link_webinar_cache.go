package cache

import "github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"

type EmailLinkWebinarCache interface {
	Set(key string, value *entity.Event)
	Get(key string) *entity.Event
	Del(key string)
}
