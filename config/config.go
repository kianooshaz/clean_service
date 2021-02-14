package config

var cfg *Config

type Config struct {
	Server
	Database
	HashId
	Authentication
}

func SetConfig(c *Config) {
	cfg = c
}

type Authentication struct {
	AccessExpirationInMinute        int    `yaml:"access_expiration_in_minute" json:"-"`
	ActivateExpirationInMinute      int    `yaml:"activate_expiration_in_minute" json:"-"`
	PasswordResetExpirationInMinute int    `yaml:"password_reset_expiration_in_minute" json:"-"`
	RefreshExpirationInMinute       int    `yaml:"refresh_expiration_in_minute" json:"-"`
	JwtSecret                       string `yaml:"jwt_secret" envconfig:"ORDER_JWT_SECRET" json:"-"`
}

type Server struct {
	Port int    `yaml:"port" envconfig:"CLEAN_SERVICE_SERVER_PORT" json:"-"`
	Host string `yaml:"host" envconfig:"CLEAN_SERVICE_SERVER_HOST" json:"-"`
}

type Database struct {
	Username string `yaml:"username" envconfig:"CLEAN_SERVICE_DB_USERNAME" json:"-"`
	Password string `yaml:"password" envconfiq:"CLEAN_SERVICE_DB_PASSWORD" json:"-"`
	DBName   string `yaml:"db_name" envconfig:"CLEAN_SERVICE_DB_NAME" json:"-"`
	Host     string `yaml:"host" envconfig:"CLEAN_SERVICE_DB_HOST" json:"-"`
	Port     string `yaml:"port" envconfig:"CLEAN_SERVICE_DB_PORT" json:"-"`
	Charset  string `yaml:"charset" json:"-"`
	SSLMode  string `yaml:"ssl_mode" envconfig:"CLEAN_SERVICE_DB_SSL_MODE" json:"-"`
	Timezone string `yaml:"timezone" envconfig:"CLEAN_SERVICE_DB_TIMEZONE" json:"-"`
}

type HashId struct {
	Salt      string `envconfig:"CLEAN_SERVICE_HASH_ID_SALT" json:"-"`
	MinLength int    `yaml:"min_length" json:"-"`
}
