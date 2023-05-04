# reporter

Send custom data to app.segment.com

# segment sdk
[https://segment.com/docs/connections/sources/catalog/libraries/server/go/](https://segment.com/docs/connections/sources/catalog/libraries/server/go/)



# install
go install github.com/bitmyth/reporter@latest

# run
nohup sh -c 'for ((;;));do reporter -k work-mini && sleep 10;done '&
