# Repository-Hasher

## Prerequisites
* Git
* Postman - In order to use & test the Repository hashing service it is recommended to use [Postman](https://www.postman.com/) or any other kind of service for sending API's.

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
1. **POST** request 
2. Type this URL
```bash
http://localhost:9091/input-path
```
3. The body of the request where you can pass the argument ***[pathInRepo]***
```json
{
    "name": "google/go-github"
}
```
<img width="839" alt="image" src="https://user-images.githubusercontent.com/113849701/191590878-8468a4f9-f9e4-416f-b2b6-1ac2868ffa94.png"> 

4. Click **Send**
5. The result should look like this:

<img width="865" alt="image" src="https://user-images.githubusercontent.com/113849701/191591440-d91defad-3e4c-4261-9862-29569b8692f7.png">
- - -
### 2. Test function  _CheckoutRef(gitRef)_ 
1. **POST** request 
2. Type this URL
```bash
http://localhost:9091/checkout
```
3. The body of the request where you can pass the argument ***[gitRef]***
```json
{
    "name":"release-cadence-readme"
}
```
<img width="830" alt="image" src="https://user-images.githubusercontent.com/113849701/191591617-90d9051f-cfd6-4cce-b0d2-c205ce5fd5a6.png"> 
4. Click **Send**
5. The result should look like this:

<img width="847" alt="image" src="https://user-images.githubusercontent.com/113849701/191591673-d40c366f-e777-49d4-908b-239b5c01d0eb.png">
- - -
### 3. Test function  _HashFiles(pathInRepo, ...)_ 
1. **POST** request 
2. Type this URL
```bash
http://localhost:9091/hash-path
```
3. The body of the request here you can pass the argument ***[pathInRepo]***
```json
"google/go-github"
```
<img width="829" alt="image" src="https://user-images.githubusercontent.com/113849701/191591828-48206663-1c38-4893-8ddd-fbfbf7dd839e.png"> 
4. Click **Send**
5. Create another **GET** request
6. Type this URL
```bash
http://localhost:9092/hashing-service
```
7. -no body- 

<img width="849" alt="image" src="https://user-images.githubusercontent.com/113849701/191592027-30d2c1fb-4491-405c-8ba4-b8dda8a09ec4.png">
8. Click **Send**
9. The result should look like this:

picture of the result
