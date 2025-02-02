FROM golang:1.23

ENV BIN_FILE /opt/calendar/calendar-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar/*

ENV CONFIG_FILE /etc/calendar/calendar_config.toml
COPY ./configs/calendar_config.toml ${CONFIG_FILE}

ENV CONFIG_MIGRATION_DIR /etc/calendar/migrations
COPY ./migrations ${CONFIG_MIGRATION_DIR}

CMD ${BIN_FILE} -config ${CONFIG_FILE}