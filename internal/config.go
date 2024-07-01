package internal

type Config struct {
	HttpAddress    string `envconfig:"BOOKMD_HTTP_ADDRESS" default:"0.0.0.0:3333"`
	ObsidianVault  string `envconfig:"BOOKMD_OBSIDIAN_VAULT" required:"true"`
	ObsidianFolder string `envconfig:"BOOKMD_OBSIDIAN_FOLDER" default:"Clippings"`
}
