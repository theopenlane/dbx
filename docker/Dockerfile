FROM gcr.io/distroless/static:nonroot

# `nonroot` coming from distroless
USER 65532:65532

# Copy the binary that was built
COPY dbx /bin/dbx

# Run the web service on container startup.
ENTRYPOINT [ "/bin/dbx" ]
CMD ["serve"]
