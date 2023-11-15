package models

type Config struct {
	App APP      `envconfig:"APP"`
	DB  Postgres `envconfig:"POSTGRES"`
}

type APP struct {
	PortAuthGRPC   string `envconfig:"PORT_GRPC_AUTH"`
	PortLoginGRPC  string `envconfig:"PORT_GRPC_LOGIN"`
	PortAccessGRPC string `envconfig:"PORT_GRPC_ACCESS"`
	PortAuthHTTP   string `envconfig:"PORT_HTTP_AUTH"`
	PortDocs       string `envconfig:"PORT_DOCS"`

	AccessSecret          string `envconfig:"ACCESS_SECRET"`
	RefreshSecret         string `envconfig:"REFRESH_SECRET"`
	AccessTTL             string `envconfig:"ACCESS_TTL_MINUTE"`
	RefreshTTL            string `envconfig:"REFRESH_TTL_MINUTE"`
	KeyForHashingPassword string `envconfig:"KEY_FOR_HASHING_PASSWORD"`
}

type Postgres struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Host     string `envconfig:"HOST" required:"true"` // localhost
	Port     string `envconfig:"PORT" required:"true"` //:5432
	Database string `envconfig:"DATABASE" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
}
