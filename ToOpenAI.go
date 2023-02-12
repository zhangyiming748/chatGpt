package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/zhangyiming748/log"
	"strings"
)

var answer = ""

func GetResponse(client gpt3.Client, ctx context.Context, quesiton string) string {
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
	ans := answer
	answer = ""
	return ans
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func ChatGPT(key, ask string) string {
	ctx := context.Background()
	client := gpt3.NewClient(key)
	fmt.Print("still alive")
	question := ask
	log.Debug.Printf("接收到了问题%v\n", question)
	return GetResponse(client, ctx, question)
}
