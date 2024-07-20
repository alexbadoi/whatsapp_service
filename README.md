
# WhatsApp Service

## Contact

For any questions, feel free to reach out on Discord: ![Discord](https://img.shields.io/badge/Discord-ak308465-7289DA)

## Objective

The main task is to implement a WhatsApp API service using Docker and integrate the WhatsApp Web API into the LizAnt UI React template provided.

## Deliverables

A Docker-compose configuration to deploy 4 services including:
- React Frontend supporting the React template
- Go main-backend
- PostgreSQL database
- WuzAPI chat service

## Getting Started

WuzAPI is an implementation of [@tulir/whatsmeow](https://github.com/tulir/whatsmeow) library as a simple RESTful API service with multiple device support and concurrent sessions. Clone the WuzAPI repository:
```bash
cd whatsapp_service
git pull https://github.com/asternic/wuzapi.git
```

Update the Docker file with:
```dockerfile
FROM golang:1.21-alpine AS build
RUN apk add --no-cache gcc musl-dev sqlite
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy
ENV CGO_ENABLED=1
RUN go build -o server .

FROM alpine:latest
RUN apk add --no-cache sqlite
RUN mkdir /app
COPY ./static /app/static
COPY --from=build /app/server /app/
VOLUME [ "/app/dbdata", "/app/files" ]
WORKDIR /app
ENV WUZAPI_ADMIN_TOKEN=YourSecureAdminTokenForAdminTasks
CMD [ "/app/server", "-logtype", "json" ]
```

Create a `docker-compose.yml`:
```yaml
version: '3.8'

services:
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    networks:
      - app-network

  backend:
    build: ./backend
    ports:
      - "8080:8080"
    networks:
      - app-network
    depends_on:
      - database
      - chat-service

  chat-service:
    image: asternic/wuzapi:latest
    ports:
      - "8081:8080"
    environment:
      - WUZAPI_ADMIN_TOKEN=YourSecureAdminToken
    networks:
      - app-network

  image: postgres:latest
  environment:
    POSTGRES_USER: user
    POSTGRES_PASSWORD: password
    POSTGRES_DB: chat_db
  volumes:
    - db_data:/var/lib/postgresql/data
  networks:
    - app-network
  ports:
    - "5432:5432"

networks:
  app-network:
    driver: bridge

volumes:
  db_data:
```

## Architecture Description

The application architecture is designed to be modular and scalable, leveraging microservices and containerization. Here is a detailed description of each component:

### Proposed Architecture Components

1. **Frontend (React)**:
    - Handles user interactions and displays chat messages.
    - Communicates with the backend services via REST API or WebSockets.

2. **Backend (Go)**:
    - Manages user sessions, authentication, and other core functionalities.
    - Handles communication with the chat service.
  
3. **Chat Service (WuzAPI)**:
    - Manages chat functionality, including message storage, retrieval, and real-time updates.
    - Integrates with the existing Go backend for user authentication.

4. **Database (PostgreSQL or MySQL)**:
    - Stores user data, chat messages, and other related information.
  
5. **Controller**:
    - Manages and routes requests to the appropriate services.
    - https://hub.docker.com/_/traefik

6. **Docker Compose**:
    - Manages multi-container Docker applications.

### Detailed Flow

1. **React Frontend**: Sends a chat message request to the Go main-backend.
2. **Go Backend**: Processes the chat request and, if needed, forwards it to the WhatsApp API.
3. **WhatsApp API**: Handles WhatsApp-specific tasks (e.g., sending messages) and returns the response to the Go backend.
4. **Go Backend**: Processes the WhatsApp API response and sends the final response back to the frontend.

## Scenario

There is an organization with multiple users using the React web-app.

1. **Register WhatsApp web** - The user must log in and receive a QR code to register WhatsApp web.
2. **De-Register WhatsApp web** - The user must be able to de-register WhatsApp web.
3. **Cat Load and Contact Merge** - WhatsApp web API is called to retrieve all contacts and historical chats. There is a pre-existing contacts table available in MySQL. The engine must join the data sets by phone numbers and identify matches.
4. **Multiple Users** - Multiple users use the same WhatsApp backend and are able to use the chat functionality offered by WuzAPI.
5. **Decoupled Backends** - We have the main-backend in Go which powers the React UI. The WhatsApp API service must include the WuzAPI package and be containerized independently of the main-backend. All requests from the frontend to the WuzAPI API should go through the Go main-backend.
6. **Postgres DB** - By default, WuzAPI is configured to use SQLite, but this implementation of [@tulir/whatsmeow](https://github.com/tulir/whatsmeow) library also supports PostgreSQL. A separate container running PostgreSQL must be deployed to support any CRUD operations from WuzAPI.

## MySQL Schema Credentials with Data Example

You can check some dummy data and data model using the credentials here:

- **MySQL Database:**
  - User: `engadev`
  - Password: `userengadev`
  - Host: `45.8.149.43:3306`
  - Schema: `data_feeds`
  - Use/insert credentials from the table: `user_ref` for login screen

## Notes on the Data

1. You can deploy a copy of the database locally by running `db/db_dump_file.sql`.
2. Once the cat app is connected via QR code, it should automatically pull in and populate the `contact_whatsapp` table with all available contacts from WhatsApp.
3. The table `tag_ref` includes available tags which should be fed into the contact card - see Enhance the contact requirements in the frontend section.
4. There is a `tag_id` field in `contact_ref` which should be displayed on the contact card in the chat app.
5. All chat data should be stored.
6. There are 3 available users in `user_ref` and they provide the login credential information for the web app.
7. Be able to add a contact from the chat app via the + button; this should populate `contact_ref`.

## Stack

- Golang + Fiber
- ReactJS

## Backend

- Use Golang with Fiber and GORM frameworks.

## Frontend

- Use the LizAnt Ant Design React Admin Dashboard template. This is already in a functioning state deployed with the Go backend. It currently only does some hotel search to Rapid API, the rest of the template is still inside.
- Modify the chat box for better screen fit.
- Add a favorites section to pin contacts.
- Enhance the contact card to integrate user-specific information and tags from the database.
- **Admin Screen:** Create a page where a user can request a QR code and register their WhatsApp. This page will handle the output from the WhatsApp service for QR code registration.
- **Chat App:** Develop a chat application for users to communicate. Ensure it can handle 3-4 concurrent users.

## Deployment

- Deploy using Docker Compose.
- Ensure the app handles 3-10 concurrent users with unique phone numbers for testing.
