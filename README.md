# go-api
Creating apis using golang  

To initialise the modules: go mod init example/go-api  
This creates the go.mod file  
To add a dependency in golang(ex):  go get github.com/gin-gonic/gin  

body.json contains the new data we want to create.  
The curl command for the post req:  
curl localhost:8008/books --include --header "Content-Type: application/json" -d @body.json --request "POST"  
-d => data  
@  => file  and the specify the request type  


gin-gonic is used for the api.  
No explicit db is used (yet)  
Required dependencies -  
net/http  
github.com/gin-gonic/gin  

gin-gonic is used for the api creation and handling.  


A little off topic,  
The gin framework features a Martini-like API, but with performance up to 40 times faster than Martini. (The Martini framework is no longer maintained)  
Their tagline being "If you need smashing performance, get yourself some Gin."  
Well, they knew what they were doing. Plus the pun intended behind gin-gonic doesn't need further explanation,duh  
