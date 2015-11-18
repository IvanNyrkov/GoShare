# Go Share

Simple file sharing service based on randomly generated and easy to remember code for file download.

[Main page](http://cs630430.vk.me/v630430677/8e/jY04CyWwBhM.jpg)

[Generated code](http://cs630430.vk.me/v630430677/95/SUyOcnvlkJo.jpg)

### How to run
    - git clone https://github.com/IvanNyrkov/Go-Share.git
    - go run server/*.go
    - open browser on localhost:3010

### Settings
To see all possible options run:
    - go run server/*.go -help  

### Backend
    - Golang
    - Negroni
    - Gorilla

### Frontend
    - AngularJS
    - Bootstrap

### Todos
    - Passphrase uniqueness
    - Scheduled auto-deletion of old files from database
    - Download by passphrase page
    - Info, Contacts pages
    - Validation and error information on user-side
