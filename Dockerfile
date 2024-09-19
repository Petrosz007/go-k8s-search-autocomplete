FROM debian
COPY ./target/search-autocomplete /search-autocomplete
ENTRYPOINT ["/search-autocomplete"]
