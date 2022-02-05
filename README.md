### OpenStreaming

# About
This is a WIP project. Do Not Use.

Some things to note, the database is no longer wiped at startup however, the structure could still change without migrations as we are not in production.

# How to run
Requriments:
- domain with valid SLL certs located in resources/certs. named fullchain1.pem and privkey1.pem
- Docker && Docker Compose installed
- A registered twitch [application](https://dev.twitch.tv/console/apps) with OAuth Redirect URLs (Currently only support /twitch) 

Setup:
- copy example.env to .env and modify accordingly
- run `docker-compose build && docker-compose up` to build and run the application

