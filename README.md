# serveit
A very simple HTTP file server

## Installation

1. Download the binary for your plaform from [here](https://github.com/samherrmann/serveit/releases/).
2. Move the binary to an installation directory of your choice.
3. Add the installation directory to your environment path.

## Usage

```shell
> cd root/of/project/you/want/to/serve
> serveit
Serving current directory on port :8080
```

Serving on a custom port...
```shell
> serveit -port=7070
Serving current directory on port :7070
```

Serving a custom file when the requested resource cannot be found...
```shell
> serveit -not-found-file=404.html
```
For single-page applications, this flag is typically set to `index.html`.

## Developing

### Required Tools

* [Go](https://golang.org/)
* Text editor (Recommended: [VSCode](https://code.visualstudio.com/) with 
   [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go))

### Build it from source

```shell
> clone https://github.com/samherrmann/serveit.git
> cd root/of/cloned/repo
> go build
```




