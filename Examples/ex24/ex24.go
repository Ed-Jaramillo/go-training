package main

func main() {
	var tier string
	getMontlyPrice(tier)
}

func getMontlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 10000
	case "premium":
		return 15000
	case "enterprise":
		return 50000
	default:
		return 0
	}
}