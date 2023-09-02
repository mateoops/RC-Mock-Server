# RC Mock Server

RC Mock Server is a GO web server application (gin) that allow you to create endpoints with anyone possible return code.
Application has easy configuration in yaml.
You can create endpoint with dynamic return code setup by request header value!

## Getting Started

### Installing

Just clone this repo, edit config.yml file, build and run!

### Configuration

YAML config is easy to read and understand.
Apart from static endpoints you can create dynamic return code endpoint by setting up return code for endpoint to 0.
To use dynamic endpoint just send request with header "RC: 405" - where 405 is your expected return code.

## License

This project is licensed under the [MIT License](LICENSE.md)
