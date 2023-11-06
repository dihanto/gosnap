# GOSNAP

![gosnap-high-resolution-logo-color-on-transparent-background](https://github.com/dihanto/gosnap/assets/39905651/704fa710-4045-44a9-a066-a2ea6daa37f4)

GoSnap is a powerful API built using the Go programming language that enables users to seamlessly upload, like, and comment on photos. With GoSnap, users can easily share their precious moments and interact with each other's photo uploads through a robust and user-friendly interface.

## Features

- **Photo Upload**: Users can upload their photos to share their special moments with others.
- **Follow Users** : Users can follow each other.
- **Like and Comment**: Users can like and comment on the photos uploaded by other users, fostering engagement and interaction within the community.

## File Structure

Explanation of the file structure:

- `/`
  - `/api/`: API Specs/Documentation.
  - `/cmd/`: Command-line interface.
    - `main.go`: Main application entry point.
  - `/database/`: Migration Files.
  - `/internal/`: Internal application code.
    - `/config/`: Configuration files.
    - `/controller/`: Request handlers and route controllers.
    - `/exception/`: Error handling utilities.
    - `/helper/`: Helper functions and utilities.
    - `/middleware/`: Middleware functions.
    - `/repository/`: Data access layer.
    - `/usecase/`: Business logic and use cases.
  - `/model/`: Data models and structures.

## Getting Started

To get started with GoSnap, follow these steps:

1. Clone the repository: `git clone https://github.com/dihanto/gosnap.git`
2. Navigate to the project directory: `cd gosnap`
3. Install dependencies: `go mod tidy`
4. Configure the application by adding the files in the `/cmd/config.json` directory as per your requirements.
5. Run the application: `go run cmd/main.go`
6. The GoSnap API will be accessible at `http://localhost:8000`.

## API Documentation

For detailed API documentation and specifications, please refer to the `/api/` directory.



## License

GoSnap is open-source software licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute the code as per the terms of the license.

## Support

For any issues, bugs, or questions related to GoSnap, please [create an issue](https://github.com/dihanto/gosnap/issues) on the GitHub repository.

