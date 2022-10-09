#!/bin/sh
echo Set GOOGLE APPLICATION CREDENTIALS
# export GOOGLE_APPLICATION_CREDENTIALS=/secrets/sa-run-exercice-lincoln
export GOOGLE_APPLICATION_CREDENTIALS=test01-lincoln-project-7c73284b4cd1.json

echo Starting Python script : feature

result=`python feature/main_feature.py`

echo $result


