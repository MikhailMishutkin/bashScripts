package service

import (
	"bashscripts/internal/models"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// ...
func (s *Service) CreateScript(ctx context.Context, sc *models.Script) error {
	output := Output(ctx, sc.Commands)

	sc.Result = output

	id, err := s.r.SaveScript(ctx, sc)
	if err != nil {
		return err
	}
	for _, v := range sc.Commands {
		err := s.r.SaveCommands(v, id)
		if err != nil {
			return err
		}
	}

	return nil
}

// ...
func (s *Service) GetScriptsList(ctx context.Context) (scriptsList []*models.Script, err error) {
	scriptsListRepo, err := s.r.GetScriptsList(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range scriptsListRepo {
		var commands []string
		commands, err = s.r.GetCommandsList(ctx, v.UUID)
		if err != nil {
			return nil, err
		}
		s := &models.Script{
			UUID:     v.UUID,
			Name:     v.Name,
			Commands: commands,
			Result:   v.Result,
		}
		scriptsList = append(scriptsList, s)
	}

	return scriptsList, nil
}

// ...
func (s *Service) GetScript(ctx context.Context, name string) (sc *models.Script, err error) {
	id, err := s.r.GetScriptIdByName(ctx, name)
	if err != nil {
		return nil, err
	}
	list, err := s.r.GetCommandsList(ctx, id)
	if err != nil {
		return nil, err
	}
	output := Output(ctx, list)

	sc = &models.Script{
		UUID:     id,
		Name:     name,
		Commands: list,
		Result:   output,
	}
	return sc, nil
}

func Output(ctx context.Context, commands []string) string {
	var output strings.Builder

	for _, command := range commands {
		cmd := exec.CommandContext(ctx, "bash", "-c", command)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		output.WriteString(string(out))
	}

	return output.String()
}
