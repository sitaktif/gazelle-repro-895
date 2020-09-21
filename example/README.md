To reproduce the issue, run:

    cd example && bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=gazelle.bzl%go_deps_local -prune=true

You will notice that the root gazelle.bzl does not contain any dependencies anymore.
