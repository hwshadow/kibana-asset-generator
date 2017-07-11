#!/bin/sh
curl -sk -XDELETE http://elastic-kag:9200/testdata.dnd/
curl -sk http://elastic-kag:9200/_template/testdata_dnd -d @./template
while read json; do curl -sk http://elastic-kag:9200/testdata.dnd/testdata/ -d "$json"; done < ./data
