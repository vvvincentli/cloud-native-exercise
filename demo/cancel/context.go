package main

import (
	"cloud-native-exercise/demo/g"
	"context"
	"fmt"
	"time"
)

type Task struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("task start.")
	var tasks []Task
	resultChan := make(chan Task)
	for i := 0; i < 100; i++ {
		tasks = append(tasks, Task{Id: i})
	}
	fmt.Println("listen result")
	g.Go(func() {
		printResult(ctx, resultChan)
	})
	fmt.Println("create task")
	for _, t := range tasks {
		g.Go(func() { task(ctx, t, resultChan) })
	}

	fmt.Println("check task canceled.")
	g.Go(func() { cancelTask(100, cancel) })
	fmt.Println("sleep 100s")
	time.Sleep(time.Second * 100)
	fmt.Println("exit.")
}

func cancelTask(id int, cancel context.CancelFunc) {
	i := 1
	for {
		//TODO: query task status by id

		fmt.Println(fmt.Sprintf("get task status %d's", i))
		if i > 10 {
			cancel()
			break
		}
		time.Sleep(time.Second * 1)
		i++
	}
}
func printResult(ctx context.Context, resultChan chan Task) {
	for {
		select {
		case val, ok := <-resultChan:
			if ok == false {
				fmt.Println("chan closed")
			}
			fmt.Println(fmt.Sprintf("receive task, id:%v,message:%v", val.Id, val.Message))
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

func task(ctx context.Context, p Task, resultChan chan Task) {
	i := 0
	t := time.After(time.Second * 200)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("task %d done", p.Id))
			p.Message = "cancel"
			resultChan <- p
			return
		case <-t:
			fmt.Println("timeout")
			return
		default:
			fmt.Println(fmt.Sprintf("task %d ,run %d.", p.Id, i))
			time.Sleep(time.Second)
		}
		i++
	}
}
