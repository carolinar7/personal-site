## TODOs:
- [ ] Add social media [meta tags.](https://css-tricks.com/essential-meta-tags-social-media/)
- [ ] Possibly update font used throughout page.

## Docker + GCP Cloud Run
Played around with containarizing my website + creating a server for it :)
### Commands to run
`docker build -t personal-site .` and then `docker run -p 8080:8080 personal-site`
### Helpfule commands
- Command for viewing generated directory: `docker run -it personal-site /bin/sh`