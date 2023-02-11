package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/zhangyiming748/log"
	"os"
	"strings"
	"time"
)

const Key = ""

var answer = ""

func GetResponse(client gpt3.Client, ctx context.Context, quesiton string) {
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

	fmt.Printf("%v\n", answer)
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	//log.SetOutput(new(NullWriter))
	ctx := context.Background()
	client := gpt3.NewClient(Key)
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("輸入你的問題(quit 離開): ")

				if !scanner.Scan() {
					break
				}

				question := scanner.Text()
				questionParam := validateQuestion(question)
				switch questionParam {
				case "quit":
					quit = true
				case "":
					continue
				default:
					GetResponse(client, ctx, questionParam)
				}
			}
		},
	}
	log.Warn.Fatal(rootCmd.Execute())
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
