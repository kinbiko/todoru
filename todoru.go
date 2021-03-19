package todoru

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(args []string, filepath string) (string, error) {
	app, err := newApplication(filepath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := app.commit(filepath); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if len(args) == 0 {
		return app.get(), nil
	}

	if args[0] == "pop" {
		got := app.get()
		if err := app.pop(); err != nil {
			return "", err
		}
		return fmt.Sprintf(`popped "%s"`, got), nil
	}
	app.add(strings.Join(args, " "))
	return "", nil
}

type application struct {
	stack []string
}

func newApplication(filepath string) (*application, error) {
	app := &application{}
	lines, err := readFile(filepath)
	if err != nil {
		return nil, err
	}
	app.stack = lines
	return app, nil
}

func readFile(filepath string) ([]string, error) {
	// TODO: This works even when the file doesn't exist, but we should
	// probably just read the file we opened instead of closing it and reading
	// it again.
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("unable to open file to read %s: %w", filepath, err)
	}
	file.Close()
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file '%s': %w", filepath, err)
	}
	if len(content) == 0 {
		return nil, nil
	}
	return strings.Split(string(content), "\n"), nil
}

func (a *application) get() string {
	if len(a.stack) > 0 {
		return a.stack[0]
	}
	return "Nothing left to do!"
}

func (a *application) add(todo string) {
	a.stack = append([]string{todo}, a.stack...)
}

func (a *application) pop() error {
	if len(a.stack) == 0 {
		return fmt.Errorf("Stack is empty. Nothing to pop")
	}
	a.stack = a.stack[1:]
	return nil
}

func (a *application) commit(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file to write '%s': %w", filepath, err)
	}
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("unable to truncate file '%s': %w", filepath, err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()
	if _, err := w.WriteString(strings.Join(a.stack, "\n")); err != nil {
		return fmt.Errorf("unable to write to file '%s': %w", filepath, err)
	}
	return nil
}
