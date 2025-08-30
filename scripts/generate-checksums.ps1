# PowerShellスクリプト：詳細なチェックサム生成
param(
    [Parameter(Mandatory=$true)]
    [string]$FilePath
)

$file = Get-Item $FilePath
$sha256 = Get-FileHash $FilePath -Algorithm SHA256
$sha1 = Get-FileHash $FilePath -Algorithm SHA1
$md5 = Get-FileHash $FilePath -Algorithm MD5

$output = @"
# Checksums for $($file.Name)

**File Information:**
- Name: $($file.Name)
- Size: $($file.Length) bytes
- Created: $($file.CreationTime)
- Modified: $($file.LastWriteTime)

**Hash Values:**
- SHA256: $($sha256.Hash)
- SHA1: $($sha1.Hash)
- MD5: $($md5.Hash)

**Verification Commands:**

Windows (CMD):
``````
certutil -hashfile $($file.Name) SHA256
``````

Windows (PowerShell):
``````
Get-FileHash $($file.Name) -Algorithm SHA256
``````

Linux/macOS:
``````
sha256sum $($file.Name)
``````
"@

$output | Out-File -FilePath "checksums.md" -Encoding UTF8
Write-Host "Checksums generated in checksums.md"