# Go Share

Simple file sharing service based on a randomly generated and easy to remember code for file download.

![Generated code](https://pp.vk.me/c638220/v638220677/1a2ee/TpPvEzfOILE.jpg)
![Main page](http://cs630430.vk.me/v630430677/8e/jY04CyWwBhM.jpg)

### How to run without docker
    - git clone https://github.com/IvanNyrkov/GoShare.git
    - sh run.sh
    - open browser on localhost:1337
 
### How to run with docker
    - git clone https://github.com/IvanNyrkov/GoShare.git
    - sh run-docker.sh
    - open browser on ${docker-host-ip}:80

### Backend
    - Go
    - Redis? (store info about saved files)

### Frontend
    - AngularJS
    - Bootstrap

### TODOs
    - Download file by generated code
    - Generated code uniqueness
    - Scheduled auto-deletion of old files from database        
    - Validation and error information on user-side
    - Info, Contacts pages
