# Go SQL Seed

This app provides an easy way to execute migrations and populate database with additional data, which is particularly useful when running a local Docker Compose with services for testing purposes.

It allows keeping separate data and migration files while applying them using a single container app.

## How to Use

The migration files must be in a format recognized by [migrate](https://github.com/golang-migrate/migrate).

### Create a local image

1. Clone this repository.

2. Create a Docker image.

3. Add a service in your `docker-compose.yml` file.

### Use an existing image

1. Pull the image from this repository package.

2. Add a service in your `docker-compose.yml` file.

For more details and examples, refer to the [docker-compose](./docker-compose.yml) file.

## Contributing

Contributions are welcome! If you find a bug or would like to suggest an enhancement, please make a fork, create a new branch with the bugfix or feature, and submit a pull request.

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) file for more information.
