#commands and steps to follow
go mod init github.com/jaipalsinghkhangarot/bookstore
go mod tidy
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/mysql
go get -u github.com/gorilla/mux
go mod tidy
#download xampp for mysql and run mysql in it
# in main folder run
go build
go run main.go