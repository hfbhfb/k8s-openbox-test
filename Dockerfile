FROM node:14.20.0-slim as node
WORKDIR /frontend
COPY ./frontend/package.json /frontend/package.json
RUN yarn config set registry https://registry.npm.taobao.org -g \
    && yarn config set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass -g \
    && yarn install

COPY ./frontend /frontend
RUN yarn config set registry https://registry.npm.taobao.org -g \
    && yarn config set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass -g \
    && yarn install \
    && yarn build

FROM golang:1.19.1 as golang
WORKDIR /backend
COPY ./backend/go.mod /backend/go.mod
COPY ./backend/go.sum /backend/go.sum
RUN go env -w GOPROXY="https://goproxy.cn,direct" && \
    go mod download
COPY ./backend /backend
COPY --from=node /frontend/dist /backend/codeasset/routers/dist
RUN go env -w GOPROXY="https://goproxy.cn,direct" && \
    go build 

FROM debian:bookworm-20220912-slim
WORKDIR /app
COPY --from=golang /backend/backend /app/backend
COPY --from=golang /backend/config-dev.json /app/config-pro.json
CMD ["./backend"]

