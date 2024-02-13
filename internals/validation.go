package internals

func BuildOtNotToBuilid(upgrades []string, latest,status_gov string  ) bool {
	for _, upgrade := range upgrades {
		if upgrade == latest && status_gov == "PROPOSAL_STATUS_PASSED" {
			return true
		}
	}
	return false
}