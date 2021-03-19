package todoru

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

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

	t.Run("add + get + pop + get", func(t *testing.T) {
		todo := "parse XML with regex"

		filepath := tempFilepath()
		got := call(strings.Split(todo, " "), filepath)
		if got != "" {
			t.Errorf(`expected "" but got "%s"`, got)
		}

		got = call([]string{}, filepath)
		if got != todo {
			t.Errorf("expected '%s' but got '%s'", todo, got)
		}

		got = call([]string{"pop"}, filepath)
		if got != fmt.Sprintf(`popped "%s"`, todo) {
			t.Errorf(`expected "popped %s" but got "%s"`, todo, got)
		}

		got = call([]string{}, filepath)
		exp := "Nothing left to do!"
		if got != exp {
			t.Errorf(`expected "%s" but got "%s"`, exp, got)
		}

	})
}
