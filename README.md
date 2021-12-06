# Golang-fizzbuzz

Golang Fizzbuzz. This README is made up of two parts: the first one will give you basic instructions to easily install this application on your local device and be able to take part in the project. The other part is focused on our exercise.

## Project README

### Dependencies

Docker 20.10  
Docker-compose 2.1  
(Golang 1.17)

### Project architecture

The project is built on top of this structure:

- **Application:** Business actions layer
  - **Command:** DTO struct to verify the command required data
  - **Handler:** Service which handle command
- **Core:** Application entry points
- **Driver:** Connector interfaces
- **Infrastructure:** External services layer
- **Presentation:** External system interaction layer, here API entry points

### First installation

Every needed environment variables are located in the .env file. The file is commited for the exercise so you can easily run the project, but I usually ignore it and provide a .env.example or .env.local commited file.

You can install project with:

```bash
$ make install
```

Access to your local service with the following ports:

- **HTTP API:** 3000

### Technical informations

You can update container for dev with:

```bash
$ make update
```

You can run test with:

```bash
$ make test || make test-cover
```

You can display app logs with:

```bash
$ make logs
```

You can have a bash command line interface with:

```bash
$ make cli
```

To get list of all command on your terminal:

```bash
$ make || make help
```

### Simple API DOC

GET /fizzbuzz

5 query parameters accepted: int1 (int), int2 (int), limit (int), str1 (string), str2 (string)

Curl request example with query parameters:

```bash
$ curl --location --request GET 'localhost:3000/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz'
```

---

GET /statistic

No query parameters accepted

Curl request example:

```bash
$ curl --location --request GET 'localhost:3000/statistic'
```

## Exam

### Exercice

Fizz buzz is a group word game where players take turns to count incrementally, replacing any number divisible by three with the word "fizz", and any number divisible by five with the word "buzz".

Your goal is to implement a ready for production and easy to maintain web server. It should expose a REST API endpoint that:

- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

Moreover, expose a statistics endpoint allowing users to know what the most frequent request is, that:

- Accept no parameter.
- Return the most used request five parameters, as well as the number of hits for this request.

### Implementation

The problematic of the exercise is technically focused on making a Golang API which provides JSON endpoints and can store data. I made for this project a simple architecture which allows you to easily maintain and expand it.

I went ahead with my own architecture to highlight one of the ways I like to structure a project, but I used to work with the classic [Go Project Layout](https://github.com/golang-standards/project-layout). This is why I choose to use net/http over gin-gonic/gin (or other) to run a vanilla http server (but I often use gin notably in a professional environment).

I opted for redis to store our needed data which is a good and simple database caching (free / open source / fast and high performance) for our case. We can easily switch to another database storage by creating a new store infrastructure file (for PostgresSQL, MongoDB, Algolia, whatever we want) and updating the implementation on the concerned presentation.

I provided an alpine-golang Docker image with required test tools for development and a minimal Docker image that we can use for production. For development mode, notice that the `make update` command can be triggered at any moment or automatically by nodemon, air, hotswap or others.

To facilitate the environment variables management for the exercise I used the joho/godotenv dependency. Tests have only been done for usecase to facilitate the exercice, as Infrastructure packages needs an implementation strategy to be correctly tested.
