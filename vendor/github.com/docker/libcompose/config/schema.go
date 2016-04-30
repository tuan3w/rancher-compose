
package config

var schemaString = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "config_schema_v1.json",

  "type": "object",

  "patternProperties": {
    "^[a-zA-Z0-9._-]+$": {
      "$ref": "#/definitions/service"
    }
  },

  "additionalProperties": false,

  "definitions": {
    "service": {
      "id": "#/definitions/service",
      "type": "object",

      "properties": {
        "build": {"type": "string"},
        "cap_add": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cap_drop": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "certs": {"$ref": "#/definitions/list_of_strings"},
        "cgroup_parent": {"type": "string"},
        "command": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "container_name": {"type": "string"},
        "cpu_shares": {"type": ["number", "string"]},
        "cpu_quota": {"type": ["number", "string"]},
        "cpuset": {"type": "string"},
        "devices": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "default_cert": {"type": "string"},
        "disks": {"type": "array"},
        "dns": {"$ref": "#/definitions/string_or_list"},
        "dns_search": {"$ref": "#/definitions/string_or_list"},
        "dockerfile": {"type": "string"},
        "domainname": {"type": "string"},
        "entrypoint": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "env_file": {"$ref": "#/definitions/string_or_list"},
        "environment": {"$ref": "#/definitions/list_or_dict"},

        "expose": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "expose"
          },
          "uniqueItems": true
        },

        "extends": {
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "object",

              "properties": {
                "service": {"type": "string"},
                "file": {"type": "string"}
              },
              "required": ["service"],
              "additionalProperties": false
            }
          ]
        },

        "extra_hosts": {"$ref": "#/definitions/list_or_dict"},
        "external_ips": {"$ref": "#/definitions/list_of_strings"},
        "external_links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "healthcheck": {"type": "object"},
        "hostname": {"type": "string"},
        "image": {"type": "string"},
        "ipc": {"type": "string"},
        "labels": {"$ref": "#/definitions/list_or_dict"},
        "links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "load_balancer_config": {"type": "object"},
        "log_driver": {"type": "string"},
        "log_opt": {"type": "object"},
        "mac_address": {"type": "string"},
        "memory": {"type": ["number", "string"]},
        "mem_limit": {"type": ["number", "string"]},
        "memswap_limit": {"type": ["number", "string"]},
        "metadata": {"type": "object"},
        "net": {"type": "string"},
        "pid": {"type": ["string", "null"]},

        "ports": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "ports"
          },
          "uniqueItems": true
        },

        "privileged": {"type": "boolean"},
        "read_only": {"type": "boolean"},
        "restart": {"type": "string"},
        "retain_ip": {"type": "boolean"},
        "scale": {"type": "number"},
        "security_opt": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "shm_size": {"type": ["number", "string"]},
        "stdin_open": {"type": "boolean"},
        "stop_signal": {"type": "string"},
        "tty": {"type": "boolean"},
        "type": {"type": "string"},
        "update_strategy": {"type": "object"},
        "ulimits": {
          "type": "object",
          "patternProperties": {
            "^[a-z]+$": {
              "oneOf": [
                {"type": "integer"},
                {
                  "type":"object",
                  "properties": {
                    "hard": {"type": "integer"},
                    "soft": {"type": "integer"}
                  },
                  "required": ["soft", "hard"],
                  "additionalProperties": false
                }
              ]
            }
          }
        },
        "user": {"type": "string"},
        "userdata": {"type": "string"},
        "vcpu": {"type": ["number", "string"]},
        "volumes": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "volume_driver": {"type": "string"},
        "volumes_from": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "working_dir": {"type": "string"}
      },

      "dependencies": {
        "memswap_limit": ["mem_limit"]
      },
      "additionalProperties": false
    },

    "string_or_list": {
      "oneOf": [
        {"type": "string"},
        {"$ref": "#/definitions/list_of_strings"}
      ]
    },

    "list_of_strings": {
      "type": "array",
      "items": {"type": "string"},
      "uniqueItems": true
    },

    "list_or_dict": {
      "oneOf": [
        {
          "type": "object",
          "patternProperties": {
            ".+": {
              "type": ["string", "number", "null", "boolean"]
            }
          },
          "additionalProperties": false
        },
        {"type": "array", "items": {"type": "string"}, "uniqueItems": true}
      ]
    },

    "constraints": {
      "service": {
        "id": "#/definitions/constraints/service",
        "anyOf": [
          {
            "required": ["build"],
            "not": {"required": ["image"]}
          },
          {
            "required": ["image"],
            "not": {"anyOf": [
              {"required": ["build"]},
              {"required": ["dockerfile"]}
            ]}
          }
        ]
      }
    }
  }
}
`