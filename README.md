# A simple Go script that scrapes a website in user-provided durations and exits upon detecting change
Command line flags:
- -d : Duration between scrapes (in Go Duration format e.g. "10m") : Default 10m
- -w : Website to scrape : Default https://info.mik.pte.hu/hu/zarovizsga
