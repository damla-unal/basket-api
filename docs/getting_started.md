
###Installing Postgres

1. Install and run Postgresql
   1. if you use Mac, you can follow this steps : 
  [Mac Setup PostgreSQL](https://sourabhbajaj.com/mac-setup/PostgreSQL/)
2. Create database that named basket_api
   >createdb basket_api
3. Run migrations from Makefile
   > make migrateup
   1. You can see the details of golang-migrate: [DB migration in Go](https://medium.com/geekculture/db-migration-in-go-lang-d325effc55de)
     