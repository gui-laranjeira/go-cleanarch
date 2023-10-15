INSERT INTO public.books (id_book,title,author,pages,publisher,"year",isbn,created_at,updated_at) VALUES
	 ('3e2dea79-8db7-4f94-9d91-db8d19a9cb27','Clean Code: A Handbook of Agile Software Craftsmanship','Robert C. Martin',464,'Pearson',2008,'978-0132350884','2023-10-07 22:41:47.3376692-03:00','0001-01-01 00:00:00Z'),
	 ('3f02f4aa-dda4-465f-a77d-a3ae5bfaa12f','Neuromancer','William Gibson',336,'Penguin Publishing Group',2000,'9780441007462','2023-10-07 22:45:50.6089105-03:00','0001-01-01 00:00:00Z'),
	 ('99d38b22-59e8-4739-afcc-6c96e1b2cc92','The Hobbit','J. R. R. Tolkien',320,'HarperCollins Publishers',2012,'9780547928227','2023-10-07 22:46:50.1233813-03:00','0001-01-01 00:00:00Z'),
	 ('436aa49a-579d-4a07-a86c-f70509c975a0','Domain-Driven Design: Tackling Complexity in the Heart of Software / Edition 1','Eric Evans',560,'Pearson Education',2003,'9780321125217','2023-10-07 22:47:38.9901889-03:00','0001-01-01 00:00:00Z'),
	 ('cc73ad03-55c8-4c7d-9ef8-53911b486329','Data Structures and Algorithms with Golang','Bhagvan Kommadi',336,'Packt Publishing',2019,'9781789618419','2023-10-07 22:49:04.2975299-03:00','0001-01-01 00:00:00Z'),
	 ('a37a2612-f192-4bf9-9ab1-2196f24ede8f','Cybersecurity','Duane C. Wilson',160,'MIT Press',2021,'9780262542548','2023-10-07 22:50:14.3486669-03:00','0001-01-01 00:00:00Z');

INSERT INTO users (id_user,email,phone,password,first_name,last_name,created_at,updated_at) VALUES
	 ('a4ce9866-c04e-4026-8ea9-d6adf761e2b8','teste2@teste.com','987654321','$2a$10$0C4yIL1XxxxhYqnCtsSfF./Ls.uuMLxh5NrFsnFt71945baHSLTvW','Maria','Santo','2023-10-12 21:10:34.679668','0001-01-01 00:00:00'),
	 ('a4494950-e5f9-4139-8d5e-de498cfcf08b','teste@teste.com','987654321','$2a$10$XTbSu2UgADI8ydhBKWRk4eXqmcNKpCLjrxQz/NvXkONt7t9gRov0q','John','Doe','2023-10-12 21:10:43.495739','0001-01-01 00:00:00');

INSERT INTO costumers (id_costumer,email,phone,address,"document",first_name,last_name,created_at,updated_at,current_book_id) VALUES
	 ('14d6c480-328f-4db9-aae4-aa82136b54de','teste@teste.com','987654321','R. das Pedras, 1','101001001','Antonia','Teste','2023-10-14 23:18:02.05326','0001-01-01 00:00:00',NULL);
