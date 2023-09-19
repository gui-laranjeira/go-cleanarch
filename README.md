# Library API in Go

A Library API developed to practice clean architecture with Golang. Feel free to suggest changes and give your opinion on the code, constructive criticism is always welcome
--- 
### What am I using
 - Chi + Postgres + Docker + Clean Architecture (literally Uncle Bob's book)

### How to run
```sh
# Prefer to use make
$ make run

```

### Main Functionalities
 - [X] Manage books database (inserting, updating, deleting and getting books from it)
 - [ ] Manage current book stock (how many books of each title and who borrowed them)
 - [ ] Manage library customers (keep track of all the books someone has borrowed)

### TODO 
 - [ ] Config docker compose to run postgres and the app
 - [ ] JWT authentication and authorization
 - [ ] Roles for admins (manage the system) and users (search for books)
 - [ ] Universal configuration file, independent of entrypoint used
 - [X] Only search titles, authors and publishers with exact name, must fix to select when contains substrings too
