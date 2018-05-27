package main

func contains(values []string, value string) bool {
	exist := false
	for _, card := range values {
		if card == value {
			exist = true
		}
	}
	return exist
}
