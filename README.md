# Reservation Service Back-End

## Table of Contents
- [Table of Contents](#table-of-contents)
- [About](#about)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running](#running)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## About
This project is a Reservation Service Back-End for the Reservation Service Application.

## Getting Started
### Prerequisites
- Go 1.18+
- Docker
- Docker Compose
- Make
- Delve (Debugger)
- Air (Live Reloader)

### Installation
1. Clone the repository
```bash
git clone https://github.com/tapiaw38/reservation-service-be.git
```
2. Install dependencies
```bash
make install-deps
```
3. Run the project
```bash
make run
```

### Running
1. Run the project
```bash
make run
```
2. Run the project on watch mode for development purposes
```bash
make run-dev
```

## Usage
The project is a RESTful API that provides the following functionalities:

### Health Check
The API provides a health check endpoint that returns a 200 status code if the service is running.

#### Request
```bash
GET /api/v1/ping
```

#### Response
```json
{
  "message": "pong"
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/) - Copyright (c) 2024 Tapia Walter <tapiawalter@gmail.com> (https://github.com/tapiaw38) 
