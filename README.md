# gothqr
The project purpose for study go programming and  application development. That coming by reuse business function from my ThaiQR PromptPay project with developed by Node.js. This project as a part of Esppree.io startup company. I build that project on MVP phase.  My intention on this job as the hobbies and share knowledge and experience to very one. 

# Hobbies Roadmap
  I will on going write the code when i have free time from normal wokring. I plan as this.
  1. Construct project and connect mongodb and write some data write and read function ( this git intial version )
  2. Enanble Web API , Echo framework for enable business function can access by Web API
  3. Complete all function for prove go coding and go language technique with old requirement from Node.js

# Working log
  Branch: master
    Initial project structure and promote mongodb connection manager and insert , read data function
  Branch: ep1
    Add Eco Http framework for start Web API development and 2 endpoint example. And added system shuting down handling.
  Branch: ep2
    Clean code by separate routing api function to other package and go file. And example endpoint return array of json object
  Branch: ep3
    Add example real endpoint for you take to use in your job, That include GET method with Query string , PathData and POST methd with extract json input data to MAP. And use string package to manipulate string data.
    
# Prerequisites
- MongoDB
- latest Go
- Go library
  go get -u github.com/labstack/echo/...
  go get gopkg.in/mgo.v2
  go get github.com/fatih/structs
  
- Instruction:
  The configruation file sit in ./config/config.json , By this file for mainly configure parameter such as db server uri , db name and other attribute as you require on the next time.
  
 # Running the tests
 go run server.go , the operation will added mongod doc. to people collection and list all document and print out on screen.
 
 # Hop you enjoy with me for new programming world with Golang.
