package cors

import "github.com/rs/cors"

// subscribe the nessesary cors configuration
func Setup() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"Token"},
	})
}
