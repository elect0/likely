# **Likely**

### A real-time "Who's More Likely To" party game.

\<p align="center"\>
\<img src="[https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge\&logo=go](https://www.google.com/search?q=https://img.shields.io/badge/Go-1.22%2B-00ADD8%3Fstyle%3Dfor-the-badge%26logo%3Dgo)" alt="Go Version"\>
\<img src="[https://img.shields.io/badge/SvelteKit-4.0+-FF3E00?style=for-the-badge\&logo=svelte](https://www.google.com/search?q=https://img.shields.io/badge/SvelteKit-4.0%2B-FF3E00%3Fstyle%3Dfor-the-badge%26logo%3Dsvelte)" alt="SvelteKit Version"\>
\<img src="[https://img.shields.io/badge/PostgreSQL-16-336791?style=for-the-badge\&logo=postgresql](https://www.google.com/search?q=https://img.shields.io/badge/PostgreSQL-16-336791%3Fstyle%3Dfor-the-badge%26logo%3Dpostgresql)" alt="PostgreSQL"\>
\<img src="[https://img.shields.io/badge/Redis-7.2-DC382D?style=for-the-badge\&logo=redis](https://www.google.com/search?q=https://img.shields.io/badge/Redis-7.2-DC382D%3Fstyle%3Dfor-the-badge%26logo%3Dredis)" alt="Redis"\>
\<img src="[https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge](https://www.google.com/search?q=https://img.shields.io/badge/License-MIT-yellow.svg%3Fstyle%3Dfor-the-badge)" alt="License: MIT"\>
\</p\>

**Likely** is a multiplayer web game where friends can join a private room and vote on who is "more likely to" do something based on daily, AI-generated questions. It features live result updates and a group chat, all powered by a high-performance Go backend designed for speed and scalability.

-----

## ✨ Features

  * **Private Game Rooms:** Create a room and invite friends with a unique, shareable code.
  * **Real-time Voting:** Cast your vote and see the results update live for everyone in the room.
  * **Live Group Chat:** Chat with your friends in the game room while you play.
  * **AI-Generated Questions:** Fresh, interesting questions are delivered to your room daily.
  * **Question History:** View the results of questions from the past 7 days.
  * **Secure Authentication:** Email and password-based user accounts.

-----

## 🚀 Tech Stack

This project uses a modern, decoupled architecture to ensure high performance and scalability.

  * **Frontend:** SvelteKit & Tailwind CSS (Deployed on Vercel)
  * **Backend:** Go (Echo Framework) & WebSockets (Deployed on Google Cloud Platform)
  * **Database:** PostgreSQL (Cloud SQL)
  * **Cache:** Redis (Cloud Memorystore) for real-time updates & caching.

-----

## 🏗️ Backend Architecture

The backend is built using a **Layered Architecture** to ensure a clear separation of concerns, making the application highly testable and maintainable.

  * **`handler`:** Manages incoming HTTP and WebSocket requests. Contains no business logic.
  * **`service`:** The core of the application, containing all business logic.
  * **`repository`:** The only layer that communicates with the database and cache.
  * **`domain`:** Contains the core data structures used throughout the application.

-----

## 🛠️ Local Development Setup

To run the backend server locally, please ensure you have **Go 1.22+** and **Docker** installed.

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/your-repo-name.git
    cd likely-game-project/backend
    ```

2.  **Set up Environment Variables:**

      * Copy the example `.env` file:
        ```bash
        cp .env.example .env
        ```
      * Edit the new `.env` file with your local configuration. See the table below for details.

3.  **Start Local Database & Cache:**

      * A `docker-compose.yml` file is included for convenience. Start local PostgreSQL and Redis containers:
        ```bash
        docker-compose up -d
        ```

4.  **Run Database Migrations:**

      * (We will add a migration tool like `golang-migrate` here later)

5.  **Install Dependencies & Run:**

    ```bash
    go mod tidy
    go run ./cmd/server/main.go
    ```

    The server will start on the port specified in your `.env` file (default: `8080`).

-----

## ⚙️ Environment Variables

The `.env` file is used to configure the application.

| Variable | Description | Default |
| --- | --- | --- |
| `PORT` | The port for the HTTP server to listen on. | `8080` |
| `DB_USER` | The username for the PostgreSQL database. | `user` |
| `DB_PASSWORD` | The password for the PostgreSQL database. | `password` |
| `DB_HOST` | The host address of the PostgreSQL database. | `localhost` |
| `DB_PORT` | The port for the PostgreSQL database. | `5432` |
| `DB_NAME` | The name of the database to use. | `likely_db` |
| `REDIS_ADDR` | The address of the Redis instance. | `localhost:6379` |
| `JWT_SECRET` | A secret key for signing authentication tokens. | `your-secret` |

-----
-----

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.
