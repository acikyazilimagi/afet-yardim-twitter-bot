package config

type Config struct {
	Server  ServerConfig
	Twitter TwitterConfig
}

type ServerConfig struct {
	HttpAddress string `required:"true" split_words:"true"`
}
type TwitterConfig struct {
	ConsumerKey       string `required:"true" split_words:"true"`
	ConsumerSecret    string `required:"true" split_words:"true"`
	AccessToken       string `required:"true" split_words:"true"`
	AccessTokenSecret string `required:"true" split_words:"true"`
}
