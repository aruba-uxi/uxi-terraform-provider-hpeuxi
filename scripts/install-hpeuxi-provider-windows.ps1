param ($VERSION)

$os="windows"
$arch="amd64"
$repo="aruba-uxi/terraform-provider-hpeuxi"
$windows_hpeuxi_dir="$env:appdata\terraform.d\plugins\registry.terraform.io\aruba-uxi\hpeuxi"

$users_pwd = Get-Location

function get_latest_release {
    Write-Host Getting latest release
    $release_url="https://api.github.com/repos/${repo}/releases/latest"
    $tag = (Invoke-WebRequest $release_url | ConvertFrom-Json)[0].tag_name
    $VERSION=${tag}

    $VERSION
}

if (!$VERSION) {
    $VERSION=get_latest_release
}

$version_number=$VERSION -replace 'v'

$dest_dir="${windows_hpeuxi_dir}\${version_number}\${os}_${arch}\"
$hpeuxi_zip="terraform-provider-hpeuxi_${version_number}_${os}_${arch}.zip"
$hpeuxi=$hpeuxi_zip -replace '.zip'
$hpeuxi_dl_url="https://github.com/${repo}/releases/download/${VERSION}/${hpeuxi_zip}"

mkdir "$dest_dir"
Set-Location "$dest_dir"

try {
    Invoke-WebRequest $hpeuxi_dl_url -Out $hpeuxi_zip
}
catch {
    Write-Host "Error: The version that was specified does not exist."

    Set-Location "${users_pwd}"
    Remove-Item -Path "${windows_hpeuxi_dir}\${version_number}" -Recurse -Force -ErrorAction SilentlyContinue

    Write-Host "Exiting..."
    Return
}

Write-Host Extracting release files
Expand-Archive $hpeuxi_zip -Force

Get-ChildItem -Path $hpeuxi -Recurse -File | Move-Item -Destination $dest_dir

Remove-Item $hpeuxi_zip -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item $hpeuxi -Recurse -Force -ErrorAction SilentlyContinue
Write-Host Complete
