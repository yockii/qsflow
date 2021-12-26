package domain

import (
	coreDomain "github.com/yockii/qscore/pkg/domain"
)

type ProcessCategory struct {
	Id           string
	Code         string
	CategoryName string
	CreateTime   coreDomain.DateTime
	UpdateTime   coreDomain.DateTime
}

type ProcessDefinition struct {
	Id           string
	CategoryCode string
	Code         string
	ProcessName  string
	CreateTime   coreDomain.DateTime
	UpdateTime   coreDomain.DateTime
}

type ProcessNodeDefinition struct {
	Id           string
	CategoryCode string
	NodeName     string
	CreateTime   coreDomain.DateTime
	updateTime   coreDomain.DateTime
}
