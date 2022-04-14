%{
    package parser

    import (
        "fmt"
    )
%}
%union{
    val int
}
%%
hello: {
    fmt.Println("Hello")
}
%%