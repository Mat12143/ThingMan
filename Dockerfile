FROM node:23-alpine3.20 AS tailwindBuild

WORKDIR /build

COPY . .
RUN npm i
RUN npm run build

FROM golang:1.24.2-alpine AS build

WORKDIR /build

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN apk add --no-cache gcc musl-dev sqlite-dev

RUN templ generate
RUN CGO_ENABLED=1 go build -ldflags="-extldflags=-static" -o ./thingman .

FROM scratch

WORKDIR /app

COPY --from=tailwindBuild /build/static/style.css ./static/style.css
COPY --from=build /build/static/alpine.js ./static/alpine.js

COPY --from=build /build/thingman .

ENTRYPOINT [ "./thingman" ]
