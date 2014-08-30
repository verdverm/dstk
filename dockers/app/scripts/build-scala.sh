#!/bin/bash

cd /app/usercode

if [ ! -f scala-library.jar ]
then
    cp /usr/lib/scala/lib/scala-library.jar .
fi

if [ ! -f spark-assembly-1.0.2-hadoop2.2.0.jar ]
then
	cp /apache/spark-1.0.2-bin-hadoop2/lib/spark-assembly-1.0.2-hadoop2.2.0.jar .
fi


echo "compiling scala"
scalac -cp "*.jar" -sourcepath src -d bin "/app/usercode/src/$APPNAME/$APPNAME.scala"

echo "linking jars"
cd bin
jar -cfm ../$APPNAME.jar ../MANIFEST.MF *
cd ..

# echo "Press [CTRL+C] to stop.."
# while :
# do
# 	sleep 60
# done

# rm scala-library.jar
# java -jar wordcount.jar

