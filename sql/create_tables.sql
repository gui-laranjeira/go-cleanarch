-- Creation of books table
CREATE TABLE IF NOT EXISTS books (
  id_book varchar(100) NOT NULL,
  title varchar(100) NOT NULL,
  author varchar(50) NOT NULL,
  pages int NOT NULL,
  publisher varchar(50) NOT NULL,
  year int,
  isbn varchar(50),
  created_at varchar(100) NOT NULL,
  updated_at varchar(100),
  PRIMARY KEY (id_book)
);

-- Creation of users table
CREATE TABLE IF NOT EXISTS users (
  id_user varchar(100) NOT NULL,
  email varchar(50) NOT NULL,
  phone varchar(20),
  password varchar(100) NOT NULL,
  first_name varchar(30) NOT NULL,
  last_name varchar(30) NOT NULL,
  created_at varchar(100) NOT NULL,
  updated_at varchar(100),
  PRIMARY KEY (id_user)
);

-- Creation of costumers table
CREATE TABLE IF NOT EXISTS costumers (
  id_costumer varchar(100) NOT NULL,
  email varchar(50) NOT NULL,
  phone varchar(20) NOT NULL,
  adress varchar(100) NOT NULL,
  document varchar(50) NOT NULL,
  first_name varchar(30) NOT NULL,
  last_name varchar(30) NOT NULL,
  created_at varchar(100) NOT NULL,
  updated_at varchar(100),
  current_book_id varchar(100),
  PRIMARY KEY (id_costumer)
);

CREATE TABLE IF NOT EXISTS borrow_history (
    id_costumer varchar(100) NOT NULL,
    id_book varchar(100) NOT NULL,
    id_stock varchar(100) NOT NULL,
    borrowed_date varchar(100) NOT NULL,
    due_date varchar(100) NOT NUll,
    returned char(1) not null
);

CREATE TABLE IF NOT EXISTS stock (
    id_stock_entry varchar(100) NOT NULL,
    id_book varchar(100) NOT NULL,
    available char(1) NOT NULL
);