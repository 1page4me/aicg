@echo off
echo Setting up PostgreSQL database...

REM Check if PostgreSQL is installed
where psql >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo PostgreSQL is not installed or not in PATH.
    echo Please follow these steps:
    echo 1. Download PostgreSQL from https://www.postgresql.org/download/windows/
    echo 2. Run the installer
    echo 3. During installation:
    echo    - Remember the password you set for postgres user
    echo    - Check "Add PostgreSQL to PATH"
    echo 4. After installation, run this script again
    pause
    exit /b 1
)

REM Check if PostgreSQL service is running
sc query postgresql-x64-15 >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo PostgreSQL service is not running.
    echo Starting PostgreSQL service...
    net start postgresql-x64-15
    if %ERRORLEVEL% neq 0 (
        echo Failed to start PostgreSQL service.
        echo Please start it manually from Services (services.msc)
        pause
        exit /b 1
    )
)

REM Set environment variables
set PGPASSWORD=your_password

REM Create database and user
echo Creating database and user...
psql -U postgres -f setup_db.sql

if %ERRORLEVEL% neq 0 (
    echo Failed to set up database.
    echo Please check:
    echo 1. PostgreSQL service is running
    echo 2. Password in setup_db.sql matches your PostgreSQL password
    echo 3. User 'postgres' exists and has correct permissions
    pause
    exit /b 1
)

echo Database setup completed successfully!
echo You can now start the application.
pause 