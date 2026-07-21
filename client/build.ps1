param(
    [string]$Output = "tcpstunc_windows_amd64.exe"
)

$previousGOOS = $env:GOOS
$previousGOARCH = $env:GOARCH

try {
    $env:GOOS = "windows"
    $env:GOARCH = "amd64"
    Push-Location $PSScriptRoot
    try {
        go build -ldflags="-s -w" -o $Output .
        if ($LASTEXITCODE -ne 0) {
            exit $LASTEXITCODE
        }
    }
    finally {
        Pop-Location
    }
}
finally {
    $env:GOOS = $previousGOOS
    $env:GOARCH = $previousGOARCH
}
