## SIMPLE CRUD MADE IN EDUCATIONAL PURPOSES

API for managing todos and users. 

### Setup

1. Clone this repository by:

```git clone https://github.com/Blthrst/go-todos```

2. Setup your `.env` file with at least 3 environment variables:

   - `MYSQL_USER` - MySQL user.
   - `MYSQL_PWD` - password for MySQL user.
   - `MYSQL_DBNAME` - database name (`go_todos` as default). Should be a name of existing database.

   **Important**. You should create your own database before starting the program.

   To create database just open your MySQL Workbench or connect to mysql-server by `mysql -u <username> -p` and type:

   ```create database if not exists <your_db_name>;```.

   Then check your databases list:

   ```show databases;```

After these steps, you can run program by `go run .` or compile it and run executable file.

HTTP server will listen `localhost:4545` as default.

### Routes

#### Users

- /users/ - get all users
- /users/create/ - create users
- /users/delete - delete users
- /users/update - update one user
- /users/get - get one user

#### Todos

- /todos/ - get all todos
- /todos/create/ - create todos
- /todos/delete - delete todos
- /todos/update - update one todos

You can find `User` and `Todo` structs at `./model/structs.go` 