Write-Host "Installing the app..." -ForegroundColor Yellow

New-Item "$($env:TEMP)\rep-checker\" -ItemType Directory | Out-Null

Invoke-WebRequest -Uri "https://github.com/GPGamer98/Repetition-Checker/releases/latest/download/CogitoErgoVet-windows-amd64.zip" -OutFile "$($env:TEMP)\rep-checker\app.zip"

$desktop = [Environment]::GetFolderPath("Desktop")

Expand-Archive -LiteralPath "$($env:TEMP)\rep-checker\app.zip" -DestinationPath $desktop -Force

Remove-Item -Path "$($env:TEMP)\rep-checker" -Recurse

Write-Host "Installation completed. The app should be on the desktop" -ForegroundColor Green

Write-Host -NoNewLine 'Press any key to close...';
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');