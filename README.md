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
describing input mock datas (see below for its format).


### Installation

`go get github.com/nlewo/consul-template-mock`


### Mock JSON file format

```json
{ "service": {
    "simple": [{"Name":"simple"}]},
  "key": {
      "/simple": "simple"
  },
  "env": {"simple": "simple"},
  "secret": {"secret/simple":
             {"simple": "simple"}},
  "file": {"/simple": "simple"}
}
```

See the `./examples` directory for more examples!


### Limitations

Only `consul-template` functions that I use are mocked, so just a
subpart of `consul-template` language is currently
supported. Contributions are welcome!
