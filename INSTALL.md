# Installing the Terraform Provider for HPE UXI

This document assumes the use of Terraform 1.7.0 or later.


## Installation

The latest release of the provider can be found on [releases](https://github.com/aruba-uxi/terraform-provider-hpeuxi/releases). You can download the appropriate version of the provider for your operating system using a command line shell or a browser.

This can be useful in environments that do not allow direct access to the Internet.

## Install Scripts

See [linux](./scripts/install-hpeuxi-provider.sh), [windows](./scripts/install-hpeuxi-provider-windows.ps1), and
[macOS](./scripts/install-hpeuxi-provider-macos.sh) install scripts that will download the latest release of the provider and install it in the appropriate location for your operating system.

Then skip to [Configure the Terraform Configuration Files](#configure-the-terraform-configuration-files) to complete the steps.

### Linux

The following examples use Bash on Linux (x64).

1. On a Linux operating system with Internet access, download the plugin from GitHub using the shell.

   ```console
   RELEASE=x.y.z
   wget -q https://github.com/aruba-uxi/terraform-provider-hpeuxi/releases/download/v${RELEASE}/terraform-provider-hpeuxi_${RELEASE}_linux_amd64.zip
   ```

2. Extract the plugin.

   ```console
   unzip terraform-provider-hpeuxi_${RELEASE}_linux_amd64.zip -d terraform-provider-hpeuxi_${RELEASE}
   ```

3. Create a directory for the provider.

   > **Note**
   >
   > The directory hierarchy that Terraform uses to precisely determine the source of each provider it finds locally.
   >
   > `$PLUGIN_DIRECTORY/$SOURCEHOSTNAME/$SOURCENAMESPACE/$NAME/$VERSION/$OS_$ARCH/`

   ```console
   mkdir -p ~/.local/share/terraform/plugins/registry.terraform.io/HewlettPackard/hpeuxi/${RELEASE}/linux_amd64
   ```

4. Copy the extracted plugin to a target system and move to the Terraform plugins directory.

   ```console
   mv terraform-provider-hpeuxi_${RELEASE}/terraform-provider-hpeuxi_v${RELEASE} ~/.local/share/terraform/plugins/registry.terraform.io/HewlettPackard/hpeuxi/${RELEASE}/linux_amd64
   ```

5. Verify the presence of the plugin in the Terraform plugins directory.

   ```console
   cd ~/.local/share/terraform/plugins/registry.terraform.io/HewlettPackard/hpeuxi/${RELEASE}/linux_amd64
   ls
   ```

### macOS

The following example uses Zsh (default) on macOS (Apple Silicon).

1. On a macOS operating system with Internet access, install wget with [Homebrew](https://brew.sh).

   ```console
   brew install wget
   ```

2. Download the plugin from GitHub using the shell.

   ```console
   RELEASE=x.y.z
   wget https://github.com/aruba-uxi/terraform-provider-hpeuxi/releases/download/v${RELEASE}/terraform-provider-hpeuxi_${RELEASE}_darwin_arm64.zip
   ```

3. Extract the plugin.

   ```console
   unzip terraform-provider-hpeuxi_${RELEASE}_darwin_arm64.zip -d terraform-provider-hpeuxi_${RELEASE}
   ```

4. Create a directory for the provider.

   > **Note**
   >
   > The directory hierarchy that Terraform uses to precisely determine the source of each provider it finds locally.
   >
   > `$PLUGIN_DIRECTORY/$SOURCEHOSTNAME/$SOURCENAMESPACE/$NAME/$VERSION/$OS_$ARCH/`

   ```console
   mkdir -p ~/.terraform.d/plugins/registry.terraform.io/HewlettPackard/hpeuxi/${RELEASE}/darwin_arm64
   ```

5. Copy the extracted plugin to a target system and move to the Terraform plugins directory.

   ```console
   mv terraform-provider-hpeuxi_${RELEASE}/terraform-provider-hpeuxi_v${RELEASE} ~/.terraform.d/plugins/registry.terraform.io/HewlettPackard/hpeuxi/${RELEASE}/darwin_arm64
   ```

6. Verify the presence of the plugin in the Terraform plugins directory.

   ```console
   cd ~/.terraform.d/plugins/registry.terraform.io/HewlettPackard/hpeuxi/${RELEASE}/darwin_arm64
   ls
   ```

### Windows

The following examples use PowerShell on Windows (x64).

1. On a Windows operating system with Internet access, download the plugin using the PowerShell.

   ```powershell
   Set-Variable -Name "RELEASE" -Value "x.y.z"
   Invoke-WebRequest https://github.com/aruba-uxi/terraform-provider-hpeuxi/releases/download/v${RELEASE}/terraform-provider-hpeuxi_${RELEASE}_windows_amd64.zip -outfile terraform-provider-hpeuxi_${RELEASE}_windows_amd64.zip
   ```

2. Extract the plugin.

   ```powershell
   Expand-Archive terraform-provider-hpeuxi_${RELEASE}_windows_amd64.zip
   cd terraform-provider-hpeuxi_${RELEASE}_windows_amd64
   ```

3. Copy the extracted plugin to a target system and move to the Terraform plugins directory.

   > **Note**
   >
   > The directory hierarchy that Terraform uses to precisely determine the source of each provider it finds locally.
   >
   > `$PLUGIN_DIRECTORY/$SOURCEHOSTNAME/$SOURCENAMESPACE/$NAME/$VERSION/$OS_$ARCH/`

   ```powershell
   New-Item $ENV:APPDATA\terraform.d\plugins\registry.terraform.io\aruba-uxi\hpeuxi\${RELEASE}\ -Name "windows_amd64" -ItemType "directory"

   Move-Item terraform-provider-hpeuxi_v${RELEASE}.exe $ENV:APPDATA\terraform.d\plugins\registry.terraform.io\aruba-uxi\hpeuxi\${RELEASE}\windows_amd64\terraform-provider-hpeuxi_v${RELEASE}.exe
   ```

4. Verify the presence of the plugin in the Terraform plugins directory.

   ```powershell
   cd $ENV:APPDATA\terraform.d\plugins\registry.terraform.io\aruba-uxi\hpeuxi\${RELEASE}\windows_amd64
   dir
   ```

### Configure the Terraform Configuration Files

A working directory can be initialized with providers that are installed locally on a system by using `terraform init`. The Terraform configuration block is used to configure some behaviors of Terraform itself, such as the Terraform version and the required providers source and version.

**Example**: A Terraform configuration block.

```hcl
terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/HewlettPackard/hpeuxi"
      version = ">= x.y.z"
    }
  }
  required_version = ">= 1.7.0"
}
```

### Verify the Terraform Initialization of a Manually Installed Provider

To verify the initialization, navigate to the working directory for your Terraform configuration and run `terraform init`. You should see a message indicating that Terraform has been successfully initialized and the installed version of the Terraform Provider for HPE UXI.

**Example**: Initialize and Use a Manually Installed Provider

```console
% terraform init

Initializing the backend...

Initializing provider plugins...
- Reusing previous version of HewlettPackard/hpeuxi from the dependency lock file
- Installing HewlettPackard/hpeuxi x.y.z...
- Installed HewlettPackard/hpeuxi x.y.z (unauthenticated)

Terraform has been successfully initialized!
```

## Get the Provider Version

To find the provider version, navigate to the working directory of your Terraform configuration and run `terraform version`. You should see a message indicating the provider version.

**Example**: Terraform Provider Version from the Terraform Registry

```console
% terraform version
Terraform x.y.z
on darwin_arm64
+ provider registry.terraform.io/HewlettPackard/hpeuxi x.y.z
```
