FROM node:21-alpine3.18 as build
WORKDIR /app

ENV MATTER_SERVER_UI_VERSION="v0.2.0"

RUN \
    apk add --no-cache git \
    && git clone https://github.com/wozz/matter-server-ui.git \
    && cd matter-server-ui \
    && git checkout ${MATTER_SERVER_UI_VERSION} \
    && npm install \
    && npm run build

FROM nginx:alpine3.18

COPY --from=build /app/matter-server-ui/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
