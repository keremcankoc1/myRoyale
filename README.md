# 👑 myRoyale - Clash Royale REST API & Analytics Service

A lightweight, modular Go REST API service backed by Microsoft SQL Server (MSSQL) that fetches, caches, and calculates Clash Royale player, card, and deck analytics.

## 🚀 Features & Endpoints

*   `GET /player/{tag}/playerStats` ➡️ Player profile details, trophies, win/loss stats, clan info, and favourite card.
*   `GET /player/{tag}/cards` ➡️ Full collection of player's cards with normalized in-game levels (1–16), evolution status, and tactical attributes (`role`, `is_flying`, `targets_air`, `is_splash`).
*   `GET /player/{tag}/deck` ➡️ Player's active 8-card battle deck with calculated **Average Elixir Cost** and **Average Card Level**.

## 🏗️ Architecture Highlights

*   **Cache-Aside Pattern (5-Min TTL):** Checks MSSQL freshness (`IsPlayerFresh`) before calling the Supercell API to minimize latency and avoid rate limits.
*   **Database Seeding & Normalization:** Automatically seeds static card mechanics (`card_catalog.json` ➡️ `card_catalog` table) on startup and merges them with player data via SQL `INNER JOIN`.

## 🛠️ Quick Start

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/keremcankoc1/myRoyale.git
    cd myRoyale/Go
    ```

2.  **Set up your environment variables:**
    *   Copy `.env.example` to `.env`
    *   Add your Clash Royale Developer API token (`CR_API_TOKEN`) and MSSQL connection string (`DB_DSN`) into `.env`

3.  **Run the server:**
    ```bash
    go run main.go
    ```

4.  **Access the API:**
    Open [http://localhost:8080/player/8GJCL2V09/deck](http://localhost:8080/player/8GJCL2V09/deck) in your browser or API client.
