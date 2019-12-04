package conf

type Config struct {
	Addr string
}

func Parse() *Config {
	return &Config{
		Addr: ":8003",
	}
}
