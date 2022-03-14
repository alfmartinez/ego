package command

import "testing"

func TestCommand(t *testing.T) {
	t.Run("CreateCommand returns wrapped function", func(t *testing.T) {
		var status bool = false
		cmd := CreateCommand(func() bool {
			return status
		})
		done := cmd.Execute()
		if done {
			t.Error("Command should not be done")
		}
		status = true
		done = cmd.Execute()
		if !done {
			t.Error("Command should be done")
		}
	})

	t.Run("Command can be aborted, acting as if it executed and done", func(t *testing.T) {
		var count int = 0
		cmd := CreateCommand(func() bool {
			count++
			return false
		})
		done := cmd.Execute()
		if done {
			t.Error("Command should not be done")
		}
		if count != 1 {
			t.Error("Cmd should have been executed only once")
		}
		cmd.Abort()
		done = cmd.Execute()
		if !done {
			t.Error("Command should be done")
		}
		if count != 1 {
			t.Error("Cmd should have been executed only once")
		}
	})
}
