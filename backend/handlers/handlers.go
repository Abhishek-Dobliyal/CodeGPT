package handlers

import (
	"log"
	"os"

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
}

func getChatGPTResponse(prompt string) models.ChatGPTApiResponse {
	// ToDo: Call ChatGPT api and fetch the reponse
	err := gptClient.ChatCompletionStream()
	return models.ChatGPTApiResponse{}
}

func RetrieveFromApi(ctx *fiber.Ctx) error {
	request := new(models.Request)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Println((err.Error()))
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return nil
}

func TestEndpoint(ctx *fiber.Ctx) error {
	testData := "testing this enpoint to check whether the fiber server is running when deployed"
	return ctx.Status(fiber.StatusOK).JSON(models.Response{
		Message:    "test endpoint",
		StatusCode: fiber.StatusOK,
		Data:       &testData,
	})
}
