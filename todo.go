package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("invalid index %d", i)
	}
	(*l)[i-1].Done = true
	(*l)[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("invalid index %d", i)
	}
	*l = append((*l)[:i-1], (*l)[i:]...)
	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (l *List) Load(filename string) error {
	js, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, l)
}
