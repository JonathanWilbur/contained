# Contained

A tool for telling if you are running in a container. Think `whoami`, but it
tells you whether or not it is executing in a container.

## Usage

`contained`

### Return codes

- `0`: Running in a container.
- `1`: Not running in a container.
- `2`: Error encountered.

## To Do

- [ ] Check for these container runtimes
  - [ ] Docker
    - [x] Linux containers
    - [ ] Windows containers
  - [x] rkt
  - [ ] Singularity
  - [ ] Nanobox
  - [ ] chroot Jail
  - [ ] vkernel
  - [ ] sysjail
  - [ ] lmctfy
  - [x] LXC
  - [ ] Solaris Zones
  - [ ] Systemd Container
- [ ] Add command-line options:
  - [ ] `--quiet`, `-q`: Do not produce output.
  - [ ] `--exit-successfully`, `-x`: Always exit with a zero-exit code.