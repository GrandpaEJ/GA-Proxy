# PteroBill AI Proxy

This is a lightweight Go-based proxy server designed to be hosted on platforms like **Render**, **Vercel**, or **Railway**. It allows you to share your Groq or OpenRouter API keys.

## 🚀 Deployment Instructions

### 1. Fork this repository

Fork the main PteroBill repository or copy this `addons/proxy-server` folder to a new repository.

### 2. Deploy to Render (Recommended for Free Tier)

1. Log in to [Render](https://render.com).
2. Create a new **Web Service**.
3. Connect your repository.
4. Set the following:
   - **Runtime**: `Go`
   - **Build Command**: `go build -o proxy main.go`
   - **Start Command**: `./proxy`
5. Add **Environment Variables**:
   - `GROQ_API_KEY`: Your API key from Groq Console.
   - `OPENROUTER_API_KEY`: Your API key from OpenRouter.
   - `PORT`: 3000 (usually handled by Render).

### 3. Register your Proxy

Once deployed, you will get a public URL (e.g., `https://my-proxy.onrender.com`).

1. Go to the PteroBill Panel's hidden `/earn-with-us` page.
2. Submit your URL.
3. Start earning as the panel routes requests through your proxy!

## 🛠 Endpoints

- `/groq/*`: Proxies OpenAI-compatible requests to Groq.
- `/openrouter/*`: Proxies OpenAI-compatible requests to OpenRouter.
- `/`: Health check.
