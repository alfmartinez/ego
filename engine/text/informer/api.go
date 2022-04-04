package informer

import (
	"fmt"
)

func Publish(message Message) RuleResult {
	return currentStory.Publish(message)
}

func GetObject(name string) Object {
	return currentStory.GetObject(name)
}

func IsSame(a, b string) bool {
	if a == "somewhere" || b == "somewhere" {
		return true
	}
	return GetObject(a) == GetObject(b)
}

func Debug() bool {
	return currentStory.Debug()
}

func Say(message string) {
	currentStory.Say(message)
}

func SetAlias(alias string, o Object) {
	currentStory.SetAlias(alias, o)
}

func Carry(subject Object, items ...Object) {
	for _, item := range items {
		fmt.Printf("Give %s to %s", subject, item)
	}
}
