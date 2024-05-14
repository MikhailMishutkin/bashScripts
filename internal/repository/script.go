package repository

import (
	"bashscripts/internal/models"
	"context"
	"fmt"
)

// ...
func (r *Repo) SaveScript(ctx context.Context, s *models.Script) (id int, err error) {
	const query = `
INSERT INTO scripts (name, result) VALUES($1, $2)  ON CONFLICT(name) DO nothing RETURNING id
`
	err = r.db.QueryRow(ctx,
		query,
		s.Name,
		s.Result,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("can't create script name, perhaps name is existing: %v", err)
	}

	return id, nil
}

func (r *Repo) SaveCommands(c string, id int) error {
	const query = `
INSERT INTO commands (command, scr_id) VALUES ($1, $2)
`
	_, err := r.db.Exec(context.Background(), query, c, id)
	if err != nil {
		return fmt.Errorf("error to save commands to db", err)
	}

	return nil
}

func (r *Repo) GetScriptsList(ctx context.Context) ([]*models.Script, error) {
	const query = `SELECT * 
	FROM scripts
	`

	var scriptsList []*models.Script

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("errot to get scripts: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.Script{}
		err := rows.Scan(&s.UUID, &s.Name, &s.Result)
		if err != nil {
			return nil, fmt.Errorf("error to scripts: %v\n", err)
		}

		scriptsList = append(scriptsList, s)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	return scriptsList, nil
}

func (r *Repo) GetCommandsList(ctx context.Context, id int) (s []string, err error) {
	const query = `SELECT command
	FROM commands
	WHERE scr_id = $1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("errot to get commands from db: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var command string
		err := rows.Scan(&command)
		if err != nil {
			return nil, fmt.Errorf("error to scan commands: %v\n", err)
		}

		s = append(s, command)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	return s, nil
}

func (r *Repo) GetScriptIdByName(ctx context.Context, name string) (id int, err error) {

	err = r.db.QueryRow(ctx,
		"SELECT id FROM scripts WHERE name = $1",
		name,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error get scr id info: %v\n", err)
	}

	return id, err
}
