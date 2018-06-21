`consul-template-mock` eats a JSON file to render a Consul template


### Usage

```$ consul-template-mock consul-template.tmpl data.json```

where `consul-template.tmpl` is a consul template file and `data.json`
is a JSON file describing datas.


### Build

`go build consul-template-mock.go`


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
