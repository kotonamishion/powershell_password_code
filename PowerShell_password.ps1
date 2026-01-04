# 1. 解决乱码
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# 2. 界面显示
Clear-Host
Write-Host "==============================" -ForegroundColor Cyan
Write-Host "   Windows 极简安全密码生成器" -ForegroundColor Cyan
Write-Host "==============================`n" -ForegroundColor Cyan

# --- 修改处：默认值设为 12 ---
$input = Read-Host "请输入密码长度 (直接回车默认为 12)"
if ($input -match '^\d+$') { 
    $length = [int]$input 
} else { 
    $length = 12 
}

# 3. 字符池
$chars = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!@#$%^&*()_+"
$password = ""

for ($i = 0; $i -lt $length; $i++) {
    $password += $chars[(Get-Random -Maximum $chars.Length)]
}

# 4. 输出并复制
Write-Host "`n生成成功！长度为: $length" -ForegroundColor Green
Write-Host "您的密码是: " -NoNewline
Write-Host $password -ForegroundColor Yellow

$password | clip
Write-Host "`n[提示] 密码已自动复制到剪贴板。" -ForegroundColor Gray

Write-Host "`n------------------------------"
Write-Host "程序将在 30 秒后自动清空剪贴板并退出..."
Write-Host "或者按回车键立即退出。"

# 5. 安全清理
$wait = 0
while($wait -lt 30 -and -not [console]::KeyAvailable) {
    Start-Sleep -Seconds 1
    $wait++
}

"" | clip
Clear-Host
Write-Host "剪贴板已清空，会话已安全退出。" -ForegroundColor Cyan
Start-Sleep -Seconds 1