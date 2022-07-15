package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// factsCmd represents the facts command
var factsCmd = &cobra.Command{
	Use:   "facts",
	Short: "Get a random Chuck Norris fact",
	Long:  "This command fetches Chuck Norris facts from the chucknorris.io api",
	Run: func(cmd *cobra.Command, args []string) {
		category, err := cmd.Flags().GetString("cat")

		if err != nil {
			fmt.Printf("Error reading cat flag: %v", err)
		}

		if category != "" {
			exist := checkValidCategory(category)

			if exist {
				getChuckNorrisFactWithCategory(category)
			} else {
				fmt.Printf("Category %s does not exist", category)
			}

		} else {
			getChuckNorrisFact()
		}
	},
}

func init() {
	rootCmd.AddCommand(factsCmd)

	rootCmd.PersistentFlags().String("cat", "", "Search category for Chuck Norris Fact")
}

type RandomFact struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type RandomFacts struct {
	Total  int          `json:"total"`
	Result []RandomFact `json:"result"`
}

func getChuckNorrisFact() {
	url := "https://api.chucknorris.io/jokes/random"

	respBody := makeChuckNorrisCall(url)
	randomFact := RandomFact{}
	err := json.Unmarshal(respBody, &randomFact)

	if err != nil {
		log.Printf("Error unmarshalling %v", err)
	}

	fmt.Println(randomFact.Value)
}

func getChuckNorrisFactWithCategory(category string) {
	url := fmt.Sprintf("https://api.chucknorris.io/jokes/search?query=%s", category)

	respBody := makeChuckNorrisCall(url)
	randomFacts := RandomFacts{}
	err := json.Unmarshal(respBody, &randomFacts)

	if err != nil {
		log.Printf("Error unmarshalling %v", err)
	}

	fact := randomiseFact(randomFacts)

	fmt.Println(fact.Value)
}

func randomiseFact(randomFacts RandomFacts) RandomFact {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := randomFacts.Total
	randomIndex := rand.Intn(max-min+1) + min
	return randomFacts.Result[randomIndex]

}

func makeChuckNorrisCall(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error making call %v", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error reading body %v", err)
	}
	err = resp.Body.Close()

	if err != nil {
		log.Printf("Error closing body %v", err)
	}

	return body
}

func checkValidCategory(category string) bool {
	categoriesList := []string{"fashion", "animal", "history"}

	for _, c := range categoriesList {
		if strings.ToLower(category) == c {
			return true
		}
	}

	return false
}

// fashion, animal, history
