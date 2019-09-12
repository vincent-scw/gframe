[![Build Status](https://travis-ci.org/vincent-scw/gframe.svg?branch=master)](https://travis-ci.org/vincent-scw/gframe)

# gframe

**Let's play a classic game --Rock-Paper-Scissors-- with others in [gframe](https://gframe.fun).**

## Motivation
This is a [proof-of-concept application](https://gframe.fun), which demostrates [Microservice Architecture Pattern](http://martinfowler.com/microservices/) using various technologies.
For technical perspective, it includes Restful Api, GraphQL, Websocket, OAuth2, Redis, Kafka, Docker, Kubernetes...etc. For deployment, I've setup Azure Kubernetes Service (AKS) for everything. Refer to [depoly](https://github.com/vincent-scw/gframe/tree/master/deploy) folder for details.
- Note: it is still in beginning stage.

## Functional services
<img with="880" alt="Functional services" src="https://github.com/vincent-scw/gframe/blob/master/gframe-functional.png" />

### Game service
Handle player login/logout, and interact with game playing.

### Broker service
Group players.

### Admin service
Setup, control and monitor a game. 

## Infrastructure services

### Redis
This project heavily depends on Redis. It uses several features in Redis, like Cache, Pub/Sub, Sets...etc

### Kafka
This project use Kafka as message queue.

### Monitoring services

#### Promethuse & Grafana (https://grafana.gframe.fun)

#### Jaeger

## Clients
- Player Client (https://www.gframe.fun)
  - User must be able to access the Player Client, and join the game with simply providing a name
  - When player presses Start, he/she must be added to a gaming group
  - Player is able to play game with opponents
- Admin Client (https://admin.gframe.fun)
  - Admin must be able to access the Admin Client with username/password
  - Admin must be able to control the game (including creation, setting, etc...)
  - Admin must be able to simulate game playing via Simulator

## How to Start
- PowerShell: `.\services\run.ps1` & `.\clients\run.ps1`
- Docker Compose: 
  - Backend: `docker-compose -f .\services\docker-compose.yml up`
  - Frontend: `docker-compose -f .\clients\docker-compose.yml up`
  - Open `Player Client` at http://localhost:8080
  - Open `Admin Client` at http://localhost:8081
  
## Key Features
- [ ] [Auth Service](https://github.com/vincent-scw/gframe/tree/master/services/oauth)
  - [x] Authorization & Authentication
  - [ ] OpenID
- [ ] Player Events
  - [x] Join
  - [ ] Leave
- [x] Players Matching
- [x] Gaming
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
