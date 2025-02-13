package config

type Settings struct {
	Port      int       `json:"port"`
	Database  Database  `json:"database"`
	Stalcraft Stalcraft `json:"stalcraft"`
}

type Database struct {
	PostgresConnection string `json:"postgresConnection"`
}

type Stalcraft struct {
	DomainApi string `json:"domainApi"`
}
