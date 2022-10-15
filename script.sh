#!/bin/sh


# https://stackoverflow.com/questions/34171568/return-value-from-python-script-to-shell-script
echo HEEEERRE
result=`python feature/main_feature.py`
echo SHELL $result 


