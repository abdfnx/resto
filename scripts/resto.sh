#!/bin/bash

restoPath=/usr/local/bin
UNAME=$(uname)
arch=$(uname -m)

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

    if [ "$UNAME" == "Linux" ]; then
        name="resto_linux_${v}_amd64"
        restoURL=$releases_api_url/$v/$name.zip

        wget $restoURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # resto
        sudo mv $name/bin/resto $restoPath

        rm -rf $name

    elif [ "$UNAME" == "Darwin" ]; then
        name="resto_macos_${v}_amd64"
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
