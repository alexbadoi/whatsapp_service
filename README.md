# Objective:
the main ask is to implement a whatsap api service using kubernetes
Integrate the WhatsApp Web API into a custom CRM (Safari Expert) for chat functionality and lead/sales management.

## Data Structure:
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
- 
## Backend:
- Use Golang with Fiber and GORM frameworks.
- Implement endpoints for user login, QR code registration, and full chat functionality.
- Connect to an external MySQL database.

## Frontend:
- Use the Lizant Ant Design React Admin Dashboard template.
- Modify the chat box for better screen fit.
- Add a favorites section to pin contacts.
- Enhance the contact card to integrate user-specific information and tags from the database.
- **Admin Screen:** Create a page where a user can request a QR code and register their WhatsApp. This page will handle the output from the WhatsApp service for QR code registration.
- **Chat App:** Develop a chat application for users to communicate. Ensure it can handle 3-4 concurrent users.


## Deployment:
- Deploy using Kubernetes.
- Ensure the app handles 3-10 concurrent users with unique phone numbers for testing.

## Architecture

### Frontend:
- Sends requests to the Go backend.

### Go Backend:
- Processes requests.
- Forwards relevant requests to the WhatsApp API.
- Returns responses to the frontend.

### WhatsApp API:
- https://github.com/chrishubert/whatsapp-api.git
- Manages WhatsApp-specific tasks.
- Communicates with the Go backend.

### Kubernetes Configuration

#### Deployments:
- backend-deployment.yaml
- frontend-deployment.yaml
- whatsapp-api-deployment.yaml

#### Services:
- services.yaml

### Deployment Steps
1. Create Deployment and Service YAML files.
2. Apply configurations using `kubectl apply -f <file.yaml>`.
3. Access services via NodePort.

### Suggested Architecture
All requests from the frontend to the WhatsApp API will go through the Go backend. Here’s the flow:

#### Flow of Requests
1. **Frontend:** Sends a request to the Go backend.
2. **Go Backend:** Processes the request and, if needed, forwards it to the WhatsApp API.
3. **WhatsApp API:** Handles WhatsApp-specific tasks (e.g., sending messages) and returns the response to the Go backend.
4. **Go Backend:** Processes the WhatsApp API response and sends the final response back to the frontend.

### Kubernetes Configuration

#### services.yaml
```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
  type: NodePort

---

apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  selector:
    app: frontend
  ports:
    - protocol: TCP
      port: 3000
      targetPort: 3000
  type: NodePort

---

apiVersion: v1
kind: Service
metadata:
  name: whatsapp-api
spec:
  selector:
    app: whatsapp-api
  ports:
    - protocol: TCP
      port: 3001
      targetPort: 3000
  type: NodePort
```

#### backend-deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: your-backend-image
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          value: "your_external_db_host"
        - name: DB_USER
          value: "root"
        - name: DB_PASSWORD
          value: "yourpassword"
        - name: DB_NAME
          value: "safari_expert"
```

#### frontend-deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
      - name: frontend
        image: your-frontend-image
        ports:
        - containerPort: 3000
```

#### whatsapp-api-deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: whatsapp-api
spec:
  replicas: 1
  selector:
    matchLabels:
      app: whatsapp-api
  template:
    metadata:
      labels:
        app: whatsapp-api
    spec:
      containers:
      - name: whatsapp-api
        image: your-whatsapp-api-image
        ports:
        - containerPort: 3000
        env:
        - name: BASE_WEBHOOK_URL
          value: http://backend:8080/webhook
        - name: ENABLE_SWAGGER_ENDPOINT
          value: "true"
        - name: API_KEY
          value: "your_api_key"
        volumeMounts:
        - name: session-volume
          mountPath: /app/session
      volumes:
      - name: session-volume
        hostPath:
          path: /path/to/your/session
```

### Summary
- **Frontend:** Sends requests to the backend via the NodePort service on port 3000.
- **Backend:** Handles requests, forwards to WhatsApp API if needed, and returns responses to the frontend. Exposed via NodePort on port 8080.
- **WhatsApp API:** Handles WhatsApp-specific tasks and communicates with the backend. Exposed via NodePort on port 3001.

This setup maintains clear communication paths and allows each service to operate independently while ensuring they can interact as needed.

Ensure you have `kubectl` configured to interact with your Kubernetes cluster. Then apply the configuration files:
```sh
kubectl apply -f backend-deployment.yaml
kubectl apply -f frontend-deployment.yaml
kubectl apply -f whatsapp-api-deployment.yaml
kubectl apply -f services.yaml
```

This setup allows you to deploy your applications to Kubernetes in a manner similar to using docker-compose, providing a straightforward way to manage your services and scale them as needed.

## Credentials:
- **MySQL Database:**
  - User: `engadev`
  - Password: `userengadev`
  - Host: `45.8.149.43:3306`
  - Schema: `data_feeds`
  - Use/insert credentials from the table: `user_ref` for login screen

## Contact
For any questions, feel free to reach out on Discord: ![Discord](https://img.shields.io/badge/Discord-ak308465-7289DA)
