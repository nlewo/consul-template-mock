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
    "a_service": [{"Name":"with_its_name"}]},
  "key": {
      "a_key": "with_its_value"
  },
  "env": {"a_environment_variable": "with_its_value"},
  "secret": {"a_secret_path":
             {"a_secret": "****"}},
  "file": {"a_filepath": "with_its_content"}
}
```

See the `./examples` directory for more examples!


### Limitations

Only `consul-template` functions that I use are mocked, so just a
subpart of `consul-template` language is currently
supported. Contributions are welcome!
