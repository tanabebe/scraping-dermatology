package domain

type DbConfig struct {
	DriverName             string
	InstanceConnectionName string
	DatabaseUser           string
	Password               string
	DatabaseName           string
}

type ScrapingList struct {
	Url      string
	Email    string
	Password string
}
