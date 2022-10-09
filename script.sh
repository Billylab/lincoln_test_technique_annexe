#!/bin/sh
echo Set GOOGLE APPLICATION CREDENTIALS
export GOOGLE_APPLICATION_CREDENTIALS=/secrets/sa-run-exercice-lincoln

echo Starting Python script : feature
python feature/main_feature.py
