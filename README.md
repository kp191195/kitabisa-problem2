# HOW TO USE
1. This API Service will be use PostgreSQL
2. Please run the structure.sql into your freshly created database
3. execute -> make copy-env-debug
4. Please edit .env.yml file inside folder config to your server configuration
5. execute -> make run (to start API SERVICE)
6. I prepare the postman collection for testing the API Service, you can import KB-PROBLEM2.postman_collection.json to your POSTMAN

# FOR TEST PURPOSE
1. Unit test will be need separate database
2. You can configure the database for unit test in config/.env.yml on database-test
3. make test for running unit test