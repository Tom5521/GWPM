powershell -c "Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser"
powershell -c "Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression"
