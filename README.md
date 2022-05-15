# coffee_api
Basic go rest api observing postgresql patterns with chi router

This api works with Postgresql, using a table named "product" which follows the patterns in models/products.go for table structure.
id = primary key, autoincrementing, name = string, description = string, price = float (I had issues with conversion from go to postgresql money values).

TO DO:
  - Add DeleteHandler
  - Add Swagger and Redoc for documentation
  - Create subroutes
  - Pair with React based front end.
