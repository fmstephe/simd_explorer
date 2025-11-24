#!/bin/bash

find . -type d -name _generate -exec go generate {} \;
