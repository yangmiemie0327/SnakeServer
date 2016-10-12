@echo off
set DEF_OUT=../../src/server/msg/snake/
set protofile=snake.proto
protoc --go_out=%DEF_OUT% ./%protofile%
cd ../client
copy ..\server\%protofile% .\protocol\
protogen -i:.\protocol\snake.proto -o:.\cs_source\MessageDefine.cs -p:detectMissing -ns:Snake3D
call todll_release.bat
pause