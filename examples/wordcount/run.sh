#!/bin/bash

docker run -it \
  -v $(pwd):/jobs/usercode \
  verdverm/dstk-jobs
