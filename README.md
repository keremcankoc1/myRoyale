# 👑 myRoyale - Clash Royale REST API Service

A lightweight, modular Go REST API service that fetches and calculates real-time Clash Royale player and card analytics.

## 🚀 Features (Endpoints)

*   `GET /player/{tag}/cards` ➡️ List of player's cards with normalized actual game levels.
*   `GET /player/{tag}/playerStats` ➡️ Player profile details, trophy counts, win/loss stats, and league/season records.

## 🛠️ Quick Start

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/KULLANICI_ADIN/myRoyale.git
    cd myRoyale/Go
    ```

2.  **Set up your environment variables:**
    *   Copy `.env.example` to `.env`
    *   Add your Clash Royale Developer API token into `.env`

3.  **Run the server:**
    ```bash
    go run .
    ```

4.  **Access the API:**
    Open [http://localhost:8080](http://localhost:8080) in your browser or API client.
