package database

func Setup(driver string) {
	dbType := driver
	if dbType == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}

	if dbType == "postgres" {
		var db = new(PgSql)
		db.Setup()
	}
}
