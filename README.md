# gosecAPI

Welcome to the gosecAPI documentation! This API serves as the backend for the gosec password manager, providing various endpoints to manage users, configurations, passwords, and authentication.

## Endpoints

### Users

#### Get All Users

- **Endpoint:** `/users`
- **Method:** GET
- **Description:** Retrieve a list of all users registered in gosec.
- **Response:** A JSON array containing user objects.

#### Get User by ID

- **Endpoint:** `/users/:id`
- **Method:** GET
- **Description:** Retrieve a specific user by their ID.
- **Parameters:** `id` (user's ID)
- **Response:** A JSON object representing the user.

### Config

#### Get All Configs

- **Endpoint:** `/config`
- **Method:** GET
- **Description:** Retrieve all configuration settings in gosec.
- **Response:** A JSON array containing config objects.

#### Get Config by User ID

- **Endpoint:** `/config/:id`
- **Method:** GET
- **Description:** Retrieve the configuration file associated with a specific user.
- **Parameters:** `id` (user's ID)
- **Response:** A JSON object representing the user's configuration.

### Password

#### Get Passwords by User ID

- **Endpoint:** `/passwords/:id`
- **Method:** GET
- **Description:** Retrieve all passwords associated with a specific user.
- **Parameters:** `id` (user's ID)
- **Response:** A JSON array containing password objects.

### Authentication

#### User Authentication

- **Endpoint:** `/auth`
- **Method:** POST
- **Description:** Authenticate a user by providing their username and password.
- **Request Parameters:** `username`, `password`
- **Response:** If the provided credentials are correct, a JSON object representing the user will be returned. If the credentials are incorrect, an error message will be returned.

## Usage

You can use the gosecAPI to manage users, configurations, passwords, and authentication for your password manager application. Here's a brief guide on how to use the API:

1. Use the `/auth` endpoint to authenticate users by sending a POST request with their `username` and `password`.

2. After authentication, you can use the retrieved user object's ID to make requests to other endpoints to manage their configurations, passwords, and more.

3. Use the `/config` and `/passwords` endpoints with the user's ID to retrieve their configuration settings and passwords, respectively.

4. If you need to retrieve a specific user, configuration, or password, you can use the endpoints that support fetching by ID.

Remember to include the appropriate HTTP methods, request parameters, and handle the JSON responses according to your application's needs.

Feel free to explore the various endpoints of gosecAPI and integrate them into your GitHub project's backend to enhance the functionality of your gosec password manager.

For more detailed information on each endpoint, request/response examples, and error handling, refer to the respective sections above.
