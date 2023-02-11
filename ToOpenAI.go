package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/zhangyiming748/log"
	"strings"
	"time"
)

var answer = ""

func GetResponse(client gpt3.Client, ctx context.Context, quesiton string, out chan string) {
	stop := make(chan bool)
	go Wait(stop)
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		one := resp.Choices[0].Text
		answer = strings.Join([]string{answer, one}, "")
	})
	if err != nil {
		log.Warn.Panicf("%v\n", err)
	}
	log.Debug.Printf("%v\n", answer)
	answer = <-out
	log.Debug.Printf("%v\n", answer)
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func ChatGPT(key string, in chan string, out chan string) {
	//log.SetOutput(new(NullWriter))
	ctx := context.Background()
	client := gpt3.NewClient(key)

	for {
		fmt.Print("still alive")
		question := <-in
		log.Debug.Printf("接收到了问题%v\n", question)
		questionParam := validateQuestion(question)
		GetResponse(client, ctx, questionParam, out)
		if _, ok := <-in; !ok {
			log.Warn.Println("程序退出")
		}
	}
}

// 这里判断问题内容是否有空格 如果有就截断 如果是关键词就返回空 这是什么用意
func validateQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}
func Wait(ch chan bool) {
	fmt.Printf("I see")
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf(".")
		if <-ch {
			close(ch)
			break
		}
	}
}

// 如果收到退出信号 退出chatgpt程序
func Exit(ch chan bool) {

}
