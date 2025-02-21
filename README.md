# SmartTaskAI
🚀 **AI-Powered Task Management System** with **Next.js & Golang**

SmartTaskAI is a **full-stack AI-powered task management application** that enables users to create, assign, and track tasks in **real-time**. Built with **Golang (Fiber)** for the backend and **Next.js (React + TypeScript)** for the frontend, it features **JWT authentication, WebSockets for real-time updates, and AI-driven task breakdowns using OpenAI/Gemini**.

📌 **Designed for developers and teams looking to optimize workflows with AI-powered task suggestions and seamless real-time collaboration.**

![HomePage](https://github.com/pooulad/nextjs-golang-crud-app/blob/main/images/home.png)

## **🛠️ Technology Stack**

### **🔹 Backend (Golang)**
- **Fiber** - Fast web framework for Golang
- **JWT Authentication** - Secure user sessions
- **WebSockets** - Real-time task updates
- **PostgreSQL / MongoDB** - Database management
- **GORM** - ORM for PostgreSQL integration
- **OpenAI API / Gemini API** - AI-powered task breakdowns
- **Docker & Kubernetes** - For containerized deployment

### **🔹 Frontend (Next.js + TypeScript)**
- **React Hook Form** - Form handling
- **Axios** - API communication
- **React Toastify** - Notifications
- **TailwindCSS** - UI styling
- **NextAuth.js** - Secure authentication
- **WebSockets** - Real-time UI updates

## **📌 Features**
✅ **JWT Authentication** - Secure login & session handling  
✅ **AI-Powered Task Suggestions** - Smart breakdown of tasks using OpenAI/Gemini  
✅ **Real-Time Updates** - WebSockets for live task tracking  
✅ **Task Management Dashboard** - Create, assign, and track tasks dynamically  
✅ **Role-Based Access Control (RBAC)** - Admin/User roles  
✅ **Cloud Deployment** - Backend on **Render/Fly.io**, Frontend on **Vercel**  
✅ **Dockerized Setup** - Run services in containers  
✅ **Slack/Discord Bot Integration** *(optional)* - For notifications  

## **📌 How to Run the Project**

### **Option 1: Run with Docker 🐳**
```bash
docker compose up -d
```
2- Run nextjs
```bash
  cd ./frontend
  yarn install
  yarn run dev
```

### **Option 2: Manual**

1-Create DB based on .env file(you can change it in .env file)

2-In root of source you should run your go project
```bash
  go mod tidy
  go run main.go
```
3-In root frontend directory you should run your next project
```bash
  yarn install
  yarn run dev
```

## API Reference

#### Local address
```bash
  http://127.0.0.1:8080
```
#### All endpoints

```http
  GET /api/users -> get all users
  POST /api/user -> create new user
  PATCH /api/user/:id -> update user
  GET /api/user/:id -> get single user
  DELETE /api/user/:id -> delete user
  POST /api/login -> login to account
```

#### TODO Checklist

This section tracks the progress of implemented features in project.

- [x] Add swagger to project.
- [x] Add Next-auth to project -> frontend directory.
- [x] Add all api response states(success,failure).
- [x] Dockerize project.

## Screenshots

![App SignUp Form](https://github.com/pooulad/nextjs-golang-crud-app/blob/main/images/sign-up.png)

![App SignIn Form](https://github.com/pooulad/nextjs-golang-crud-app/blob/main/images/sign-in.png)




