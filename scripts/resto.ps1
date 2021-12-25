# get latest release
$release_url = "https://api.github.com/repos/abdfnx/resto/releases"
$tag = (Invoke-WebRequest -Uri $release_url -UseBasicParsing | ConvertFrom-Json)[0].tag_name
$loc = "$HOME\AppData\Local\resto"
$url = ""
$arch = $env:PROCESSOR_ARCHITECTURE

if ($arch -eq "AMD64") {
    $url = "https://github.com/abdfnx/resto/releases/download/$tag/resto_windows_${tag}_amd64.zip"
} elseif ($arch -eq "x86") {
    $url = "https://github.com/abdfnx/resto/releases/download/$tag/resto_windows_${tag}_386.zip"
} elseif ($arch -eq "arm") {
    $url = "https://github.com/abdfnx/resto/releases/download/$tag/resto_windows_${tag}_arm.zip"
} elseif ($arch -eq "arm64") {
    $url = "https://github.com/abdfnx/resto/releases/download/$tag/resto_windows_${tag}_arm64.zip"
}

if (Test-Path -path $loc) {
    Remove-Item $loc -Recurse -Force
}

Write-Host "Installing resto version $tag" -ForegroundColor DarkCyan

Invoke-WebRequest $url -outfile resto_windows.zip

Expand-Archive resto_windows.zip

New-Item -ItemType "directory" -Path $loc

Move-Item -Path resto_windows\bin -Destination $loc

Remove-Item resto_windows* -Recurse -Force

[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$loc\bin", [System.EnvironmentVariableTarget]::User)

if (Test-Path -path $loc) {
    Write-Host "Thanks for installing Resto! Refresh your powershell" -ForegroundColor DarkGreen
    Write-Host "If this is your first time using the CLI, be sure to run 'resto --help' first." -ForegroundColor DarkGreen
} else {
    Write-Host "Download failed"
    Write-Host "Please try again later"
}
