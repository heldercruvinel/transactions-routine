
# </br>Transactions Routine</br>

The project was created to simulate a credit transactions routine.

## </br>Dependencies

To run this project, you need to install some tools:

1. [Git](https://git-scm.com/downloads)
2. [Docker (Docker Compose is installed together)](https://docs.docker.com/get-started/get-docker/)

## </br>Running The Project

To run this project, you can use pure Docker commands or Docker Compose.

To make our API and PostgreSQL use the same database user and password, we need to set the following environment variables: **`DATABASE_USER`** and **`DATABASE_PASSWORD`**.

Everything inside "[ ]" we need to change for our own values, for example, where we find **`[database_password]`** we need to change for **`my_password_123`**.

For the purpose of correctly running the project, we have to follow these steps:

1. [Cloning the repository.](#clonning-repository)
2. [Using Docker commands.](#docker)
    1. [Setting up the environment variables (Works with Docker Compose too)](#env-docker)
    2. [Running the **`postgresql`** container.](#postgresql-container)
    3. [Build and running the **`api`** container.](#api-container)
3. [Using Docker Compose commands.](#docker-compose)
    1. [Setting up the environment variables](#env-docker-compose)
    2. [Running all services](#services)


<h3 id="clonning-repository"></br>Cloning The Repository</h3>

> </br>
> 
> ```bash
> $ cd ./my/repositories/folder/
> $ git clone git@github.com:heldercruvinel/transactions-routine.git && cd ./transactions-routine
> ```
>
> </br>

<h3 id="docker"></br>Using <bold><code>Docker</code></bold> commands.</h3>

>
> <h4 id="env-docker"></br>- Setting up the environment variables (Works with Docker Compose too)</h4>
>
> <sub>Shell</sub>
> ```bash
> $ export DATABASE_HOST="postgresql"
> $ export DATABASE_USER=[my_database_user]
> $ export DATABASE_PASSWORD=[my_database_password]
> ```
> <sub>Command Prompt</sub>
> ```powershell
> $ setx DATABASE_HOST "postgresql"
> $ setx DATABASE_USER "[my_database_user]"
> $ setx DATABASE_PASSWORD "[my_database_password]"
> ```
> <sub>PowerShell</sub>
> ```powershell
> $ [Environment]::SetEnvironmentVariable("DATABASE_HOST", "postgresql", "User")
> $ [Environment]::SetEnvironmentVariable("DATABASE_USER", "[my_database_user]", "User")
> $ [Environment]::SetEnvironmentVariable("DATABASE_PASSWORD", "[my_database_password]", "User")
> ```
>
> <sub>Creating the docker network</sub>
> ```bash
> $ docker network create -d bridge backend
> ```
>
> <h4 id="postgresql-container"></br>- Running the <bold><code>postgresql</code></bold> container</h4>
> 
> ```bash
> $ docker run --name postgresql -v db:/var/lib/postgresql/data --network=backend -p 5432:5432 -e POSTGRES_DB="financial" -e POSTGRES_PASSWORD=$DATABASE_PASSWORD -e POSTGRES_USER=$DATABASE_USER -d postgres:18.0-alpine3.22
> ```
>
> <h4 id="api-container"></br>- Bulding and running the <bold><code>api</code></bold> container</h4>
>
> ```bash
> ## Inside the cloned repository folder
> $ docker build -t api .
> $ docker run -d --name api --network=backend -p 8080:8080 -e DATABASE_HOST=$DATABASE_HOST -e DATABASE_USER=$DATABASE_USER -e DATABASE_PASSWORD=$DATABASE_PASSWORD -it api api
> ```
>
> </br>

<h3 id="docker-compose"></br>Using <bold><code>Docker Compose</code></bold> commands.</h3>

>
> <h4 id="env-docker-compose"></br>- Setting up the environment variables</h4>
>
> Create a **`.env`** file and set the environmet variables inside it.
> 
> <sub>Shell</sub>
> ```bash
> $ touch .env
> $ echo DATABASE_HOST=postgresql > .env
> $ echo DATABASE_USER=[my_database_user] >> .env
> $ echo DATABASE_PASSWORD=[my_database_password] >> .env 
>
> ```
> <sub>Command Prompt</sub>
> ```powershell
> $ type nul > .env
> $ echo DATABASE_HOST=postgresql >> .env
> $ echo DATABASE_USER=[my_database_user] >> .env
> $ echo DATABASE_PASSWORD=[my_database_password] >> .env
> ```
> <sub>PowerShell</sub>
> ```powershell
> $ New-Item -Path .env -ItemType File -Force
> $ Add-Content -Path .env -Value "DATABASE_HOST=postgresql"
> $ Add-Content -Path .env -Value "DATABASE_USER=[my_database_user]"
> $ Add-Content -Path .env -Value "DATABASE_PASSWORD=[my_database_password]"
> ```
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

## </br>Extra Notes

### </br>Links used during the development

> - [Markdown Cheat-Sheet](https://www.markdownguide.org/cheat-sheet/)
> - [Make a Readme](https://www.makeareadme.com/)
> - [Shields.io](https://shields.io/)
> - [Golang Documentation](https://go.dev/doc/)
> - [Docker Reference](https://docs.docker.com/reference/)
> - [Docker Hub](https://hub.docker.com)
> - [PostgreSQL Documentation](https://www.postgresql.org/docs/)


### </br>Development decisions

> - For this test, I preferred to use the standard library whenever possible to remain independent from external libraries. However, when necessary, mainly in production, I don't see any problem with using external libraries.
> - All the project configuration was created intending to be simple as possible, but considering the diversity in the computer system environments, I used some environment variables in order to make the configuration process easier.

## </br>License

[MIT License](https://choosealicense.com/licenses/mit/)





docker network create -d bridge my-net


