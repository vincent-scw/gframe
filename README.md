# gframe
The goal of this project is to practice microservices architecture. It contains two web clients - player client & admin client. 
Players can join the online game from player client. In admin client, we can monitor what is happening.

## Architecture
![Architecture](https://github.com/vincent-scw/gframe/blob/master/gframe.png)

## How to Start
- PowerShell: `.\services\run.ps1` & `.\clients\run.ps1`
- Docker Compose: 
  - Backend: `docker-compose -f .\services\docker-compose.yml up`
  - Frontend: `docker-compose -f .\clients\docker-compose.yml up`

## Key Features
- [ ] [Auth Service](https://github.com/vincent-scw/gframe/tree/master/services/oauth)
  - [x] Authorization & Authentication
  - [ ] OpenID
- [ ] Player Events
  - [x] Join
  - [ ] Leave
- [ ] Players Matching
- [ ] Gaming
- [ ] Admin Client & Simulator
  - [ ] Admin Client
  - [ ] Simulator
    - [x] Inject Players
    - [ ] Game Play
- [ ] Log Monitoring
- [ ] Docker Supporting
  - [x] Docker Compose
  - [ ] Kubernetes
- [ ] Devops
  - [ ] CI/CD
  - [ ] Deployment
  - [ ] Microservices Monitoring
