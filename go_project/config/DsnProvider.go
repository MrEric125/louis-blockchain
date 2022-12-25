package config

type DsnProvider interface {
	Dsn() string
}
