This is a very crude CLI utility which can generate a statistics report from a web server access log.

I created this just to learn. Performance, stability, etc etc have not been considered.

I highly recommend GoAccess instead of this. GoAccess is a well built bit of kit. <https://goaccess.io/>

- There's all sorts of hard coded values (like the unit of measurement in the date graphs etc)
- All of the problems from the library used to get the statistics from the access log are inherited. <https://github.com/Tech-Arch1tect/access-log-stats>
- Not designed to really be used
- Templates are not yet embedded within the binary, so if you build a binary you'll also need to carry a copy of the templates with you wherever you need to run it
- Error handling is bad/non existent

If you still want to give it a go:

1. Get your access log, from here on out: /path/to/access.log
1. Copy access-log-cli.toml-sample to access-log-cli.toml
1. Work out the regexes that you need for your access log format (the regex's I am using for caddy's console format are provided as an example - I hope to add more examples for different formats / web servers at a later date)
1. Run the report generation:
    ```
    If you are in the same directory as access-log-cli.toml:

    go run main.go -l /path/to/access.log -o report-output.html

    If you are not in the same dir as the configuration, or you have changed the name of your config file:

    go run main.go -l /path/to/access.log -u report-output.html -c /path/to/access-log-cli.toml
    ```