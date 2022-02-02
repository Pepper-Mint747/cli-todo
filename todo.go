package todo

import (
	"fmt"
	"time"
)

type item struct {
	// Item structrepresents a ToDo item
	Task        string
	Done        bool
	CreatedAt   time.Time
	completedAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		completedAt: time.Time{},
	}
	*l = append(*l, t)
}

//Complete method marks a ToDo item as completed by setting Done = true
//and ComletedAt to the currrent time
func (l *List) Complete(i int) error {
	ls := *l
	if i < +0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjust index for 0 based index
	ls[i-1].Done = true
	ls[i-1].completedAt = time.Now()

	return nil
}

// Delete method delees a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	//Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}
