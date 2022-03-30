FROM debian:11

### variables ###
ARG UPD="apt-get update"
ARG UPD_s="sudo $UPD"
ARG INS="apt-get install"
ARG INS_s="sudo $INS"

RUN $UPD && $INS -y build-essential software-properties-common && $UPD && \
    locale-gen en_US.UTF-8 && \
    mkdir /var/lib/apt/abdcodedoc-marks && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* && \
    $UPD

ENV LANG=en_US.UTF-8

### git ###
RUN $INS -y git && \
    rm -rf /var/lib/apt/lists/* && \
    $UPD

### sudo ###
RUN $UPD && $INS -y sudo && \
    adduser --disabled-password --gecos '' rs && \
    adduser rs sudo && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

ENV HOME="/home/rs"
WORKDIR $HOME
USER rs

### go ###
COPY --from=golang /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

### resto ###
RUN go install github.com/abdfnx/resto@latest
