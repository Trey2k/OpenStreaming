### OpenStreaming

# About
This is a WIP project. To not use.

Some things to note, since this project is in active development and not in use. The database is wiped every start up. This will be changed later on down the line. But it has made things more convenient for me currently.

# How to run
Requriments:
- domain with valid SLL certs located in resources/certs. named fullchain1.pem and privkey1.pem
- Docker && Docker Compose installed
- A registered twitch [application](https://dev.twitch.tv/console/apps) with OAuth Redirect URLs (Currently only support /twitch) 

Setup:
- copy example.env to .env and modify accordingly
- run `docker-compose build && docker-compose up` to build and run the application

# TODO
- Replace all hardcoded URLS with os.Getenv("URL")
