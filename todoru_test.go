package todoru

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestApplication(t *testing.T) {
	t.Run("when no items are stored then none are returned", func(t *testing.T) {
		app := application{}
		if got := app.get(); got != "" {
			t.Errorf(`expected to get "" returned when no items are added`)
		}
	})

	t.Run("when adding two items then getting repeatedly only the first item is returned", func(t *testing.T) {
		app := application{}
		app.add("some other todo item that should't appear")
		todo := "do the dishes"
		app.add(todo)
		for i := 1; i <= 3; i++ {
			if got := app.get(); got != todo {
				t.Errorf(`expected to get "%s" returned when multiple items are added (iteration: %d) but got "%s"`, todo, i, got)
			}
		}
	})

	t.Run("popping when the stack is empty should give an error", func(t *testing.T) {
		app := application{}
		err := app.pop()
		if err == nil {
			t.Fatalf(`expected error but got none`)
		}
		if got := err.Error(); !strings.Contains(got, "nothing") {
			t.Fatalf(`unexpected error message: "%s"`, got)
		}
	})

	t.Run("popping works as a stack", func(t *testing.T) {
		app := application{}
		for i := 0; i < 5; i++ {
			app.add(fmt.Sprintf("item number %d", i))
		}
		for i := 4; i >= 0; i-- {
			exp := fmt.Sprintf("item number %d", i)
			if got := app.get(); got != exp {
				t.Errorf(`expected to get "%s" but got "%s"`, exp, got)
			}
			if err := app.pop(); err != nil {
				t.Errorf("unexpected error on pop(): %s", err)
			}
		}
		if got := app.get(); got != "" {
			t.Errorf(`expected to get "" returned when all items are popped but got "%s"`, got)
		}
	})
}

func TestRun(t *testing.T) {
	tempFilepath := func() string {
		f, err := os.CreateTemp("", "todo-stack-test")
		if err != nil {
			t.Fatal(err)
		}
		filepath := f.Name()
		os.Remove(filepath)
		f.Close()
		return filepath
	}

	call := func(args []string, filepath string) string {
		got, err := Run(args, filepath)
		if err != nil {
			t.Fatalf("unexpected error %s", err.Error())
		}
		return got
	}

	noArgs := []string{}

	t.Run("get nothing", func(t *testing.T) {
		if got := call(noArgs, tempFilepath()); got != "" {
			t.Errorf(`expected "" but got '%s'`, got)
		}
	})

	t.Run("add + get + pop + get", func(t *testing.T) {
		todo := "parse XML with regex"

		filepath := tempFilepath()
		got := call([]string{"add", todo}, filepath)
		if got != "" {
			t.Errorf(`expected "" but got "%s"`, got)
		}

		got = call([]string{}, filepath)
		if got != todo {
			t.Errorf("expected '%s' but got '%s'", todo, got)
		}

		got = call([]string{"pop"}, filepath)
		if got != "" {
			t.Errorf(`expected "" but got "%s"`, got)
		}

		got = call([]string{}, filepath)
		if got != "" {
			t.Errorf(`expected "" but got "%s"`, got)
		}

	})
}
