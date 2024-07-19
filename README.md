## Contact

For any questions, feel free to reach out on Discord: ![Discord](https://img.shields.io/badge/Discord-ak308465-7289DA)

# Objective:
The main ask is to implement a whatsap api service using docker
Integrate the WhatsApp Web API into the LizAnt UI react template given.

## Deliverable:
A docker-compose configuration to deploy 4 containers includeing:
- React Frontend supporting the React template
- Go main-backend
- Postgresql database
- Qazapi api

# To get started:

WuzAPI is an implementation of @tulir/whatsmeow library as a simple RESTful API service with multiple device support and concurrent sessions.
Clone the Wuzapi repository:
```
cd whatsapp_service
git pull https://github.com/asternic/wuzapi.git
```

update the Docker file with:

```
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
ENV WUZAPI_ADMIN_TOKEN SetToRandomAndSecureTokenForAdminTasks
CMD [ "/app/server", "-logtype", "json" ]

```

create docker-compose.yml
```
version: '3.8'

services:
  wuzapi:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./dbdata:/app/dbdata
      - ./files:/app/files
    environment:
      - WUZAPI_ADMIN_TOKEN=YourSecureAdminToken

```
check users inside the container: 
```
docker exec -it docker_container_id  sh
sqlite3 dbdata/users.db "select * from users"
```


add user inside the container sql:
```
docker exec -it docker_container_id  sh
cd /app/dbdata
sqlite3 users.db "insert into users (name, token) values ('John', '1234ABCD')"
```

# Scenario:

There is an organisation with multiple users using the react web-app.

1. **Register WhatsApp web** - The user must log and receive QR code to register whats-app web
2. **De-Register WhatsApp web** - The user must be able to de-register whats-app web
3. **Cat Load and Contact Merge** - Whatsapp web api i called to retreive all contacts and hostorical chats. there is a pre-existing contacts table available in mysql. the engine must join the data sets by phone numbers and identify matches.
4. **Multiple Users** - Multiple users use the same whatsapp back-end and are able to use the chat functionality offered by wazapi
5. **Decoupled backends** - We have the main-backend in go which powers the react UI. the whatsapp API service must include the wazapi package and containerised independently of the main-backend which is present here in this repo. All requests from the frontend to the wazapi API should go through the Go main-backend. Here’s the flow:
- 
  - 
    - **React Frontend:** Sends a chat message i.e. request to the Go main-backend.
    - **Go Backend:** Processes the chat request and, if needed, forwards it to the WhatsApp API.
    - **WhatsApp API:** Handles WhatsApp-specific tasks (e.g., sending messages) and returns the response to the Go backend.
    - **Go Backend:** Processes the WhatsApp API response and sends the final response back to the frontend.
6. **Postgres DB** By default wazapi is configured to use sqlite but this is an implementation of @tulir/whatsmeow library which also suports PostgreSQL. A separate container running PostgreSQL must be deployed to dupport any CRUD operations from wazapi.


## mysql schema credentials with data example:

You can check some dummy data and data model using the credentials here:

- **MySQL Database:**
  - User: `engadev`
  - Password: `userengadev`
  - Host: `45.8.149.43:3306`
  - Schema: `data_feeds`
  - Use/insert credentials from the table: `user_ref` for login screen


## Noted on the data:
1. you can deploy a copy of the db loclaly by running db/db_dump_file.sql
2. once the cat app is connected via qr code it should automaticlaly pull in and populate contact_whatsapp table with all available contacts from whstapp
3. table tag_ref includes available tags which should be fed into the contact card - see Enhance the contact requirements in the frontend section.
4. there is a tag_id field in contact ref this should be displayed on teh contact card in the chat app. 
5. all chat data should be stored
6. there are 3 available users in user_ref and they provide the login credential information for the web app
7. be able to add a contact from chat app via + button this should populate contact_ref

## Stack:
- Golang + Fiber
- ReactJS

 
## Backend:
- Use Golang with Fiber and GORM frameworks.

## Frontend:
- Use the Lizant Ant Design React Admin Dashboard template. this is already in a functioning state deployed with go backend. it currently only does some hotel search to rapid API the rest of the template still inside.
- Modify the chat box for better screen fit.
- Add a favorites section to pin contacts.
- Enhance the contact card to integrate user-specific information and tags from the database.
- **Admin Screen:** Create a page where a user can request a QR code and register their WhatsApp. This page will handle the output from the WhatsApp service for QR code registration.- **Chat App:** Develop a chat application for users to communicate. Ensure it can handle 3-4 concurrent users.

## Deployment:
- Deploy using docker-compose.
- Ensure the app handles 3-10 concurrent users with unique phone numbers for testing.
