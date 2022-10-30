module github.com/sdq-codes/maze-app

// +heroku goVersion go1.15
go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/getsentry/sentry-go v0.11.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.4.0
	github.com/mailgun/mailgun-go/v4 v4.6.0
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
