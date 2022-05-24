# coffee_api
Basic go rest api observing postgresql patterns with chi router.
Non Standard Library imports: 
  - pq driver github.com/lib/pq
  - chi router github.com/go-chi/chi/v5

This api works with Postgresql, using a table named "product" which follows the patterns in models/products.go for table structure.
id = primary key, autoincrementing, name = varchar, description =  varchar , price = float (I had issues with conversion from go to postgresql money values).

TO DO:
  - Add DeleteHandler
  - Add json validation logic with https://github.com/go-playground/validator
  - Add Swagger and Redoc for documentation
  - Create subroutes
  - Pair with React based front end.
