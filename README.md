##go-rollercoaster-api

### What is it and what does it do?
A REST API built with go, that allows you to read and add rollercoaster data.

### How to use it?
1. Make sure you have Go 1.16 installed
2. Set the environment variable "ADMIN_PASSWORD=secret"
3. Run the main function in the 'server.go' file  

    ####Routes
   * /costers -> shows all rollercoaster data
   * /coasters/\<id> -> show rollercoaster data for specific id
   * /addCoaster -> add a rollercoaster
   * /admin -> login to see the super secret admin page

### What did I learn?
* Go
* Go http package

### Disclaimers
This project was created following kubucation's video on YouTube.  
You can find it here: https://www.youtube.com/watch?v=2v11Ym6Ct9Q&ab_channel=kubucation  

Functionality I added includes:
* Refactored code and created handlers folder for improved readability
* "/" redirects to "/coasters" and anything else gets "page not found" html as a response
* Add "WWW-Authenticate: Basic" header to the response in case of wrong login credentials for "/admin"
* Show form under "/addCoaster" to add a coaster from the browser



### Curl commands for testing the API
You can also use the browser to test the routes

#### Post Data:  
curl -v localhost:8080/coasters -X POST -d '{"name": "Taron", "inPark": "Phantasialand", "height": 30, "manufacturer": "Intamin"}' -H "Content-Type: application/json"  

#### Get Data:  
curl -v localhost:8080/coasters | jq  

#### Get Random Coaster (with redirect):
curl localhost:8080/coasters/random -L  

#### Test Authentication
curl localhost:8080/admin -u admin:secret

