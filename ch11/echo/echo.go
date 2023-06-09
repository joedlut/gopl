package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var(
	n = flag.Bool("n",false,"omit trailing newline")
	s = flag.String("s"," ","separator")
)

var out io.Writer = os.Stdout

//执行go test main函数被忽略，没有执行
func main(){
	flag.Parse()
	if err := echo(!*n,*s,flag.Args());err !=nil{
		fmt.Fprintf(os.Stderr,"echo %v\n",err)
		os.Exit(1)
	}
}

func echo(newline bool,sep string,args []string)error{
	fmt.Fprint(out,strings.Join(args,sep))
	if newline{
		fmt.Fprintln(out)
	}
	return nil
}
