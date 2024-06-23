@echo off
setlocal

:install_winget
echo Installing winget...
powershell -Command "Start-Process ms-windows-store://pdp/?productid=9NBLGGH4NNS1 -Wait -NoNewWindow"
echo Please install winget from the Microsoft Store and then re-run this script.
exit /b 1

winget --version >nul 2>&1
if %errorlevel% neq 0 (
    echo winget is not installed. Attempting to install winget...
    call :install_winget
) else (
    echo winget is installed.
)

go version >nul 2>&1
if %errorlevel% neq 0 (
    echo Go is not installed. Installing Go...
    winget install --id GoLang.Go --source winget
) else (
    echo Go is already installed.
)

fnm --version >nul 2>&1
if %errorlevel% neq 0 (
    echo fnm is not installed. Installing fnm...
    winget install Schniz.fnm
) else (
    echo fnm is already installed.
)

echo Checking Node.js version 20...
fnm use --install-if-missing 20

echo Verifying Node.js installation...
node -v
if %errorlevel% neq 0 (
    echo Node.js installation failed. Exiting...
    exit /b 1
)

echo Verifying npm installation...
npm -v
if %errorlevel% neq 0 (
    echo npm installation failed. Exiting...
    exit /b 1
)

echo Installing node modules and dependencies...
npm install

echo Running frontend server...
start cmd /k "npm run dev"

echo Running backend server...
cd backend
start cmd /k "go run main.go"

echo Opening http://localhost:5173/ in the default browser...
start http://localhost:5173/

endlocal
echo "Application is running. Please press Ctrl+C to end the process."
pause
