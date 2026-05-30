$RepoDomain = "github.com"
$RepoOwner = "spenserblack"
$RepoName = "zal-go"
$ProjectName = "zalgo"
$InstallDir = "$Env:LOCALAPPDATA\Programs\$ProjectName"
$InstallPath = "$InstallDir\$ProjectName.exe"

# Architectures
$Amd64 = "amd64"

$Os = "windows"

$Repo = "$RepoDomain/$RepoOwner/$RepoName"

if ($Env:PROCESSOR_ARCHITECTURE -eq "AMD64") {
    $Arch = "$Amd64"
}
else {
    Write-Error "Unknown processor architecture: $Env:PROCESSOR_ARCHITECTURE"
    exit 1
}

$AssetName = "$ProjectName-$Os-$Arch"
$CompressedAssetName = "$AssetName.tar.gz"
$DecompressedAssetName = "$AssetName.exe"
$Url = "https://$Repo/releases/latest/download/$CompressedAssetName"

$TempDir = [System.IO.Path]::GetTempPath()
$DownloadPath = "$TempDir" + "$CompressedAssetName"

Write-Output "Downloading $Url to $DownloadPath"
Invoke-WebRequest -Uri $Url -OutFile $DownloadPath

Write-Output "Creating $InstallDir..."
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

Write-Output "Extracting from $DownloadPath..."
tar.exe -xzf "$DownloadPath"

Write-Output "Installing $DecompressedAssetName to $InstallPath..."
Move-Item -Path $DecompressedAssetName -Destination $InstallPath -Force

$UserPath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "Adding to PATH..."
    [System.Environment]::SetEnvironmentVariable("PATH", "$UserPath;$InstallDir", "User")
    $Env:PATH += ";$InstallDir"
}

Write-Host "Cleaning up $DownloadPath..."
Remove-Item -Path $DownloadPath

Write-Host "Done!"
Write-Host "You may need to restart your terminal"
