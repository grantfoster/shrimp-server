# shrimp-server

practicing being serious coder by making a multiplayer game server framework

## prerequisites ‼️

go 1.24.0 

postgresql db with db named "shrimp_server"

## initial setup for local dev 💻

1. first things first blast off a ```go mod tidy``` to grab all deps
2. ```go run ./cmd/migrate/``` to migrate that "shrimp_server" db up
3. smash ```go run ./cmd/server/``` to start the bad boy

you can now use postman or curl to hit /ping and netcat to write messages to udp and see them in your server stdout.
```curl localhost:7777/ping```
```nc -u localhost 6666```
📝
## db & migrations 🦜

again, you'll need a psql running with a db name of "shrimp_server". we use this db for long lived data such as user profiles or session info.
we use golang-migrate to manage our db and then we use sqlc to generate beautiful go code from queries written in the queries folder. later on we will automat migration generation from with the admin console for custom dbs that the server manager can define via ui or code.

- to generate a new migration ```go run ./cmd/generate/ -n=your_migration_name_snake_case```
    - uses raw psql
- to migrate up or down (change d flag to change direction) ```go run ./cmd/migrate/ -d=up```
- to run sqlc codegen:
    - write any valid query in an sql file in the queries folder with the proper inline comment found in sqlc docs.
    - run ```sqlc generate``` from project root.
    - use cool code in project

## project notes 📝
~~no ai allowed~~
note: damnit that didn't last long. how about this:
no copied ai codegen

- research big names and figure out mvp
    - nakama is the big one and it's written in golang so i can get inspiration
- must be testable every step of the way
    - coming soon lol

### general goals ⚽️
1. write a multiplayer game server framework with go called "shrimp-server"
    * uses udp/grpc/websockets do we even want REST?
    * try to keep it general/agnostic for multiple game engines
    * focus on performance, configurability, and simplicity
    * unit tests 100% coverage lol
    * load tests as soon as possible
2. write a simple game server with shrimp-server package
    * persistent or temp profiles
    * customizable avatars
    * screen name
    * live chat
    * position/animation
    * emotes
    * rtmp streaming
3. write simple 2d chatroom like game client
    * customizable avatars that can walk around a simple pixel world w a few animations
    * one or more multiplayer minigames - GAMBLING ASAP
    * stream a simple show over rtmp/obs to an ingame "cinema" with a few clients watching

### weekend 1 - done ✅
* write a server and use netcat and postman to test. connect them via udp/tcp and send text messages between them.

### weekend 2 - done ✅
* setup migrations for psql
* setup sqlc for db interactions
* set up script for generating new sequential migrations 
    * (all this does is make a valid) prefixed file to create an sql migration in

### weekend 3 - todo
* postgresql_lsp documented and using env vars perhaps?
* change from unix socket in db to userpass
* make into a package rather than just a go program
* write dummy simple udp server for text chat with no auth
* write dummy simple client for text chattin using no auth
* maybe dockerize idk
* maybe get redis going idk
