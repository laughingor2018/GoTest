package main

import(
	"os/exec"
	//"strings"
	//"bytes"
	"context"
	"time"
	"fmt"
)

type result struct {
	out []byte
	err error
}

var resultChan chan *result

func main(){
	/*f, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)*/


	//cmd := exec.Command("tr","a-z","A-Z") 
	//cmd := exec.Command("ls","-al","/root") 
	//cmd.Stdin = strings.NewReader("something input")
	//cmd.Stdin = strings.NewReader("/root")
	//var out bytes.Buffer
	//cmd.Stdout = &out

	//err := cmd.Run()

	//out,err := cmd.CombinedOutput()

	//out,err := cmd.Output()

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Println(out.String())
	//fmt.Println(string(out))


	var (
		ctx context.Context
		cancelFunc context.CancelFunc
	)

	resultChan = make(chan *result,1000)

	ctx,cancelFunc = context.WithCancel(context.TODO())

	go func(){
		cmd := exec.CommandContext(ctx,"/bin/bash","-c","sleep 3;echo hello;")
		out,err := cmd.CombinedOutput()
		fmt.Println("out:",string(out))
		fmt.Println("err:",err)
		resultChan <- &result{
			out:out,
			err:err,
		}		
	}()

  	time.Sleep(time.Second)	

	//cancelFunc()

	r := <-resultChan

	fmt.Println("r.err:",r.err)
	fmt.Println("r.out:",string(r.out))

	cancelFunc()
}
