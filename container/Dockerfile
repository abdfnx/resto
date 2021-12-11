FROM debian:latest

### variables ###
ARG UPD="apt-get update"
ARG UPD_s="sudo $UPD"
ARG INS="apt-get install"
ARG INS_s="sudo $INS"
ENV PKGS="zip unzip multitail curl lsof wget ssl-cert asciidoctor apt-transport-https ca-certificates gnupg-agent bash-completion build-essential htop jq software-properties-common less llvm locales man-db nano vim ruby-full"
ENV BUILDS="build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev libbz2-dev"

RUN $UPD && $INS -y $PKGS && $UPD && \
    locale-gen en_US.UTF-8 && \
    mkdir /var/lib/apt/abdcodedoc-marks && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* && \
    $UPD

ENV LANG=en_US.UTF-8

### git ###
RUN $INS -y git && \
    rm -rf /var/lib/apt/lists/* && \
    $UPD

### go ###
COPY --from=golang:1.17-alpine /usr/local/go/ /usr/local/go/

### sudo ###
RUN $UPD && $INS -y sudo && \
    adduser --disabled-password --gecos '' rs && \
    adduser rs sudo && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

ENV HOME="/home/rs"
WORKDIR $HOME
USER rs

### resto ###
RUN go install github.com/abdfnx/resto@latest

### zsh ###
ENV src=".zshrc"

RUN $INS_s zsh -y
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    $UPD_s && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

### rm old ~/.zshrc ###
RUN sudo rm -rf $src

### wget new files ###
RUN wget https://docker.secman.dev/resto/.zshrc
RUN wget https://docker.secman.dev/resto/README

CMD /bin/bash -c "cat README && zsh"
