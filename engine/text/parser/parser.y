%{
    package parser

    import (
        "fmt"
        "unicode"
    )
%}
%union{
    val int
}

%token <val> DIGIT LETTER 
%%
hello: {
    fmt.Println("Hello")
}
%%
type InformerLex struct {
	s string
	pos int
}


func (l *InformerLex) Lex(lval *InformerSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c) - 'a'
		return LETTER
	}
	return int(c)
}

func (l *InformerLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
