# find-up

Stupid little dir/file "finder" for going up a directory tree to find things like the nearest ".git" dir, "Dockerfile", etc.

Say you have a docker project in `/foo`. Sometimes you're in `/foo/bar`, sometimes `/foo/baz`,
sometimes `/foo/bar/baz`, etc. You need your bash script to find the "nearest" `Dockerfile`.
Just run `find-up Dockerfile` and use the path it outputs! Easy! And probably overkill!
