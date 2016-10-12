set DEF_OUT=E:\code\unity\snake\Assets\Plugins
cd precompile
precompile ..\DLL\MessageDefine.dll -o:ProtobufSerializer.dll -t:Snake3D.ProtobufSerializer

copy ..\DLL\MessageDefine.dll %DEF_OUT%\MessageDefine.dll

copy ProtobufSerializer.dll %DEF_OUT%\ProtobufSerializer.dll

rem copy protobuf-net.dll %DEF_OUT%\protobuf-net.dll

rem svn commit %DEF_OUT%\MessageDefine.dll %DEF_OUT%\ProtobufSerializer.dll  -m "protobuf gen"
pause