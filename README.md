# Go Share

Simple file sharing service based on randomly generated and easy to remember code for file download.

![Main page](http://cs630430.vk.me/v630430677/8e/jY04CyWwBhM.jpg)

![Generated code](http://cs630430.vk.me/v630430677/95/SUyOcnvlkJo.jpg)

### How to run without docker
    - git clone https://github.com/IvanNyrkov/GoShare.git
    - sh run.sh
    - open browser on localhost:1337
 
 ## How to run with docker
    - git clone https://github.com/IvanNyrkov/GoShare.git
    - sh run-docker.sh
    - open browser on ${docker-host-ip}:80

### Backend
    - Go
    - Echo
    - Redis? (store info about saved files)

### Frontend
    - AngularJS
    - Bootstrap

### TODOs
    - Download by passphrase page
    - Passphrase uniqueness
    - Scheduled auto-deletion of old files from database    
    - Info, Contacts pages
    - Validation and error information on user-side
