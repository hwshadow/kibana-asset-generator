#!/bin/sh
cd "/project/testdata/dnd/"
curl -sk -XDELETE http://elastic-kag:9200/testdata.dnd/
curl -sk -XDELETE http://elastic-kag:9200/.kibana/
#curl -sk http://elastic-kag:9200/_template/testdata_dnd -d @./mapping_template2
curl -sk http://elastic-kag:9200/_template/testdata_dnd -d @./mapping_template5
while read json; do curl -sk http://elastic-kag:9200/testdata.dnd/testdata/ -d "$json"; done < ./documents
../../bin/kag -template="./kag_template/" -write=true -dashTitle="dash" -pre="testdata-dnd-" -idx="testdata.dnd" -timeField="dob"
curl -sk http://elastic-kag:920/_cat/indices
