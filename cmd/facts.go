/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// factsCmd represents the facts command
var factsCmd = &cobra.Command{
	Use:   "facts",
	Short: "Get a random Chuck Norris fact",
	Long:  "This command fetches Chuck Norris facts from the chucknorris.io api",
	Run: func(cmd *cobra.Command, args []string) {
		getChuckNorrisFact()
	},
}

func init() {
	rootCmd.AddCommand(factsCmd)
}

type RandomFact struct {
	ID    string `json:"id"`
	Value string `json:"value"`
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
