package main

import (
	"fmt"
	"time"
	"github.com/gorhill/cronexpr"
)


func main(){

	var (
		err error
		expr *cronexpr.Expression
		now time.Time
		next time.Time
	)

	if expr,err = cronexpr.Parse("*/5 * * * * * *");err != nil {
		fmt.Println("err:",err)
		return
	}

	fmt.Println("expr:",expr)

	now = time.Now()
	next = expr.Next(now)
	
	fmt.Println("now:",now)
	fmt.Println("next:",next)
	
	time.AfterFunc(next.Sub(now),func(){
		fmt.Println("Now:",time.Now())
	})

	time.Sleep(20 * time.Second)
}
