#!/bin/bash -e

# Increment the version 
# $1 == Project to version
# $2 == Major/Minor version
# e.g. ./version.sh PondDisplay Major

ProjectPath=$1
VersionToUpdate=$2

echo Incrementing $VersionToUpdate version of $ProjectPath

# Get the current version
CurrentVersion=$(cat $ProjectPath/VERSION)

echo $ProjectPath currently at version $CurrentVersion
IFS='.' read -ra CurrentVersionArray <<< "$CurrentVersion"

MajorVersion=${CurrentVersionArray[0]}
MinorVersion=${CurrentVersionArray[1]}

# Increment the version
if [ "$VersionToUpdate" = "Major" ]
then
	MajorVersion=$(($MajorVersion+1))
elif [ "$VersionToUpdate" = "Minor" ]
then
	MinorVersion=$(($MinorVersion+1))
else
	echo "Invalid argument. Syntax is: version.sh <Project Folder> <Major/Minor>"
fi

# Increment the version
NewVersion=$MajorVersion.$MinorVersion
echo Updating $ProjectPath to $NewVersion
echo $NewVersion > $ProjectPath/VERSION

# Check it back into SCM
git add $ProjectPath/VERSION
git commit -m "v$NewVersion"
