package internal

type Config struct {
	HttpAddress string `envconfig:"BOOKMD_HTTP_ADDRESS" default:"0.0.0.0:3333"`
}
