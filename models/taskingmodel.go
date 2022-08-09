package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/entities"
)

type TaskingModel struct {
	conn *sql.DB
}

func NewTaskingModel() *TaskingModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &TaskingModel{
		conn: conn,
	}

}

func (p *TaskingModel) FindAll() ([]entities.Tasking, error) {

	rows, err := p.conn.Query("select * from task")
	if err != nil {
		return []entities.Tasking{}, err
	}
	defer rows.Close()

	var dataTasking []entities.Tasking
	for rows.Next() {
		var tasking entities.Tasking
		rows.Scan(&tasking.Id, &tasking.Task, &tasking.Assigne, &tasking.DeadLine, &tasking.IsDone, &tasking.UpdateDate)

		dataTasking = append(dataTasking, tasking)
	}

	return dataTasking, nil

}

func (p *TaskingModel) Create(tasking entities.Tasking) bool {

	result, err := p.conn.Exec("insert into task (task, assigne, deadline) values (?,?,?)",
		tasking.Task, tasking.Assigne, tasking.DeadLine)

	if err != nil {
		fmt.Println(err)
		return false
	}

	LastInsertId, _ := result.LastInsertId()

	return LastInsertId > 0

}

func (p *TaskingModel) Find(id int64, tasking *entities.Tasking) error {
	return p.conn.QueryRow("select * from task where id = ?", id).Scan(
		&tasking.Id,
		&tasking.Task,
		&tasking.Assigne,
		&tasking.DeadLine,
		&tasking.IsDone,
		&tasking.UpdateDate)
}

func (p *TaskingModel) Update(tasking entities.Tasking) error {

	_, err := p.conn.Exec(
		"update task set task = ?, assigne = ?, deadline = ?, updatedate=now() where id = ? ",
		tasking.Task, tasking.Assigne, tasking.DeadLine, tasking.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *TaskingModel) Delete(id int64) {
	p.conn.Exec("delete from task where id = ?", id)
}

func (p *TaskingModel) IsDone(id int64) {
	p.conn.Exec("update task set isdone=1, updatedate=now() where id = ?", id)
}

// func (p *TaskingModel) UpdateDate(id int64) {
// 	p.conn.Exec("update task set updatedate=now() where id = ?", id)
// }
