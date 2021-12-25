# Build From Source
$loc = "$HOME\AppData\Local\resto"

go run build/date.go >> date.txt

$LATEST_VERSION=git describe --abbrev=0 --tags
$DATE=cat date.txt

# Build
go mod tidy
go build -o resto.exe -ldflags "-X main.version=$LATEST_VERSION -X main.versionDate=$DATE"

# Setup
$BIN = "$loc\bin"
New-Item -ItemType "directory" -Path $BIN
Move-Item resto.exe -Destination $BIN
[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$BIN", [System.EnvironmentVariableTarget]::User)

if (Test-Path -path $loc) {
    Write-Host "Resto was built successfully, refresh your powershell and then run 'resto --help'" -ForegroundColor DarkGreen
} else {
    Write-Host "Build failed" -ForegroundColor Red
}
