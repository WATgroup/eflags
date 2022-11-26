package eflags

import (
    "github.com/spf13/pflag"
    "fmt"
    "testing"
)

var a string

func init() {
    x := StringChoice(&a,`default`,[]string{`one`,`two`,`three`})

    pflag.VarP(x,"choice","s","Help for this choice")
}


func TestStringChoice(t *testing.T) {

    cl := pflag.CommandLine
    cl.Parse([]string{"-s","four"})
    fmt.Println("Flag is:",a)

    cl.Parse([]string{"-s","one"})
    fmt.Println("Flag is:",a)

    cl.Parse([]string{"-s","two"})
    fmt.Println("Flag is:",a)

    cl.Parse([]string{"--choice","three"})
    fmt.Println("Flag is:",a)
}
