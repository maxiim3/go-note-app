package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

// Response from the API
type AiResponse struct {
	response openai.ChatCompletionMessage
}

func main() {
	// note directory
	const DIRPATH string = "/Users/maxi/notes"
	var ai_response []AiResponse

	client := openAiClientPool()

	files := findFiles(DIRPATH)

	if len(ai_response) == 0 {
		fmt.Println("fetching AI")
		for _, file := range files {
			content := getFileContent(file)
			resp := askOpenAI(&client, file.name, content)
			ai_response = append(ai_response, resp)
		}
	}

	for _, response := range ai_response {
		// Here we want to prompt the user with the response from the AI.
		// Y/N to rename the file

		// os.WriteFile(
		// 	fmt.Sprint(DIRPATH, "/", "response.txt"),
		// 	[]byte(response.response.Content),
		// 	0666,
		// )

	}
	// fmt.Println(ai_response)
}

type OpenAIClient struct {
	client *openai.Client
	ctx    context.Context
}

func getOpenAIApiKey() string {
	return os.Getenv("OPEN_AI_API")
}

// pool client
func openAiClientPool() OpenAIClient {
	client := openai.NewClient(getOpenAIApiKey())
	ctx := context.Background()

	return OpenAIClient{client: client, ctx: ctx}
}

// Curl Open AI API
func askOpenAI(client *OpenAIClient, filename string, content string) AiResponse {
	prompt := fmt.Sprint("Based on the content of the file, you will propse a name for this file. the current name is compose like : YYYY-MM-DD.md. You will add a context before the extension. example: YYYY-MM-DD-your_answer_here.md .\n Current name : %s. \n Content: %s \n Your answer contains only the name of the file. Keep in mind to keep the extension.", filename, content)

	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	resp, err := client.client.CreateChatCompletion(client.ctx, request)

	if err != nil {
		fmt.Printf("Error from API : %v\n", err)
	}
	return AiResponse{response: resp.Choices[0].Message}
	// http request here
}

type File struct {
	name string
	path string
}

// Find recurseively all files from the path
func findFiles(path string) []File {
	var files []File
	res, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}
	for _, element := range res {

		if element.IsDir() {
			fileName := strings.Split(string(element.Name()), "")

			if fileName[0] != "." {
				concatPath := fmt.Sprint(path, "/", element.Name())

				fmt.Println(concatPath)
				recursiveFiles := findFiles(concatPath)
				files = append(files, recursiveFiles...)
			}
			//? } else if noSuffix(element) {
		} else {
			file := File{name: element.Name(), path: path}
			files = append(files, file)
		}
	}
	return files
}

// todo check if the file is already renamed
func noSuffix(element string) {
	// check if file is format : YYYY-MM-DD.md
	// NOT YYYY-MM-DD-<...>.md
}

// Get the content of the File
func getFileContent(file File) string {
	content, err := os.ReadFile(fmt.Sprint(file.path, "/", file.name))

	if err != nil {
		panic(err)
	}

	fileContent := string(content)

	return fileContent
}
