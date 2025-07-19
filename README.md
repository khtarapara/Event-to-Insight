# Incident AI Assistant

A full-stack AI-powered Incident Triage application using:

- **Go (Gin)** for backend
- **React (Vite + Material UI)** for frontend
- **OpenAI/OpenRouter** for AI classification
- **MySQL** for persistence
- **Playwright**, **Vitest**, **Zap**, **Swagger** for full testing and logging

---

## 🧠 Features

- Submit incidents via form
- AI-classified severity and category
- View all incidents in a styled table
- Incident detail screen
- Fully tested: unit, integration, E2E
- Swagger-based API documentation

---

## 🚀 Running the App

### 🔧 1. Backend (Go)

```bash
cd incident-ai-backend
go run main.go
```

Create \`.env\`:

```env
PORT=8080
DB_USER=root
DB_PASS=yourpass
DB_HOST=127.0.0.1:3306
DB_NAME=incident_db
FRONTEND_ORIGIN=http://localhost:5173
OPENROUTER_API_KEY=your-openrouter-key
```

Install deps:

```bash
go get
```

Run:

```bash
go run main.go
```

Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

### 💻 2. Frontend (React)

```bash
cd incident-ai-frontend
npm install
npm run dev
```

---

## 🧪 Testing

### ✅ Backend

```bash

# Unit + integration tests

go test ./tests/...
```

### ✅ Frontend

```bash

# Component/unit tests

npx vitest

# E2E tests with Playwright

npx playwright test
```

---

## 📚 API Endpoints

| Method | Endpoint       | Description                |
| ------ | -------------- | -------------------------- |
| POST   | /incidents     | Create new incident        |
| GET    | /incidents     | List all incidents         |
| GET    | /incidents/:id | Get single incident detail |

Swagger: [localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🧠 AI Integration

AI classification is handled via OpenRouter-compatible LLMs.

You can toggle between:

- \`openai/gpt-3.5-turbo\`
- \`google/gemma-3n-e2b-it:free\`

---

## 📁 Folder Structure Highlights

```
incident-ai-backend/
├── models/
├── controllers/
├── config/
├── routes/
├── utils/logger.go

incident-ai-frontend/
├── src/pages/home/
├── src/pages/incident/
├── src/components/
├── src/api/
├── src/tests/
```

---

## 🧩 Logging

Structured logging with \`log\`. Usage:

```go
logger.Logger.Error("incident created", zap.String("title", incident.Title))
```

---

## 🧪 Test Coverage

| Layer         | Library          |
| ------------- | ---------------- |
| Backend unit  | testify + sqlite |
| API           | httptest + Gin   |
| Frontend unit | RTL + Vitest     |
| E2E           | Playwright       |

---

## 🧠 Author

[Kaniya Tarapara](https://github.com/khtarapara/)
