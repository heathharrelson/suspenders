FROM gcr.io/distroless/static

ARG BINARY=suspenders-linux-amd64
COPY dist/${BINARY} /suspenders
COPY templates /templates
COPY ui/static /ui/static

EXPOSE 8080
USER nobody
CMD [ "/suspenders" ]
