package service

import (
	"bashscripts/internal/models"
	"context"
)

type Service struct {
	r Repositorier
}

func NewService(r Repositorier) *Service {
	return &Service{r: r}
}

type Repositorier interface {
	SaveScript(context.Context, *models.Script) (int, error)
	SaveCommands(string, int) error
	GetScriptsList(context.Context) ([]*models.Script, error)
	GetCommandsList(context.Context, int) ([]string, error)
	GetScriptIdByName(context.Context, string) (int, error)
}
