package models

import (
	"fmt"
	"log"

	"todo/database"
)

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{}

func (task Task) Save() error {
	query := `
    INSERT INTO tasks(title)
    VALUES ($1)
  	RETURNING id
  `
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var taskID int64
	err = stmt.QueryRow(task.Title).Scan(&taskID)
	if err != nil {
		log.Fatal(err)
	}

	task.ID = taskID
	fmt.Printf("task.ID inside the save func: %d\n", task.ID)

	// _, err = stmt.Exec(task.Title)
	// if err != nil {
	// return err
	// }
	//
	return nil
}

func GetAllTasks() ([]Task, error) {
	rows, err := database.DB.Query("select * from tasks")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Done,
		)
		if err != nil {
			fmt.Println("error in the loop.")
			return nil, err
		}

		tasks = append(tasks, task)

	}
	return tasks, nil
}

func GetTask(id int64) (*Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1"
	row := database.DB.QueryRow(query, id)

	var task Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Done,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (task Task) Update() error {
	query := `
	update tasks
	set title=$1, done=$2
	where id = $3
	`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Done, task.ID)
	return err
}

func (task Task) Delete() error {
	query := "delete from tasks where id=$1"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.ID)
	return err
}
