package types

import (
	"encoding/json"
	"fmt"
)

type ToDo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func (hr *HelloResponse) String() string {
	return fmt.Sprint(json.Marshal(hr))
}

func (td *ToDo) String() string {
	//return fmt.Sprintf("{ID = %d, Title = %s, Description = %s", td.ID, td.Title, td.Description)
	return fmt.Sprint(json.Marshal(td))
}
