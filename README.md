# Porsche API

<img src="https://logodownload.org/wp-content/uploads/2021/02/porsche-logo-0.png"/>

Free fake API, written with <a href="https://go.dev/" target="_blank">Golang</a> for testing and prototyping.

## Why ?
Are you tired of spending valuable time registering for complex APIs when all you need is some data for testing or prototyping?

Introducing Porsche API – a lightweight Golang-based API, designed to provide quick and easy access to dummy data without the hassle of registration or complex API documentation.

## Features
* No registration
* Basic API
* Cross-domain
* Supports GET, POST, PUT, DELETE
* HTTP
* Compatible with React, Angular, Vue, Ember, ...

## Set up locally

Before you dive in, make sure to set up your environment variables by following these steps:

1. Clone the project: Open your terminal and clone our project repository by running the following command:
`git clone https://github.com/vitolinho/porsche-api.git`

2. **Grab the `.env.example`**: Head over to your project directory and find the `.env.example` file. This serves as a template for your environment variables.

3. **Fill in the `.env` file**: Duplicate the `.env.example` file and rename it to `.env`. Open it up and fill in the required values based on your setup. Don't worry, the example file provides hints on what each variable should be set to.

4. **Ready to Launch!**: Once your `.env` file is properly filled, you're almost ready to go!

5. **Start Docker**: If you haven't already, make sure you have Docker installed. [Install Docker](https://docs.docker.com/get-docker/) if you haven't already.

6. **Run the command**: In your terminal, simply execute `make up` or `docker-compose up -d --build` to start up your local environment.

## Ressources

Porsche-api comes with a set of 1 common resource:

`/cars` 50 cars

## Routes

**GET, POST, PUT, DELETE** HTTP methods are supported. You can use http for your requests.<br>

*GET* `/api/v1/cars`<br>

*GET* `/api/v1/cars/1`<br>

*POST* `/api/v1/cars`<br>

*PUT* `/api/v1/cars/1`<br>

*DELETE* `/api/v1/cars/1`<br>

## Technical stack

[Golang](https://go.dev/) Programing Language<br>
[Fiber](https://docs.gofiber.io/) Web Framework<br>
[GORM](https://gorm.io/) ORM library
