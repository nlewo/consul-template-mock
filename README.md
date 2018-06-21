`consul-template-mock` eats a JSON file to render a [`consul-template`](https://github.com/hashicorp/consul-template) template
for testing and development purposes.


### Usage

```
$ consul-template-mock examples/trivial.tmpl examples/trivial.json
Rendered without Consul :/
```

where [examples/trivial.tmpl](examples/trivial.tmpl) is a
`consul-template` template file and
[examples/trivial.json](examples/trivial.json) is a JSON file
describing input datas (see below for its format).


### Installation

`go get github.com/nlewo/consul-template-mock`


### Input JSON file format

The JSON top level attributes are
- service: object where keys are service name
- secret: object where keys are secret name
- key: object where keys are Consul key and value associated value
- env: object where keys are environment variable name and value the variable value
- file: object where keys are file name and value file contents

Check the `./examples` directory for examples!


### Limitations

Only `consul-template` functions I use are mocked, so only a subpart
of `consul-template` language is supported. But contributions are
welcome.
