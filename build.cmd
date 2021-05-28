@echo off
REM *****************************************************
REM * This file will build the tail.exe executable file *
REM *****************************************************

echo Cleaning up
go clean

echo Building binary file for Windows
set GOARCH=amd64
set GOOS=windows
go build -ldflags="-s -w" main.go > NUL
move main.exe tail.exe
echo Done building Windows binary