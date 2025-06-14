package store

import (
	"todo-app/model"
)

func (ps *PostgresStore) Create(task model.Task) (model.Task, error) {
	query := `INSERT INTO tasks (title, content, done) VALUES ($1, $2, $3) RETURNING id`
	err := ps.DB.QueryRow(query, task.Title, task.Content, false).Scan(&task.ID)
	return task, err
}
func (ps *PostgresStore) List() ([]model.Task, error) {
	rows, err := ps.DB.Query(`SELECT id, title, content, done FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func (ps *PostgresStore) Delete(id int) error {
	_, err := ps.DB.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	return err
}
func (ps *PostgresStore) MarkDone(id int) error {
	_, err := ps.DB.Exec(`UPDATE tasks SET done = TRUE WHERE id = $1`, id)
	return err
}
