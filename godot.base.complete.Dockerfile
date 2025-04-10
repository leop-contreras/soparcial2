#docker build -t godot -f godot.base.complete.Dockerfile .
#docker run godot

FROM ubuntu:24.04

RUN apt update && apt install -y wget unzip libfontconfig1

RUN wget https://github.com/godotengine/godot/releases/download/4.4-stable/Godot_v4.4-stable_linux.x86_64.zip -O /tmp/godot.zip \
&& unzip /tmp/godot.zip -d /usr/local/bin/ \
&& mv /usr/local/bin/Godot_v4.4-stable_linux.x86_64 /usr/local/bin/godot \
&& rm /tmp/godot.zip

ENV PATH="/usr/local/bin:${PATH}"

CMD ["godot", "--headless"]