package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/arunkpatra/todo/internal/app/todo/types"
)

var MyTodoService ITodoService

type TodoService struct {
	DB *sql.DB
}

// connect returns SQL database connection from the pool
func (s *TodoService) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (t *TodoService) Read(ctx context.Context) (*types.ToDo, error) {

	// get SQL connection from pool
	c, err := t.connect(ctx)
	if err != nil {
		panic("DB get connection from pool failed")
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description` FROM todo WHERE `ID`=?",
		1)
	if err != nil {
		panic("READ query failed")
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("Next failed")
		return nil, err
	}
	var td types.ToDo

	if err := rows.Scan(&td.ID, &td.Title, &td.Description); err != nil {
		fmt.Println(err)
		//panic("Read record failed. Was expecting a single record")
		return nil, err
	} else {
		return &td, nil
		//fmt.Println(td)
	}
}

type ITodoService interface {
	Read(context.Context) (*types.ToDo, error)
	Hello(context.Context) (*types.HelloResponse, error)
}

func (t *TodoService) Hello(ctx context.Context) (*types.HelloResponse, error) {
	var hr = &types.HelloResponse{Message: "Hello World!"}
	return hr, nil
}
