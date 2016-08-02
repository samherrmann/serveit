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

Serving root "root/of/project/you/want/to/serve" on port :8080
```

Serving on a custom port...
```shell
> serveit --port=7070

Serving root "root/of/project/you/want/to/serve" on port :7070
```

## Build it from source

```shell
> clone https://github.com/samherrmann/serveit.git
> cd root/of/cloned/repo
> go build
```




