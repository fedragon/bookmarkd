package internal

type Config struct {
	HttpAddress string `envconfig:"BOOKMARKD_HTTP_ADDRESS" default:"0.0.0.0:11235"`
}
