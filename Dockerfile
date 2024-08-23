# Dockerfile

# Step 1: Build the frontend
FROM node:20 AS build-frontend
WORKDIR /app
COPY frontend/ .
RUN npm install && npm run build

# Step 2: Build the backend
FROM golang:1.19 AS build-backend
WORKDIR /go/app
COPY backend/ .
RUN go build -buildvcs=false

# Step 3: Final image
FROM ubuntu:latest
WORKDIR /app

# Copy backend binary
COPY --from=build-backend /go/app/personal-site .

# Copy frontend build
COPY --from=build-frontend /app/.svelte-kit/output ./frontend

# Ensure the `personal-site` file has execute permissions
RUN chmod +x personal-site

# Serve the frontend
CMD ["./personal-site"]
