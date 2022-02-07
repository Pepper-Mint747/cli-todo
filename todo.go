package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	// Item struct represents a To Do item
	Task        string
	Done        bool
	CreatedAt   time.Time
	completedAt time.Time
}

// List represents a list of To Do items
type List []item

// Add creates a new to do item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		completedAt: time.Time{},
	}
	*l = append(*l, t)
}

//Complete method marks a To Do item as completed by setting Done = true
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

// Delete method deletes a To Do item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	//Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as JSON and saves it using the
//provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes the JSON
// data and parses into a list
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// String prints out a formatted list
//Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}

		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
