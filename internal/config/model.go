package config

type Config struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

type Server struct {
	Host string
	Port string
}
type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Timezone string
}

type Jwt struct {
	RefreshKey string
	AccessKey  string
}
