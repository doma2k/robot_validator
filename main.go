package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	// "os"
	// "path/filepath"
)

const propsURL = "http://162.55.135.119:1317/cosmos/gov/v1beta1/proposals?proposal_status=PROPOSAL_STATUS_PASSED&pagination.limit=10&pagination.reverse=true"
const propType = "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"

type Proposals struct {
	Proposals []Proposal `json:"proposals"`
}

type Proposal struct {
	ProposalID string `json:"proposal_id"`
	Content    struct {
		Type        string `json:"@type"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"content"`
	Status string `json:"status"`
}

func getProposals() {

	resp, err := http.Get(propsURL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var proposals Proposals
	err = json.Unmarshal(body, &proposals)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	for _, proposal := range proposals.Proposals {
		if propType == proposal.Content.Type {
			log.Printf("Proposal ID: %s\n", proposal.ProposalID)
			log.Printf("Title: %s\n", proposal.Content.Title)
			log.Printf("Status: %s\n", proposal.Status)
		}
	}
}

func main() {
	getProposals()
}
