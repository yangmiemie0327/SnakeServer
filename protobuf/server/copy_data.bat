@echo off
set DATA_PATH=..\..\bin\gamedata
if exist "%DATA_PATH%" rmdir /s /q "%DATA_PATH%"
mkdir "%DATA_PATH%"
copy ..\..\..\..\unity\snake\Assets\Resources\Configs\*.xml %DATA_PATH%
pause