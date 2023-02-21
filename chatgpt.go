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
	
	resp, err := client.CompletionWithEngine(ctx,"davinci", gpt3.CompletionRequest{
		Prompt:    []string{"I am a highly intelligent question answering bot. If you ask me a question that is rooted in truth, I will give you the answer. If you ask me a question that is nonsense, trickery, or has no clear answer, I will respond with \"Unknown\".\n\nQ: What is human life expectancy in the United States?\nA: Human life expectancy in the United States is 78 years.\n\nQ: Who was president of the United States in 1955?\nA: Dwight D. Eisenhower was president of the United States in 1955.\n\nQ: Which party did he belong to?\nA: He belonged to the Republican Party.\n\nQ: What is the square root of banana?\nA: Unknown\n\nQ: How does a telescope work?\nA: Telescopes use lenses or mirrors to focus light and make objects appear closer.\n\nQ: Where were the 1992 Olympics held?\nA: The 1992 Olympics were held in Barcelona, Spain.\n\nQ: How many squigs are in a bonk?\nA: Unknown\n\nQ:" + question+"?\nA:"},
		MaxTokens: gpt3.IntPtr(100),
		Stop:      []string{"\n"},
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