package database

var connection string

func init() {
	connection = "Mysql"
}

func GetDataBase() string {
	return connection;
}