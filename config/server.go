package config

type ServerConfig struct {
	Host            string `env:"HOST"          envDefault:"localhost:9000" json:"host"`
	PublicHost      string `env:"PUBLIC_HOST"   envDefault:"localhost"      json:"public_host"`
	PublicScheme    string `env:"PUBLIC_SCHEME" envDefault:"http"           json:"public_scheme"`
	Address         string `env:"ADDRESS"       envDefault:":9000"          json:"address"`
	HttpIdleTimeout int    `env:"HTTP_IDLE_TIMEOUT"       envDefault:"1" `
	WorkersTimeout  int    `env:"WORKERS_TIMEOUT"       envDefault:"10" `
}
