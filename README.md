## microservices-example
This is a todo app based on microservice architecture. I made this project while learning about Docker, Kubernetes, gRPC, Service Meshes.

## Architecture
Here is an overview of the application architecture.

![Architecture Overview](./readme-assets/architecture.jpg)

| Service | Language, Framework | Description |
| --- | --- | --- |
| [frontend](./frontend/) | HTML, CSS, JavaScript | This is a simple web app written in plain HTML, CSS, JavaScript. It's docker image uses `apache httpd` to serve the web app. |
| [api-service](./api-service/) | Python, FastAPI | This service exposes RESTful APIs that the `frontend` communicates with. It doesn't have much business logic but instead further communicates with `auth-service` and `todo-service` over `gRPC` to handle the requests. The protobuf definitions can be found [here](./protobuf/). |
| [auth-service](./auth-service/) | Go | This service handles all the authentication, registration, etc. logic. It stores all its data in a `PostgreSQL` database. |
| [todo-service](./todo-service/) | TypeScript | This service basically handles all the CRUD operations of todos. It stores all its data in a `MongoDB` database. |

## Screenshots
### Todo App

| Home Page | Login Page |
| --- | --- |
| ![Home Page](./readme-assets/app-main.png) | ![Login Page](./readme-assets/app-login.png) |

### Services Graph by Kiali in Istio Service Mesh
![Kiali Graph](./readme-assets/kiali-graph.png)