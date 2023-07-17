package database

var connection string

func init() {
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
