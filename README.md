# Library API in Go

A Library API developed to practice clean architecture with Golang. Feel free to suggest changes and give your opinion on the code, constructive criticism is always welcome
--- 
### What am I using
 - Chi + Postgres + Docker + Clean Architecture (literally Uncle Bob's book)

### How to run
```sh
# Prefer to use make
$ make run

# but you can run the regular entrypoint go run cmd/server/main.go
$ go run cmd/server/main.go
```

### Main Functionalities
 - [ ] Manage books database (inserting, updating, deleting and getting books from it)
 - [ ] Manage current book stock (how many books of each title and who borrowed them)
 - [ ] Manage library customers (keep track of all the books someone has borrowed)

### TODO 
 - [ ] JWT authentication and authorization
 - [ ] Roles for admins (manage the system) and users (search for books)
 - [ ] Universal configuration file, independent of entrypoint used
 - [ ] Only search titles, authors and publishers with exact name, must fix to select when contains substrings too
