# Exo-planets Golang microservice

Note: Open this file in an markdown supported IDE for appearance

#### Golang microservice for supporting the space voyagers who are embarking on a journey to study different exoplanets. Exo-planets are the planets outside of our solar systems

## Implemented Tasks
- Add an Exoplanet
- List Exoplanets
- Get Exoplanet by ID
- Update Exoplanet
- Delete Exoplanet
- Fuel estimation

## Extras Implemented Tasks
- Added swagger UI for better understanding and testing
- Implement validation for explanets data, ensuring that fields like planet's name, description, distance, radius and mass are provided correctly.
- Add unit tests to ensure the reliability of the service.
- Implement sorting based on radius and mass & filtering options for listing planets if time permits. (based on mass and name)

## How to run
- Run using go command
  - install go 1.21 or later
  - Go to source directory(/exoplanets)
  - Run go mod tidy
  - Run `go run server.go`
  - Open browser and go to http://localhost:1323. It will open swagger UI. Then play with the endpoints.
- Run using docker
  - Install docker in your system
  - Go to source directory(/exoplanets)
  - Run `docker build -t exoplanets:1.0.0 .`
  - Run `docker run -d -p=1323:1323 exoplanets:1.0.0 .`
  - `-d` is used to run the container in detached mode
   `-p=1323:1323` will forward port to host port 1323
  - Open browser and go to http://localhost:1323. It will open swagger UI. Then play with the endpoints.
  - Once the testing is done, stop the container. `docker ps` and find the container ID and stop `docker stop containerID`


  ### Announcement: 
  I would do so many things here. Due to time constraint, I have done these much. Hoping you like it 🙂.
