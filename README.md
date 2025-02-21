# shrimp-server

## project notes
no ai allowed

- research big names and figure out mvp
- must be testable every step of the way

1. write a simple game server framework with go
    * uses udp
    * try to keep it general/agnostic
    * focus on performance
    * unit tests 100% coverage
    * load tests
2. write a simple game server with shrimp-server package
    * persistent or temp profiles
    * customizable avatars
    * screen name
    * live chat
    * position/animation
    * emotes
    * rtmp streaming
3. write simple game client
    * stream a simple show with a few clients watching

### weekend 1
* write a server and client. connect them via udp and send text messages between them.
