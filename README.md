# Repository-Hasher

## Prerequisites
* Docker
* Postman - In order to use & test the Repository hashing service it is recommended to use [Postman](https://www.postman.com/) or any other kind of platform for sending API's.

### Clone the project:

```bash
git clone https://github.com/EllaRashty/Repository-Hasher.git
```

### Build and run the project with Docker:

Open a terminal window and navigate to the directory where you cloned the project to.

_Run the following command:_

```bash
docker-compose -f [example/downloads/Repository-Hasher/docker-compose.yml] up
```
## Microservices:
1. **Hasher** runs on port 9092
2. **Repository** runs on port 9091


# Usage

### 1. Test function  _GetFileContents(pathInRepo)_ 
Send a **POST** request with the URL:
```bash
http://localhost:9091/input-path
```
The body of the request where you can pass the argument ***[pathInRepo]***
```json

    "google/go-github"

```
<img width="847" alt="image" src="https://user-images.githubusercontent.com/48799296/191605458-7bc34500-278a-456f-ae1d-1e50fad4716f.png">

The result should look like this:

<img width="847" alt="image" src="https://user-images.githubusercontent.com/48799296/191605696-91c645b7-1702-464c-a03b-cd50ba0b58b0.png">


- - -

### 2. Test function  _CheckoutRef(gitRef)_ 
Send a **POST** request with the URL:

```bash
http://localhost:9091/checkout
```
The body of the request where you can pass the argument ***[gitRef]***
```json

    "release-cadence-readme"
    
```
<img width="847" alt="image" src="https://user-images.githubusercontent.com/48799296/191606292-005c2372-cdcb-4b8e-abdf-a65fed96d876.png">

The result should look like this:

<img width="847" alt="image" src="https://user-images.githubusercontent.com/48799296/191606856-25a263f4-9e6e-43f6-bf52-14c3fadc9f8a.png">

- - -

### 3. Test function  _HashFiles(pathInRepo, ...)_ 
**POST** request with the URL: 
```bash
http://localhost:9091/hash-path
```
The body of the request here you can pass the argument ***[pathInRepo]***
```json
"google/go-github"
```
<img width="829" alt="image" src="https://user-images.githubusercontent.com/113849701/191591828-48206663-1c38-4893-8ddd-fbfbf7dd839e.png"> 

 **Now** create another **GET** request with the URL:
 
 (pay attention that the port is 909**2**)
 
```bash
http://localhost:9092/hashing-service
```
<img width="849" alt="image" src="https://user-images.githubusercontent.com/113849701/191592027-30d2c1fb-4491-405c-8ba4-b8dda8a09ec4.png">

The result should look like this:

<img width="849" alt="image" src="https://user-images.githubusercontent.com/48799296/191608128-1f3a66da-e058-44d9-82ec-721d4f25e969.png">
