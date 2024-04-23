# Subnet Server

Subnet Server is a lightweight, secure web server designed to execute `ipc-cli` commands and expose them through a RESTful API. It's ideal for integrating with other systems or providing a network-based interface to `ipc-cli` functionalities.

## Features

- **Simple API**: Exposes `ipc-cli` commands through HTTP POST requests.
- **Flexible Configuration**: Supports custom port and optional authorization token for enhanced security.
- **Easy Deployment**: Quick to set up with minimal configuration, perfect for both development and production environments.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installing

Subnet Server can be downloaded directly as a pre-built binary from the releases page. Use the following `curl` command to download the latest release:

```bash
curl -L -O https://github.com/myesterylabs/subnet-server/releases/download/v0.0.3/subnet-server-linux-amd64

chmod +x subnet-server-linux-amd64
```


Running the Server
Run the server using the following command:

./subnet_server --port 8818 --authToken "your_auth_token"

--port: Specifies the port number on which the server will listen (default is 8818).

--authToken: An optional security token for client authentication. If specified, all incoming requests must include this token in their headers. (default is 098790879089789)


```

### Opening the Port
Ensure that the specified port is open on your server to allow traffic. The method to open a port varies depending on your server's operating system and firewall settings.

### Example Usage
You can interact with the server using curl as shown in this example:

```bash
curl -X POST http://localhost:8818/ipc-cli/ls \
     -H "Authorization: your_auth_token" \
     -d '{"args":["-l", "/tmp"]}' \
     -H "Content-Type: application/json"
```

### Security
If you choose to use the --authToken flag, ensure that your token is kept secure. Do not expose the server to the internet without proper authentication and data validation mechanisms in place.

### Contributing
Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change.

### License
This project is licensed under the MIT License - see the LICENSE.md file for details.

### Acknowledgments
Thanks to everyone who has contributed to developing ipc-cli.
Inspired by the needs of network administrators and automation enthusiasts.






