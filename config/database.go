package config

type DatabaseConfig struct {
	Host             string `env:"DB_HOST" envDefault:"localhost" json:"host"`
	Port             string `env:"DB_PORT" envDefault:"5432"      json:"port"`
	User             string `env:"DB_USER,required"               json:"user"`
	Password         string `env:"DB_PASSWORD,required"           json:"password"`
	Dbname           string `env:"DB_NAME,required"               json:"dbname"`
	MaxOpenConns     int    `env:"MAX_OPEN_CONNS,required" envDefault:"1000"`
	MaxOpenIdleConns int    `env:"MAX_OPEN_IDLE_CONNS,required" envDefault:"1000"`
	CacheEnable      bool   `env:"DB_CACHE_ENABLE,required" envDefault:"true"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost" json:"host"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"      json:"port"`
	Password string `env:"REDIS_PASSWORD"                    json:"password"`
	Db       int    `env:"REDIS_DB,required"                 json:"db"`
}
