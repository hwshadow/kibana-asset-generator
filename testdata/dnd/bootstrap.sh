#!/bin/sh
cd "/project/testdata/dnd/"
printf "\n\nZeroing testdata.dnd index\n"
curl -sk -XDELETE http://elastic-kag:9200/testdata.dnd/
printf "\n\nZeroing .kibana index\n"
curl -sk -XDELETE http://elastic-kag:9200/.kibana/
ESINFO=$(curl -sk http://elastic-kag:9200/?filter_path=version.number)
printf "\n\nInjecting testdata.dnd mappings\n"
if echo $ESINFO | grep -q "2."; then
  printf "Assuming ES 2.x\n"
  curl -sk http://elastic-kag:9200/_template/testdata_dnd -d @./mapping_template2
elif echo $ESINFO | grep -q "5."; then
  printf "Assuming ES 5.x\n"
  curl -sk http://elastic-kag:9200/_template/testdata_dnd -d @./mapping_template5
fi
printf "\n\nLoading dnd dataset\n"
while read json; do curl -sk http://elastic-kag:9200/testdata.dnd/testdata/ -d "$json"; done < ./documents
printf "\n\nExecuting KAG with dnd templates\n"
../../bin/kag -template="./kag_template/" -write=true -dashTitle="dash" -pre="testdata-dnd-" -idx="testdata.dnd" -timeField="dob"
curl -sk http://elastic-kag:920/_cat/indices
