FROM golang:1.23.4 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/dbx -a -ldflags '-linkmode external -extldflags "-static"' .

FROM gcr.io/distroless/static:nonroot

# `nonroot` coming from distroless
USER 65532:65532

# Copy the binary that goreleaser built
COPY --from=builder /go/bin/dbx /bin/dbx

# Run the web service on container startup.
ENTRYPOINT [ "/bin/dbx" ]
CMD ["serve"]
