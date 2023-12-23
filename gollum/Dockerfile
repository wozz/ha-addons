FROM golang:1.21 as gobuilder

WORKDIR /app
COPY proxy/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prxy .

FROM gollumwiki/gollum:v5.3.2

WORKDIR /data
RUN apk add --upgrade \
            --no-cache \
            curl \
            libcurl \
            jq \
            supervisor \
    && rm -rf /var/cache/apk/*

RUN curl -sSL https://raw.githubusercontent.com/totoroot/gollum-dracula-theme/main/custom.css -o /custom.css

COPY --from=gobuilder /app/prxy /prxy
COPY docker-run.sh /docker-run.sh
RUN chmod +x /docker-run.sh

# fix home button
RUN find / -name navbar.mustache -exec sed -i -e 's#id="minibutton-home" href="{{page_route}}">#id="minibutton-home" href="{{page_route}}/">#g' {} \;

ADD supervisor /supervisor

ENV GOLLUM_AUTHOR_USERNAME=homeassistant
ENV GOLLUM_AUTHOR_EMAIL=homeassistant@ha.local

ENTRYPOINT ["supervisord","-c","/supervisor/service.conf"]
