# Go class: Build example
Docker build example (using sort animation)

You can find the contain on Docker Hub under [matthol2/sort-anim](https://hub.docker.com/repository/docker/matthol2/sort-anim). (I may rename that account to matt4biz to match these repos; if you don't see it, search on "sort-anim".)

## Sort animation

Routes:

| Route    | Description |
| -------- | ----------- |
| insert   | insertion sort |
| qsort    | Quicksort, pivot on high element  |
| qsortm   | Quicksort, pivot on middle element |
| qsort3   | Quicksort, pivot on median-of-3 element  |
| qsorti   | Quicksort, pivot on median-of-3 element; use insertion sort on small array |
| qsortf   | Quicksort, Dutch flag (3-way) partition |
| version  | show the version |

Each sort algorithm route takes two optional parameters:

| Parameter| Description |
| -------- | ----------- |
| loop     | animation loop, default 1 (use 0 to suppress looping) |
| delay    | delay between frames (ms), default 8 |

## Building

The Makefile has the following targets:

| Target    | Description |
| --------  | ----------- |
| sort      | build the program (default target) |
| lint      | run golangci-lint |
| committed | verify the repo is not dirty (doesn't verify it's pushed/tagged) |
| docker    | make the docker container |
| publish   | push the docker container; must be committed |

