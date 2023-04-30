package handlers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Abhishek-Dobliyal/backend/constants"
	"github.com/Abhishek-Dobliyal/backend/models"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	apiKey    string
	gptClient gpt3.Client
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load API key from env", err.Error())
	}
	apiKey = os.Getenv("API_KEY")
	gptClient = gpt3.NewClient(apiKey)

	fmt.Println(apiKey, gptClient)
}

func getChatGPTResponse(prompt string) models.ChatGPTApiResponse {
	var chatGPTResp models.ChatGPTApiResponse

	err := gptClient.CompletionStreamWithEngine(
		context.Background(),
		gpt3.TextDavinci003Engine,
		gpt3.CompletionRequest{
			Prompt:      []string{prompt},
			MaxTokens:   gpt3.IntPtr(500),
			Temperature: gpt3.Float32Ptr(0),
		},
		func(cr *gpt3.CompletionResponse) {
			chatGPTResp = models.ChatGPTApiResponse{
				Result: &cr.Choices[0].Text,
			}
		},
	)
	if err != nil {
		log.Println("could not fetch the reponse", err.Error())
		return models.ChatGPTApiResponse{}
	}

	return chatGPTResp
}

func createPromptFromAction(request *models.Request) string {
	var prompt string

	switch request.Action {

	case constants.ActionRefactor:
		addTestFlag, multipleFuncFlag := request.AddTest, request.MultipleFunc
		if addTestFlag != nil && *addTestFlag {
			prompt += "\nadd test cases for the given coode if possible"
		}
		if multipleFuncFlag != nil && *multipleFuncFlag {
			prompt += "\nseparate code into multiple functions if possible."
		}

	case constants.ActionPrettify:
		indentSize, bestPracticeFlag := request.IndentSize, request.BestPractice
		if indentSize > 0 {
			prompt += fmt.Sprintf("\nuse indent size of %d\n", indentSize)
		}
		if bestPracticeFlag != nil && *bestPracticeFlag {
			prompt += "\nuse the best practices as per language standards"
		}

	case constants.ActionConvert:
		convertToLang := request.ConvertTo
		if convertToLang != nil {
			prompt += fmt.Sprintf("\nconvert th given code to %s", *convertToLang)
		}

	case constants.ActionComment:
		cmtEachLineFlag, conciseCmtFlag := request.CommentEachLine, request.ConciseComment
		if cmtEachLineFlag != nil && *cmtEachLineFlag {
			prompt += "\nadd comments to every line"
		}
		if conciseCmtFlag != nil && *conciseCmtFlag {
			prompt += "\nadd comments and make comments concise and to the point"
		}
	}

	return prompt
}

func ProcessCode(ctx *fiber.Ctx) error {
	request := new(models.Request)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Println((err.Error()))
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message:    "could not parse request body",
			StatusCode: fiber.StatusBadRequest,
		})
	}

	prompt := createPromptFromAction(request)
	resp := getChatGPTResponse(prompt)
	if resp.Result == nil {
		log.Println("unable to retreive response from chatgpt api")
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message:    "unable to retreive response from chatgpt api",
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(models.Response{
		Message:    "success",
		StatusCode: fiber.StatusOK,
		Data:       resp.Result,
	})
}

func TestEndpoint(ctx *fiber.Ctx) error {
	testData := "testing this enpoint to check whether the fiber server is running when deployed"
	return ctx.Status(fiber.StatusOK).JSON(models.Response{
		Message:    "test endpoint",
		StatusCode: fiber.StatusOK,
		Data:       &testData,
	})
}
