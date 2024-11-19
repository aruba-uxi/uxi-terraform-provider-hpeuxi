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

## Building and Distribution

All builds must be signed by HPE Code Sign before distribution.
Read the [Notices Report](public/README.md) and ensure a report is added to every build which is published
All open source sourcecode must be submitted to the HPE DSM to make this available for customers to download on request.
