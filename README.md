# Library API in Go

A Library API developed to practice clean architecture with Golang. Feel free to suggest changes and give your opinion on the code, constructive criticism is always welcome
--- 
### What am I using
 - Chi + Viper + Postgres + Docker + Clean Architecture (literally Uncle Bob's book)

### How to run

- The first step is to create a config.toml file following the example with the database environment variables. 
- After that edit the connection string in `/internal/infrastructure/db/connection.go` file to match your run option: docker or make. (Docker use the container host and make use localhost, still working on a better solution for this problem.)

You can run using docker: 

```sh
$ docker compose up
```

If you prefer to use make, you just need to set up the database locally and after that:

```sh
$ make run
```

### Main Functionalities
 - [X] Manage books database (inserting, updating, deleting and getting books from it)
 - [ ] Manage current book stock (how many books of each title and who borrowed them)
 - [ ] Manage library customers (keep track of all the books someone has borrowed)
 - [ ] Browse books (costumer role)

### TODO 
 - [ ] Config docker compose to run postgres and the app
 - [ ] JWT authentication and authorization
 - [ ] Roles for users (authentication needed) and costumers (no authentication needed)
 - [ ] Universal configuration file, independent of entrypoint used
 - [X] Only search titles, authors and publishers with exact name, must fix to select when contains substrings too
 - [ ] Use Redis to cache main books to display
