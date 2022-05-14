package config

type UserId string

type ApiKey struct {
	Telegram string
	Site     string
}

type Config struct {
	ApiKey     ApiKey
	Moderators map[UserId]string
	Dsn        string
}
