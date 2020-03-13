FROM flungo/inkscape

WORKDIR /home

CMD ["--help"]
ENTRYPOINT ["/usr/bin/inkscape"]