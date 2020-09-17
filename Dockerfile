FROM gcr.io/distroless/static

ARG BINARY=suspenders-linux-amd64
COPY dist/${BINARY} /suspenders
COPY ui/dist /ui/dist

EXPOSE 8080
USER nobody
CMD [ "/suspenders" ]
