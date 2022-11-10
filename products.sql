CREATE TABLE products (
id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
name TEXT NOT NULL,
description TEXT,
price MONEY NOT NULL);

# docker run -it --name coffee_DB -p 5444:5432 -e POSTGRES_PASSWORD=mypassword -e POSTGRES_USER=coffeeguy -e POSTGRES_DB=products -d postgres