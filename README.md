https://disintegration.dev/

jq .[0].images[].url | sed -e 's/^"//' -e 's/"$//' | sort | uniq
