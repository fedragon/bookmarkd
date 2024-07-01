package internal

type Config struct {
	HttpPort       int    `envconfig:"BOOKMD_HTTP_PORT" default:"3000"`
	ObsidianVault  string `envconfig:"BOOKMD_OBSIDIAN_VAULT" required:"true"`
	ObsidianFolder string `envconfig:"BOOKMD_OBSIDIAN_FOLDER" default:"Clippings"`
}
