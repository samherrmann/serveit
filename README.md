# serveit
A simple HTTP file server in a single executable file.

## Installation

1. Download the binary for your plaform from [here](https://github.com/samherrmann/serveit/releases/).
2. Move the binary to an installation directory of your choice.
3. Add the installation directory to your environment path.

## Usage

```sh
cd root/of/project/you/want/to/serve
serveit
```

### Serving on custom port
```sh
serveit -port=7070
```

### Serving custom file when requested resource cannot be found
```sh
serveit -not-found-file=404.html
```
For single-page applications, this flag is typically set to `index.html`.

### Serving over HTTPS
Serveit makes it simple to serve files over HTTPS by automatically creating
self-signed certificates:

```sh
serveit -tls
```
This command automatically generates the following files if they don't already exist:
|                       |                                              |
| --------------------- | -------------------------------------------- |
| `serveit_root_ca.key` | Private key for Root Certificate Aurthority. |
| `serveit_root_ca.crt` | Public certificate for Root Certificate Aurthority. |
| `serveit.key`         | Private key for server. |
| `serveit.crt`         | Public certificate for server, signed with `serveit_root_ca.key` and `serveit_root_ca.crt`. |

Install `serveit_root_ca.crt` as a _Trusted Root Certificate Authority_ in your
client device to have the browser trust the connection to serveit at
`https://localhost:8080`.

#### More HTTPS options
Create certificate for IP address:
```sh
serveit -tls -hosts 192.168.0.1
```
Create certificate for domain name:
```sh
serveit -tls -hosts example.com
```
Create certificate for multiple domain names or IP addresses:
```sh
serveit -tls -hosts localhost,192.168.0.1,example.com
```
Serve on default HTTPS port (443):
```sh
serveit -tls -hosts example.com -port 443
```

## Developing

### Required Tools

* [Go](https://golang.org/)
* [GNU Make](https://www.gnu.org/software/make/)
* Text editor (Recommended: [VSCode](https://code.visualstudio.com/) with 
   [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go))

### Build

```sh
make build
```

### Test

```sh
make test
```




