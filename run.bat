@echo off
setlocal

:install_winget
echo Installing winget...
powershell -Command "Start-Process ms-windows-store://pdp/?productid=9NBLGGH4NNS1 -Wait -NoNewWindow"
echo Please install winget from the Microsoft Store and then re-run this script.
pause
exit /b 1

:check_winget
winget --version >nul 2>&1
if %errorlevel% neq 0 (
    echo winget is not installed. Attempting to install winget...
    call :install_winget
    pause
    exit /b 1
) else (
    echo winget is installed.
)

call :check_winget

go version >nul 2>&1
if %errorlevel% neq 0 (
    echo Go is not installed. Installing Go...
    winget install --id GoLang.Go --source winget
    if %errorlevel% neq 0 (
        echo Go installation failed. Exiting...
        pause
        exit /b 1
    )
    echo Go installed successfully.
    pause
) else (
    echo Go is already installed.
)

node -v >nul 2>&1
if %errorlevel% neq 0 (
    fnm --version >nul 2>&1
    if %errorlevel% neq 0 (
        echo fnm is not installed. Installing fnm...
        winget install Schniz.fnm
        if %errorlevel% neq 0 (
            echo fnm installation failed. Exiting...
            pause
            exit /b 1
        )
        echo fnm installed successfully.
        pause
    ) else (
        echo fnm is already installed.
    )

    echo Node.js is not installed. Installing Node.js version 20...
    fnm use --install-if-missing 20
    if %errorlevel% neq 0 (
        echo Node.js installation failed. Exiting...
        pause
        exit /b 1
    )
    echo Node.js installed successfully.
    pause
) else (
    echo Node.js is already installed.
)

echo Verifying Node.js installation...
node -v
if %errorlevel% neq 0 (
    echo Node.js installation failed. Exiting...
    pause
    exit /b 1
)

echo Verifying npm installation...
npm -v
if %errorlevel% neq 0 (
    echo npm installation failed. Exiting...
    pause
    exit /b 1
)

echo Installing dependencies...
npm install
if %errorlevel% neq 0 (
    echo Dependency installation failed. Exiting...
    pause
    exit /b 1
)

echo Running frontend server...
start cmd /k "npm run dev"

echo Running backend server...
cd backend
start cmd /k "go run main.go"

echo Opening http://localhost:5173/ in the default browser...
start http://localhost:5173/

endlocal
echo Application is running. Please press Ctrl+C to end the process.
pause