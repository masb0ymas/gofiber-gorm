package constants

func AllowedOrigins() []string {
	local := []string{"http://localhost:3000", "http://localhost:3333"}
	internal := []string{"https://masb0ymas.com"}

	result := append(local, internal...)

	return result
}
