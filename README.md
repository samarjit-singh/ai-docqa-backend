# AI Document Q&A Backend

A Go-based backend service that enables users to upload documents (PDF/TXT) and ask questions about their content using AI. The service processes the documents, extracts their content, and uses the Gemini AI model to provide relevant answers.

## Features

- Document upload support (PDF and TXT files)
- Text extraction from documents
- AI-powered question answering
- Query history tracking
- RESTful API endpoints
- Built with Go and Fiber framework
- PostgreSQL database with Prisma ORM

## Prerequisites

- Go 1.x
- PostgreSQL database
- Node.js (for Prisma)
- Gemini API key

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd ai-docqa-backend
```

2. Install dependencies:

```bash
go mod download
```

3. Install Prisma CLI:

```bash
npm install -g prisma
```

4. Set up environment variables by creating a .env file:

```plaintext
DATABASE_URL="postgresql://username:password@localhost:5432/your_database_name?schema=public"
GEMINI_API_KEY="your-gemini-api-key"
PORT=8080
```

5. Generate Prisma client:
```bash
npx prisma generate
````

6. Run database migrations:

```bash
npx prisma migrate dev
```

## Running the Application

Start the server:

```bash
go run main.go
```

The server will start on the configured port (default: 8080).

## API Endpoints

### Upload Document and Ask Question

```plaintext
POST /query
Content-Type: multipart/form-data

Parameters:
- document: File (PDF or TXT)
- question: String
```

Response:

```json
{
  "filename": "example.pdf",
  "question": "What is this document about?",
  "answer": "AI-generated answer based on document content"
}
```

### View Query History

```plaintext
GET /history
```

Returns a list of previous queries and their answers.
