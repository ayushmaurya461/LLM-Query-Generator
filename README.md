# 🧠 LLM Query Generator (Go + Ollama)

This is a lightweight Go server that uses local LLMs (like Mistral, Phi, etc. via [Ollama](https://ollama.com/)) to **generate SQL or MongoDB queries** from natural language prompts and your database schema.

> You control your data. The model only **returns** the query — it does **not execute** anything.

---

## 🚀 Features

- Query generation via local LLMs (runs on your machine via [Ollama](https://ollama.com))
- Supports **MongoDB** and **SQL**
- Simple JSON API
- CLI support or `curl` for fast interaction
- No data leaves your machine

---

## 📦 Prerequisites

- Go >= 1.20
- [Ollama](https://ollama.com/download) (installed and running locally)
- Models like:
  - `ollama run mistral`
  - `ollama run phi`
  - `ollama run codellama`

---

## 📁 Project Structure

llm-query-generator/
├── main.go # Fiber server
├── handler/ # API handler
├── ollama_client/ # Makes calls to local Ollama instance
├── go.mod / go.sum



---

## ⚙️ How to Run

1. Start your model in the background:

```bash
ollama run mistral

go run main.go

```

### MongoDB Example
```
curl -s -X POST http://localhost:3000/generate-query \
  -H "Content-Type: application/json" \
  -d '{
    "schema": "users: { _id: ObjectId, name: String, email: String }, orders: { _id: ObjectId, user_id: ObjectId, amount: Number }",
    "prompt": "Generate a MongoDB query to list users who spent more than $500. Only return the query. No explanation.",
    "model": "mistral"
  }' | jq -r .query
```
Sample Output
```
db.users.aggregate([
  {
    $lookup: {
      from: "orders",
      localField: "_id",
      foreignField: "user_id",
      pipeline: [
        { $match: { amount: { $gt: 500 } } },
        { $group: { _id: "$user_id", total: { $sum: "$amount" } } },
        { $project: { _id: 1, total: 1, _iduser: "$_id" } }
      ],
      as: "orders"
    }
  },
  { $unwind: "$orders" },
  { $group: { _id: "$_id", total: { $sum: "$orders.total" } } },
  { $match: { total: { $gt: 500 } } }
]);
```
### SQL Example
```
curl -X POST http://localhost:3000/generate-query \
  -H "Content-Type: application/json" \
  -d '{
    "schema": "users(id INT, name TEXT, email TEXT), orders(id INT, user_id INT, amount DECIMAL)",
    "prompt": "List users who spent more than $500",
    "model": "mistral"
  }'
```
Sample Output
```
{
  "query": "SELECT users.name FROM users JOIN orders ON users.id = orders.user_id WHERE orders.amount > 500;"
}
```
🛠️ API Endpoint
POST /generate-query

Request JSON:

{
  "schema": "table/collection schema here",
  "prompt": "natural language query",
  "model": "mistral" // or phi, codellama etc.
}
