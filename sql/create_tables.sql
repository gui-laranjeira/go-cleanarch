-- Creation of books table
CREATE TABLE IF NOT EXISTS books (
  id_book VARCHAR(100) NOT NULL,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(50) NOT NULL,
  pages INT NOT NULL,
  publisher VARCHAR(50) NOT NULL,
  year INT,
  isbn VARCHAR(50),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  PRIMARY KEY (id_book)
);

-- Creation of users table
CREATE TABLE IF NOT EXISTS users (
  id_user VARCHAR(100) NOT NULL,
  email VARCHAR(50) NOT NULL,
  phone VARCHAR(20),
  password VARCHAR(100) NOT NULL,
  first_name VARCHAR(30) NOT NULL,
  last_name VARCHAR(30) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  PRIMARY KEY (id_user)
);

-- Creation of costumers table
CREATE TABLE IF NOT EXISTS costumers (
  id_costumer VARCHAR(100) NOT NULL,
  email VARCHAR(50) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  address VARCHAR(100) NOT NULL,
  document VARCHAR(50) NOT NULL,
  first_name VARCHAR(30) NOT NULL,
  last_name VARCHAR(30) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  current_book_id VARCHAR(100),
  PRIMARY KEY (id_costumer)
);

CREATE TABLE IF NOT EXISTS borrow_history (
    id_costumer VARCHAR(100) NOT NULL,
    id_book VARCHAR(100) NOT NULL,
    id_stock VARCHAR(100) NOT NULL,
    borrowed_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NUll,
    returned CHAR(1) not null
);

CREATE TABLE IF NOT EXISTS stock (
    id_stock_entry VARCHAR(100) NOT NULL,
    id_book VARCHAR(100) NOT NULL,
    available CHAR(1) NOT NULL,
    id_costumer VARCHAR(100)
);