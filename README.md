# Exoplanet Microservice API

The Exoplanet Microservice API is a Go-based microservice designed to support space voyagers studying different exoplanets. It provides functionalities to manage various exoplanets, including adding, listing, updating, and deleting them. Additionally, it offers fuel estimation for trips to specific exoplanets based on crew capacity.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
    - [Adding an Exoplanet](#adding-an-exoplanet)
    - [Listing Exoplanets](#listing-exoplanets)
    - [Getting an Exoplanet by ID](#getting-an-exoplanet-by-id)
    - [Updating an Exoplanet](#updating-an-exoplanet)
    - [Deleting an Exoplanet](#deleting-an-exoplanet)
    - [Fuel Estimation](#fuel-estimation)
3. [Docker](#docker)
4. [Testing](#testing)
5. [Contributing](#contributing)
6. [License](#license)

## Installation

To install the Exoplanet Microservice API, you need to have Go installed on your system. Clone the repository and build the project using the following commands:

```bash
git clone https://github.com/your-username/exoplanet-microservice.git
cd exoplanet-microservice
go build
```

## Usage
### Adding an Exoplanet

To add a new exoplanet, send a POST request to /exoplanets with the required properties:

```json
{
    "name": "Exoplanet Name",
    "description": "Description of the exoplanet",
    "distance": 100, // Distance from Earth (light years)
    "radius": 1,    // Radius of the exoplanet (Earth-radius units)
    "mass": 1,      // Mass of the exoplanet (Earth-mass units, required only for terrestrial type)
    "type": "GasGiant" // Type of exoplanet (GasGiant or Terrestrial)
}
```

### Listing Exoplanets

To retrieve a list of all available exoplanets, send a GET request to /exoplanets. You can also use query parameters for sorting and filtering the results.

### Getting an Exoplanet by ID

To retrieve information about a specific exoplanet by its unique ID, send a GET request to /exoplanets/{id}.

### Updating an Exoplanet

To update the details of an existing exoplanet, send a PUT request to /exoplanets/{id} with the updated properties.

### Deleting an Exoplanet

To remove an exoplanet from the catalog, send a DELETE request to /exoplanets/{id}.

### Fuel Estimation

To retrieve an overall fuel cost estimation for a trip to any particular exoplanet for a given crew capacity, send a GET request to /exoplanets/{id}/fuel with the query parameter crewCapacity.

### Docker

You can also run the Exoplanet Microservice API using Docker. Build the Docker image and run the container using the following commands:

```bash
docker build -t exoplanet-microservice .
docker run -p 8080:8080 exoplanet-microservice
```
### Testing

Run the tests using the `go test` command.