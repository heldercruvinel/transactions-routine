
# </br>Transactions Routine</br>

The project was created to simulate a credit transactions routine.

## </br>Dependencies

To run this project, you need to install some tools:

1. [Git](https://git-scm.com/downloads)
2. [Docker (Docker Compose is installed together)](https://docs.docker.com/get-started/get-docker/)

## </br>Running The Project

To run this project, you can use pure Docker commands or Docker Compose.

To make our API and PostgreSQL use the same database user and password, we need to set the following environment variables: **`DATABASE_HOST`**, **`POSTGRES_DB`**, **`POSTGRES_USER`** and **`POSTGRES_PASSWORD`**.

Everything inside "[ ]" we need to change for our own values, for example, where we find **`[database_password]`** we need to change for **`my_password_123`**.

For the purpose of correctly running the project, we have to follow these steps:

1. [Cloning the repository.](#clonning-repository)
2. [Setting up the environment variables](#envs)
3. [Starting the application with start.sh.](#startsh)
4. [Running the unit tests with coverage with tests.sh.](#testssh)
5. [Using Docker Compose commands.](#docker-compose)
    2. [Running all services](#services)
6. [Using Docker commands.](#docker)
    1. [Setting up the environment variables (Works with Docker Compose too)](#env-docker)
    2. [Running the **`postgresql`** container.](#postgresql-container)
    3. [Build and running the **`api`** container.](#api-container)
7. [Testing the api endpoints](#endpoints)
    1. [Insert a Account](#insert-account)
    2. [Get a Account](#get-account)
    3. [Insert a Transaction](#insert-transaction)


<h3 id="clonning-repository"></br>Cloning The Repository</h3>

> </br>
> 
> ```bash
> $ cd ./my/repositories/folder/
> $ git clone git@github.com:heldercruvinel/transactions-routine.git && cd ./transactions-routine
> ```
>
> </br>

<h3 id="envs"></br>Setting up the environment variables</h4>

> </br>
>
> Create a **`.env`** file and set the environmet variables inside it.
> 
> <sub>Shell</sub>
> ```bash
> $ touch .env
> $ echo DATABASE_HOST=postgresql > .env
> $ echo POSTGRES_DB=financial >> .env
> $ echo POSTGRES_USER=[my_database_user] >> .env
> $ echo POSTGRES_PASSWORD=[my_database_password] >> .env 
>
> ```
> <sub>Command Prompt</sub>
> ```powershell
> $ type nul > .env
> $ echo DATABASE_HOST=postgresql >> .env
> $ echo POSTGRES_DB=financial >> .env
> $ echo POSTGRES_USER=[my_database_user] >> .env
> $ echo POSTGRES_PASSWORD=[my_database_password] >> .env
> ```
> <sub>PowerShell</sub>
> ```powershell
> $ New-Item -Path .env -ItemType File -Force
> $ Add-Content -Path .env -Value "DATABASE_HOST=postgresql"
> $ Add-Content -Path .env -Value "POSTGRES_DB=financial"
> $ Add-Content -Path .env -Value "POSTGRES_USER=[my_database_user]"
> $ Add-Content -Path .env -Value "POSTGRES_PASSWORD=[my_database_password]"
> ```
>
> </br>

<h3 id="startsh"></br>Starting the application with <bold><code>start.sh</code></bold></h3>

> </br>
> 
> ```bash
> $ ./start.sh
> ```
>
> </br>

<h3 id="testssh"></br>Running the unit tests and coverage check with <bold><code>tests.sh</code></bold></h3>

> </br>
> 
> ```bash
> $ ./tests.sh
> ```
>
> </br>

<h3 id="docker-compose"></br>Using <bold><code>Docker Compose</code></bold> commands.</h3>

>
> <h4 id="services"></br>- Running all services</h4>
>
> When executing this command, the **`api`** service will wait the **`postgresql`** service automatically.
> 
> ```bash
> $ docker compose up -d
> ```
> 
> </br>

<h3 id="docker"></br>Using <bold><code>Docker</code></bold> commands.</h3>

>
> <sub>Creating the docker network</sub>
> ```bash
> $ docker network create -d bridge backend
> ```
>
> <h4 id="postgresql-container"></br>- Running the <bold><code>postgresql</code></bold> container</h4>
> 
> ```bash
> $ docker run -d --env-file ./.env --name postgresql -v db:/var/lib/postgresql/data --network=backend -p 5432:5432 postgres:18.0-alpine3.22
> ```
>
> <h4 id="api-container"></br>- Bulding and running the <bold><code>api</code></bold> container</h4>
>
> ```bash
> ## Inside the cloned repository folder
> $ docker build -t api .
> $ docker run -d --env-file ./.env --name api --network=backend -p 8080:8080 -it api api
> ```
>
> </br>


<h3 id="endpoints"></br>Testing the api endpoints</h3>

> </br>
>
> <h4 id="insert-account"></br>- Insert a Account</h4>
> 
> ```curl
> $ curl --location 'http://localhost:8080/accounts/' \
>   --header 'Content-Type: application/json' \
>   --data '{
>   	"account_code": "0009992229x"
>   }'
> ```
>
>
> <h4 id="get-account"></br>- Get a Account</h4>
> 
> ```curl
> $ curl --location 'http://localhost:8080/accounts/c8ef914e-c492-4a43-b3d3-00dac9676ced'
> ```
>
>
> <h4 id="insert-transaction"></br>- Insert a Transaction</h4>
> 
> ```curl
> $ curl --location 'http://localhost:8080/transactions/' \
>   --header 'Content-Type: application/json' \
>   --data '{
>   	"operation_id": 3,
>       "account_id":"351ecd0e-991b-45a3-984a-ee577736cb18",
>       "amount":100.00
>   }'
> ```
>
> </br>

## </br>Extra Notes

### </br>Links used during the development

> </br>
>
> - [Markdown Cheat-Sheet](https://www.markdownguide.org/cheat-sheet/)
> - [Make a Readme](https://www.makeareadme.com/)
> - [Shields.io](https://shields.io/)
> - [Golang Documentation](https://go.dev/doc/)
> - [Docker Reference](https://docs.docker.com/reference/)
> - [Docker Hub](https://hub.docker.com)
> - [PostgreSQL Documentation](https://www.postgresql.org/docs/)
>
> </br>


### </br>Development decisions

> </br>
>
> - For this test, I preferred to use the standard library whenever possible to remain independent from external libraries. However, when necessary, mainly in production, I don't see any problem with using external libraries.
> - All the project configuration was created intending to be simple as possible, but considering the diversity in the computer system environments, I used some environment variables in order to make the configuration process easier.
> - Architecture Desing based on SOLID, Clean Architecture, KISS, DRY and others best pratices.
>
> </br>

## </br>License

[MIT License](https://choosealicense.com/licenses/mit/)


