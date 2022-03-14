package command

import "testing"

func TestCommandStream(t *testing.T) {
	t.Run("Execute empty stream returns done", func(t *testing.T) {
		stream := CreateCommandStream()
		done := stream.Execute()
		if !done {
			t.Errorf("Empty Command stream should return done")
		}
	})

	t.Run("Single command, done in one go", func(t *testing.T) {
		stream := CreateCommandStream()
		cmd := CreateCommand(func() bool {
			return true
		})
		stream.After(cmd)
		done := stream.Execute()
		if !done {
			t.Errorf("Empty Command stream should return done")
		}
	})

	t.Run("Single command, done in two go", func(t *testing.T) {
		stream := CreateCommandStream()
		count := 2
		cmd := CreateCommand(func() bool {
			count--
			return count == 0
		})
		stream.After(cmd)
		done := stream.Execute()
		if done {
			t.Errorf("Command should not be empty yet")
		}
		done = stream.Execute()
		if !done {
			t.Errorf("Empty Command stream should return done")
		}
	})

	t.Run("Single command, done in two go, closure style", func(t *testing.T) {
		stream := CreateCommandStream()
		closure := func() func() bool {
			var count int = 2
			return func() bool {
				count--
				return count == 0
			}
		}()
		cmd := CreateCommand(closure)
		stream.After(cmd)
		done := stream.Execute()
		if done {
			t.Errorf("Command should not be empty yet")
		}
		done = stream.Execute()
		if !done {
			t.Errorf("Empty Command stream should return done")
		}
	})

	t.Run("Two commands, After", func(t *testing.T) {
		stream := CreateCommandStream()

		var first, second bool = false, false
		c1 := CreateCommand(func() bool {
			first = true
			return true
		})
		c2 := CreateCommand(func() bool {
			second = true
			return true
		})
		stream.After(c1)
		stream.After(c2)
		if first || second {
			t.Error("stream should not have executed any command")
		}
		stream.Execute()
		if !first || second {
			t.Error("stream should have executed first command")
		}
		stream.Execute()
		if !first || !second {
			t.Error("stream should have executed second command too")
		}
	})

	t.Run("Two commands, Before", func(t *testing.T) {
		stream := CreateCommandStream()

		var first, second bool = false, false
		c1 := CreateCommand(func() bool {
			first = true
			return true
		})
		c2 := CreateCommand(func() bool {
			second = true
			return true
		})
		stream.After(c1)
		stream.Before(c2)
		if first || second {
			t.Error("stream should not have executed any command")
		}
		stream.Execute()
		if first || !second {
			t.Error("stream should have executed second command")
		}
		stream.Execute()
		if !first || !second {
			t.Error("stream should have executed first command too")
		}
	})

}
