package main

import (
	"io"
	"log"
	"net/http"
	"robot-validator/utils"
	"robot-validator/internals"
	"robot-validator/types"
)

var (
	api_addr  = "http://localhost:1317"
	props_url = api_addr + "/cosmos/gov/v1beta1/proposals?proposal_status=PROPOSAL_STATUS_PASSED&pagination.limit=10&pagination.reverse=true"
	repo_url  = "https://api.github.com/repos/elys-network/elys/releases/latest"
	propType  = "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"
)

func fetchProposals(url string) []byte {
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

func queryProp(body []byte, r types.Release) string {
	var proposals types.Proposals
	var status string
	utils.UnmarshalJSON(body, &proposals)
	for _, proposal := range proposals.Proposals {
		if propType == proposal.Content.Type && r.Tag_name == proposal.Content.Title {
			status = proposal.Status
		}
	}
	return status
}

func main() {
	proposal_list := fetchProposals(props_url)
	latest_release := internals.GetLatestRelease(repo_url)
	proposa_status := queryProp(proposal_list, latest_release)
	staged_binaries := internals.GetStagedBinaries()
	log.Println("<<< Atomated Validation >>> ")
	log.Println("-------------------------------------------------")
	log.Printf("Cosmovisor staged binaries: %v\n", staged_binaries)
	log.Println("-------------------------------------------------")
	log.Printf("Elys/releases/latest: %v %v\n", latest_release.Tag_name, proposa_status)
	log.Println("-------------------------------------------------")
	log.Println("No upgrade needed")
}
