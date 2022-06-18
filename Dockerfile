FROM ubuntu:latest
ADD dist/versioner /usr/local/bin/versioner
RUN export PATH=/usr/local/bin/versioner:${PATH}
RUN chmod a+x /usr/local/bin/versioner
ENTRYPOINT ["versioner"]