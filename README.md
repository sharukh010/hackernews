# HackerNews GraphQL API (Go)

This project is a simple **GraphQL API** implemented using **Golang**, inspired by the Hacker News application. It served as a learning project to explore the foundations of **GraphQL**, including its schema design, mutations, queries, and integration with authentication.

## 🌱 Objective

The goal of this project was to:

- Understand how **GraphQL** works in a real-world API.
- Learn about **queries**, **mutations**, and **input schemas**.
- Explore how to build GraphQL servers in **Go** using `gqlgen`.
- Implement **JWT-based authentication** and **token refreshing**.
- Use **MySQL** for data persistence with basic user and link models.

## 🚀 Features

- 📌 GraphQL API for managing links and users.
- 🧩 JWT authentication and token refresh functionality.
- 💾 MySQL-backed database using `gorm`.
- 🌐 Built using `gqlgen`, `chi`, and `.env` config via `godotenv`.
- 🛠️ Database migrations with `golang-migrate`.

## 🧠 Learnings

Throughout the project, I gained hands-on experience in:

- Writing **GraphQL schemas** with complex input and output types.
- Setting up **mutations** to create users and links.
- Implementing **query resolvers** to fetch nested data.
- Adding **middleware for JWT auth** and token parsing.
- Structuring a modular **Go backend** project with clear separation of concerns.

## 🗃️ GraphQL Schema Overview

```graphql
type Link {
  id: ID!
  title: String!
  address: String!
  user: User!
}

type User {
  id: ID!
  name: String!
}

type Query {
  links: [Link!]!
}

input NewLink {
  title: String!
  address: String!
}

input RefreshTokenInput {
  token: String!
}

input NewUser {
  username: String!
  password: String!
}

input Login {
  username: String!
  password: String!
}

type Mutation {
  createLink(input: NewLink!): Link!
  createUser(input: NewUser!): String!
  login(input: Login!): String!
  refreshToken(input: RefreshTokenInput): String!
}
````

## 🛠 Tech Stack

* **Language:** Go 1.24.2
* **Frameworks & Libraries:**

  * `gqlgen` - GraphQL server
  * `chi` - HTTP router
  * `jwt-go` - Authentication
  * `godotenv` - Config management
  * `mysql` - Database
  * `golang-migrate` - DB migrations

## 📂 Project Structure

```
github.com/sharukh010/hackernews/
│
├── graph/               # GraphQL schema, resolvers
├── database/            # DB connection and migrations
├── middleware/          # JWT auth middleware
├── models/              # User and Link models
├── utils/               # Helper functions
├── .env                 # Environment variables
├── gqlgen.yml           # gqlgen config
└── main.go              # App entrypoint
```

## 🧪 Running the Project

1. **Clone the repo**

   ```bash
   git clone https://github.com/sharukh010/hackernews.git
   cd hackernews
   ```

2. **Create a `.env` file** with your DB and secret config:

   ```
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=hackernews
   JWT_SECRET=yourjwtsecret
   ```

3. **Run database migrations**

   ```bash
   migrate -path ./migrations -database "mysql://root:password@tcp(localhost:3306)/hackernews" up
   ```

4. **Run the server**

   ```bash
   go run main.go
   ```

5. **Access GraphQL Playground** at `http://localhost:8080/query`

## 🔗 Repository

[github.com/sharukh010/hackernews](https://github.com/sharukh010/hackernews)

## 📚 Reference 

[GraphQL Tutorial](https://www.howtographql.com/graphql-go/0-introduction/)

---

## 🧭 Future Improvements

* Add pagination to link queries.
* Implement user roles and permissions.
* Add logging and metrics (OpenTelemetry ready).

---


