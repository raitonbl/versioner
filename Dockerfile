FROM ubuntu:latest
ADD dist/versioner /usr/local/bin/versioner
RUN chmod a+x /usr/local/bin/versioner
ENTRYPOINT ["versioner"]
CMD ["help"]