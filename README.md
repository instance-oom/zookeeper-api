# zookeeper-api
zookeeper rest api by golang

# Usage

Clone And Install Dependencies
```cmd
git clone https://github.com/lon-yang/zookeeper-api.git
cd ./zookeeper-api/src/zookeeper-api
glide install
```

Build
```cmd
go build zookeeper-api
```

Run
```cmd
zookeeper-api.exe
```

# API

<b>Get Childs</b>
>http://localhost:8000/za/v1/childs/{path}
  
Example   
```cmd
CURL http://localhost:8000/za/v1/childs/zookeeper
```
```json
{
  "Childs": [
    "quota"
  ],
  "Path": "/zookeeper"
}
```

<b>Get Node</b>  
>http://localhost:8000/za/v1/node/{path}

Example
```cmd
CURL http://localhost:8000/za/v1/node/zookeeper
```
```json
{
  "ChildNum": 1,
  "InDate": 0,
  "LastEditDate": 0,
  "Path": "/zookeeper",
  "Value": "",
  "Version": 0
}
```

<b>Create Node</b>
>`POST` - http://localhost:8000/za/v1/node/{path}

Example
```json
//URL: http://localhost:8000/za/v1/node/test
//Request Body
{
    "value": "test"
}

//Response
{
  "Path": "/test"
}
```

<b>Update Node</b>
>`PUT` - http://localhost:8000/za/v1/node/{path}

Example
```json
//URL: http://localhost:8000/za/v1/node/test
//Request Body
{
    "value": "test2"
}

//Empty Response
```

<b>Delete Node</b>
>`DELETE` - http://localhost:8000/za/v1/node/{path}

Example
```json
//URL: http://localhost:8000/za/v1/node/test
//Request Body
{
    "value": "test2"
}

//Empty Response
```

<b>Server State</b>
>`GET` - http://localhost:8000/za/v1/stat/{ip:port}

Example
```cmd
CURL http://localhost:8000/za/v1/stat/127.0.0.1:2181,127.0.0.1:2182,127.0.0.1:2183
```
```json
[                                
  {                              
    "Server": "127.0.0.1:2181",
    "NodeCount": 280,            
    "MinLatency": 0,             
    "AvgLatency": 0,             
    "MaxLatency": 103,          
    "Connections": 35,           
    "Mode": "follower",          
    "Version": "3.4.5-1392090",  
    "Error": null                
  },
  {                              
    "Server": "127.0.0.1:2182",
    "NodeCount": 280,            
    "MinLatency": 0,             
    "AvgLatency": 0,             
    "MaxLatency": 158,          
    "Connections": 10,           
    "Mode": "leader",          
    "Version": "3.4.5-1392090",  
    "Error": null                
  },
  {                              
    "Server": "127.0.0.1:2183",
    "NodeCount": 280,            
    "MinLatency": 0,             
    "AvgLatency": 0,             
    "MaxLatency": 514,          
    "Connections": 2,           
    "Mode": "follower",          
    "Version": "3.4.5-1392090",  
    "Error": null                
  }                                  
]                                
```