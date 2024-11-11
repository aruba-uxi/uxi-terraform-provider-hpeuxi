# configuration api terraform provider

A repository to hold the terraform provider as well as the configuration-api client used by the terraform provider.

## Setup provider development environment
Add the appropriate terraform provider dev override to your `~/.terraformrc` file to ensure that that terraform operations is performed against the local provider.

go_path = $(go env GOPATH)/bin

```
dev_overrides {
    "registry.terraform.io/arubauxi/hpeuxi" = "<go_path>"
}
```

Example `~/.terraformrc` file
```
provider_installation {
  dev_overrides {
      "registry.terraform.io/arubauxi/hpeuxi" = "/Users/<user>/go/bin"
  }
  direct {}
}
```
