package config

func HashIdSalt() string {
	return cfg.HashId.Salt
}

func HashIdMinLength() int {
	return cfg.HashId.MinLength
}
