// README.md

# Project Support

###Introduction
CMS is a platform that enables users perform CRUD - Create, Read, Update, Delete - operations on their personal information which includes name as well as other details deemed as important

### Project Support Features

-   Users can create and save a collection of important information about themselves
-   Users can retrieve said information via their id
-   Users can edit and update stored information
-   Users can delete stored information

### Installation Guide

-   Clone the repository [here](https://www.github.com/Huey-Emma/)
-   The develop branch is the most stable branch at any given time. Ensure you are working from it

-   Run `go mod download` to install all dependencies
-   You can either work with PGAdmin or PSQL to access your datasbase
-   Create a `.env` file in the project root folder and add your variables. See `.env.sample` for assistance

### API Endpoints

| HTTP Verbs | Endpoints  | Action                               |
| ---------- | ---------- | ------------------------------------ |
| POST       | `/api`     | Add personal information             |
| GET        | `/api/:id` | Retrieve stored personal information |
| PATCH      | `/api/:id` | Edit and update personal information |
| DELETE     | `/api/:id` | Delete personal information          |

### Technologies Used

-   [Golang]("https://go.dev")
    An open source programming language supported by Google. It allows for installation and management of dependencies and communication with databases.

-   [Postgres]("https://www.postgresql.org")
    A powerful open source object relational database that has a strong reputation for reliability, feature robustness and performance

### License

This project is available for use under the MIT License
