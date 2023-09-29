package postgres

type Config struct {
	Database string `config:"DATABASE_NAME" yaml:"database"`
	User     string `config:"DATABASE_USER" yaml:"user"`
	Password string `config:"DATABASE_PASSWORD" yaml:"password"`
	Host     string `config:"DATABASE_HOST" yaml:"host"`
	Port     int    `config:"DATABASE_PORT" yaml:"port"`
	Retries  int    `config:"DB_CONNECT_RETRY" yaml:"retries"`
	PoolSize int    `config:"DB_POOL_SIZE" yaml:"pool_size"`
}
