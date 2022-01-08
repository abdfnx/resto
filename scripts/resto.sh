#!/bin/bash

installPath=$1
restoPath=""

if [ $installPath != "" ]; then
    restoPath=$installPath
else
    restoPath=/usr/local/bin
fi

UNAME=$(uname)
ARCH=$(uname -m)

rmOldFiles() {
    if [ -f $restoPath/resto ]; then
        sudo rm -rf $restoPath/resto*
    fi
}

v=$(curl --silent "https://api.github.com/repos/abdfnx/resto/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

releases_api_url=https://github.com/abdfnx/resto/releases/download

successInstall() {
    echo "🙏 Thanks for installing Resto! If this is your first time using the CLI, be sure to run `resto --help` first."
}

mainCheck() {
    echo "Installing resto version $v"
    name=""

    if [ "$UNAME" == "Linux" ]; then
        if [ $ARCH = "x86_64" ]; then
            name="resto_linux_${v}_amd64"
        elif [ $ARCH = "i686" ]; then
            name="resto_linux_${v}_386"
        elif [ $ARCH = "i386" ]; then
            name="resto_linux_${v}_386"
        elif [ $ARCH = "arm64" ]; then
            name="resto_linux_${v}_arm64"
        elif [ $ARCH = "arm" ]; then
            name="resto_linux_${v}_arm"
        fi

        restoURL=$releases_api_url/$v/$name.zip

        wget $restoURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # resto
        sudo mv $name/bin/resto $restoPath

        rm -rf $name

    elif [ "$UNAME" == "Darwin" ]; then
        if [ $ARCH = "x86_64" ]; then
            name="resto_macos_${v}_amd64"
        elif [ $ARCH = "arm64" ]; then
            name="resto_macos_${v}_arm64"
        fi

        restoURL=$releases_api_url/$v/$name.zip

        wget $restoURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # resto
        sudo mv $name/bin/resto $restoPath

        rm -rf $name

    elif [ "$UNAME" == "FreeBSD" ]; then
        if [ $ARCH = "x86_64" ]; then
            name="resto_freebsd_${v}_amd64"
        elif [ $ARCH = "i386" ]; then
            name="resto_freebsd_${v}_386"
        elif [ $ARCH = "i686" ]; then
            name="resto_freebsd_${v}_386"
        elif [ $ARCH = "arm64" ]; then
            name="resto_freebsd_${v}_arm64"
        elif [ $ARCH = "arm" ]; then
            name="resto_freebsd_${v}_arm"
        fi

        restoURL=$releases_api_url/$v/$name.zip

        wget $restoURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # resto
        sudo mv $name/bin/resto $restoPath

        rm -rf $name
    fi

    # chmod
    sudo chmod 755 $restoPath/resto
}

rmOldFiles
mainCheck

if [ -x "$(command -v resto)" ]; then
    successInstall
else
    echo "Download failed 😔"
    echo "Please try again."
fi
