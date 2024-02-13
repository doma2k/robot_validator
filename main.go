package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	api_addr  = "http://localhost:1317"
	props_url = api_addr + "/cosmos/gov/v1beta1/proposals?proposal_status=PROPOSAL_STATUS_PASSED&pagination.limit=10&pagination.reverse=true"
	repo_url  = "https://api.github.com/repos/cosmos/cosmos-sdk/releases/latest"
	propType = "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"
)

func fetchProposals(url string) []byte {
	log.Println("Fetching proposals...")
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}

func printProposals(body []byte) {
	var proposals Proposals
	err := json.Unmarshal(body, &proposals)
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
	printProposals(fetchProposals(props_url))
	LatestRelease:=getLatestRelease(repo_url)
	stagedBinaries:=getStagedBinaries()
	log.Printf("%+v :%v \n",LatestRelease, stagedBinaries)
}
