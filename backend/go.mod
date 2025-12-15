module github.com/AntonKhPI2/nim-one-pile

go 1.24.2

require (
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/joho/godotenv v1.5.1
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	golang.org/x/net v0.48.0
	gorm.io/driver/sqlite v1.6.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/text v0.32.0 // indirect
)

replace github.com/AntonKhPI2/nim-backend => .

replace github.com/AntonKhPI2/nim-one-pile => .
