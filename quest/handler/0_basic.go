package handler

import (
	questshare "quest/share"
)

type BaseTaskHandler interface {
	CanUpdate(args ...any) bool
	Count() questshare.Progress
	GetCriteriaType() questshare.CriteriaType
	GetProgressSetType() questshare.ProgressSetType
}
