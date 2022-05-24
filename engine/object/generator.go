//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"go/format"
	"log"
	"os"
	"text/template"
)

type genData struct {
	Types []TypeData
}

type TypeData struct {
	TypeName   string
	Components []string
}

var data = genData{
	Types: []TypeData{
		{
			TypeName: "Camera",
			Components: []string{
				"Transform2D",
				"Camera",
			},
		},
		{
			TypeName: "Mob",
			Components: []string{
				"Transform2D",
				"SpriteRenderer",
				"RigidBody2D",
				"Collider2D",
			},
		},
	},
}

func main() {
	generateObject(data)
}

func generateObject(objectData genData) {
	tmpl := template.Must(template.ParseFiles("object.gotmpl"))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, "object.gotmpl", objectData)
	if err != nil {
		log.Fatalf("Unable to parse template : %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		panic(err)
	}
	pwd, _ := os.Getwd()
	f, _ := os.Create(pwd + "/generated_objects.go")
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()

}
