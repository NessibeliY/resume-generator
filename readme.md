# 📝 Resume Generator (Go + React + OpenAI)

An AI-powered resume generator built with a Go backend (Gin) and a React frontend. It uses OpenAI to help users automatically generate personalized resume content and download it as a PDF.

---

## 📁 Project Structure

resume-builder/
├── backend/          # Go backend (API, OpenAI, PDF)
│   ├── cmd/              # main.go entrypoint
│   ├── internal/
│   │   ├── config/       # .env loader
│   │   ├── handler/      # HTTP endpoints
│   │   ├── service/      # Business logic (AI, PDF)
│   │   ├── models/       # DTOs / structs
│   │   ├── repository/   # DB (optional)
│   ├── .env
│   └── go.mod
│
├── frontend/         # React + Vite
│   ├── src/
│   │   ├── pages/        # Pages (Form, Resume)
│   │   ├── components/   # Reusable UI
│   │   └── api/          # API calls (Axios)
│   ├── package.json
│   └── vite.config.js
│
└── README.md

---

## 🚀 Features

- ✨ Fill out a simple form with your info
- 🤖 Let AI generate resume content for you using OpenAI
- 📄 Export generated content as a downloadable PDF
- 🧠 Clean codebase with modular architecture

---

## 🛠️ Tech Stack

- **Backend:** Go, Gin, OpenAI API
- **Frontend:** React, Vite, Tailwind CSS
- **Other:** dotenv, CORS, REST API

---

## 🔐 Environment Setup

### 1. Clone the repository

```bash
git clone git@github.com:NessibeliY/resume-generator.git
cd resume-generator


⸻

2. Backend Setup

✅ Requirements
	•	Go 1.20+
	•	OpenAI API Key (paid)


🔐 Create .env file inside backend/:

OPENAI_API_KEY=your_openai_key
PORT=:8080

▶️ Run the backend

cd backend
go run cmd/main.go



⸻

3. Frontend Setup

📦 Install dependencies

cd frontend
npm install

▶️ Run frontend

npm run dev

Frontend will be available at:
http://localhost:5173

Make sure the backend is running on port 8080. The frontend uses a proxy to forward API requests.

⸻

🔗 API Example

POST /generate

Request:

{
  "name": "John Doe",
  "skills": ["Go", "Docker", "PostgreSQL"],
  "experience": "Backend Developer with 5 years experience..."
}

Response:

{
  "resumeText": "Generated resume content here..."
}



