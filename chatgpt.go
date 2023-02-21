package main
import (
	"context"
	"fmt"
	"log"
	"os"
	"bufio"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}
var question string;
fmt.Println("Enter your question");
scanner := bufio.NewScanner(os.Stdin)

if scanner.Scan() {
	question = scanner.Text()
}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	
	resp, err := client.CompletionWithEngine(ctx,"text-davinci-003", gpt3.CompletionRequest{
		Prompt:    []string{question},
		MaxTokens: gpt3.IntPtr(100),
		Stop:      []string{"."},
		Echo:      false,
		Temperature: gpt3.Float32Ptr(0),
		TopP: gpt3.Float32Ptr(1),
		FrequencyPenalty: *gpt3.Float32Ptr(0),
		PresencePenalty: *gpt3.Float32Ptr(0),
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}