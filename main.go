package main

import (
	"bufio"
	"fmt"
	"os"

	"chatbot-cli/db"
	"chatbot-cli/ollama"
	"chatbot-cli/prompt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🤖 Chatbot MySQL (ketik 'exit' untuk keluar)")
	for {
		fmt.Print("You: ")
		q, _ := reader.ReadString('\n')

		if q == "exit\n" {
			break
		}

		context, _ := db.GetFinanceStatement(database)
		fullPrompt := prompt.Build(context, q)
		// fmt.Print(fullPrompt)
		answer, _ := ollama.Generate(fullPrompt)
		fmt.Println("Bot:", answer)

	}
}
