# gframe

In [gframe](https://gframe.fun), you can play a classic game --Rock-Paper-Scissors-- with others.

## Motivation
This repository is for practicing microservices architecture. It contains two web clients - Player Client & Admin Client, and several backend services as well. 
I will use as many technologies as possible in different projects. It includes Restful Api, GraphQL, Websocket, OAuth2, Redis, Kafka, Docker, Kubernetes...etc. For deployment, I've setup Azure Kubernetes Service (AKS) with 2 nodes for everything. Refer to [depoly](https://github.com/vincent-scw/gframe/tree/master/deploy) folder for details.
- Note: it is still in beginning stage.

## Architecture
![Architecture](https://github.com/vincent-scw/gframe/blob/master/gframe.png)

## How to Start
- PowerShell: `.\services\run.ps1` & `.\clients\run.ps1`
- Docker Compose: 
  - Backend: `docker-compose -f .\services\docker-compose.yml up`
  - Frontend: `docker-compose -f .\clients\docker-compose.yml up`
  - Open `Player Client` at http://localhost:8080
  - Open `Admin Client` at http://localhost:8081

## Behaviors
- Player Client (https://www.gframe.fun)
  - User must be able to access the Player Client, and join the game with simply providing a name
  - When player presses Start, he/she must be added to a gaming group
  - Player is able to play game with opponents
- Admin Client (https://admin.gframe.fun)
  - Admin must be able to access the Admin Client with username/password
  - Admin must be able to control the game (including creation, setting, etc...)
  - Admin must be able to simulate game playing via Simulator
  
## Key Features
- [ ] [Auth Service](https://github.com/vincent-scw/gframe/tree/master/services/oauth)
  - [x] Authorization & Authentication
  - [ ] OpenID
- [ ] Player Events
  - [x] Join
  - [ ] Leave
- [x] Players Matching
- [ ] Gaming
- [ ] Admin Client & Simulator
  - [ ] Admin Client
  - [ ] Simulator
    - [x] Inject Players
    - [ ] Game Play
- [ ] Log Monitoring
- [x] Docker Supporting
  - [x] Docker Compose
  - [x] Kubernetes
- [ ] Devops
  - [x] CI/CD
  - [ ] Microservices Monitoring
