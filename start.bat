@echo off
REM GoSecureTransfer Start Script for Windows

setlocal enabledelayedexpansion

color 0A
echo.
echo ========================================
echo Starting GoSecureTransfer...
echo ========================================
echo.

REM Check if required tools are available
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Go is not installed or not in PATH
    pause
    exit /b 1
)

where npm >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] npm is not installed or not in PATH
    pause
    exit /b 1
)

REM Start Backend in a new window with labeled output
echo [BACKEND] Starting backend server...
start "GoSecureTransfer - Backend" cmd /c "cd backend && echo. && echo [BACKEND] Backend server starting... && echo. && go run ./cmd/server/main.go 2>&1 | findstr /R /C:"^" && pause"

REM Wait a moment for backend to start
timeout /t 2 /nobreak

REM Start Frontend in a new window with labeled output
echo [FRONTEND] Starting frontend dev server...
start "GoSecureTransfer - Frontend" cmd /c "cd frontend && echo. && echo [FRONTEND] Frontend dev server starting... && echo. && npm run dev"

REM Display info
echo.
echo ========================================
echo Both services are starting...
echo ========================================
echo [BACKEND]  - Server running in separate window
echo [FRONTEND] - Dev server running in separate window
echo.
echo Press Ctrl+C in each window to stop the respective service
echo.
pause
