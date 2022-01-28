module github.com/sajeelwaien/gopro

go 1.16

replace github.com/sajeelwaien/gopro/database => ./Database

replace github.com/sajeelwaien/gopro/migrations => ./Migrations

replace github.com/sajeelwaien/gopro/models => ./Models

replace github.com/sajeelwaien/gopro/schemas => ./Schemas

require (
	github.com/cosmtrek/air v1.27.6 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/graphql-go/graphql v0.8.0 // indirect
	github.com/jackc/pgx/v4 v4.14.0 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/sajeelwaien/gopro/database v1.0.0
	github.com/sajeelwaien/gopro/migrations v1.0.0
	github.com/sajeelwaien/gopro/models v1.0.0
	github.com/sajeelwaien/gopro/schemas v1.0.0
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	gorm.io/driver/postgres v1.2.2 // indirect
	gorm.io/gorm v1.22.3 // indirect
)
