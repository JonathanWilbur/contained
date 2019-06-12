# How to Check For Containerization

This utility in `systemd`,
[`systemd-detect-virt`](https://github.com/systemd/systemd/blob/03b35f8775c3a20ad05ff1cbac74e15cc24d8990/src/detect-virt/detect-virt.c),
can be used as a reference in checking for containerization.
[This file](https://github.com/systemd/systemd/blob/6c8a2c679313f8283514923daf65f5e9d050d94c/src/basic/virt.c)
contains the actual functions used in checking for containerization.

## Docker

Check `/proc/1/cgroup` for `/docker/`.

- [1](https://stackoverflow.com/questions/23513045/how-to-check-if-a-process-is-running-inside-docker-container)

## rkt



## Singularity



## Nanobox



## chroot Jail

Check that the inode for `/` is always 2. If it is not, you are in a chroot jail.

Also, in `virt.c`:

```c
int running_in_chroot(void) {
    int r;
    if (getenv_bool("SYSTEMD_IGNORE_CHROOT") > 0)
            return 0;
    r = files_same("/proc/1/root", "/", 0);
    if (r < 0)
            return r;
    return r == 0;
}
```

```c
int files_same(const char *filea, const char *fileb, int flags) {
        struct stat a, b;
        assert(filea);
        assert(fileb);
        if (fstatat(AT_FDCWD, filea, &a, flags) < 0)
                return -errno;
        if (fstatat(AT_FDCWD, fileb, &b, flags) < 0)
                return -errno;
        return a.st_dev == b.st_dev &&
                a.st_ino == b.st_ino;
}
```

As you can see, it looks like you just check if the device and inode is the
same for `/proc/1/root` and `/`. If they are the same, you are _not_ running
in a container.

- [1](https://stackoverflow.com/questions/75182/detecting-a-chroot-jail-from-within)

## vkernel



## sysjail



## lmctfy



## LXC

Check `/proc/1/cgroup` for `/lxc/`.

Sources:

- [1](https://stackoverflow.com/questions/20010199/how-to-determine-if-a-process-runs-inside-lxc-docker)

## Solaris Zones



## Systemd Container


