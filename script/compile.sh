#!/usr/bin/env bash
#

set -axe

for F in $(ls build/*.tex); do
    pdflatex "${F}"
    pdflatex "${F}"
done
