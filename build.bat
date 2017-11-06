set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\gochart_example\=\%
set GOPATH=%BASEDIR%;%CURDIR%\Godeps
set GOBIN=%CURDIR%\bin
go install .
pause