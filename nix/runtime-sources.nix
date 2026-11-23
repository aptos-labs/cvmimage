let
  nvidiaCompute = {
    name = "libnvidia-compute";
    url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/libnvidia-compute_595.112.07-1ubuntu1_amd64.deb";
    sha256 = "843a6adfea7bc16f2aad3d3d2daca931fd67cb57cf32e2622b17b0eba914bec2";
  };
in
{
  busybox = {
    name = "busybox-static";
    url = "https://snapshot.ubuntu.com/ubuntu/20261123T000000Z/pool/main/b/busybox/busybox-static_1.37.0-7ubuntu1_amd64.deb";
    sha256 = "fd605342f62268753076aa7d9321ff098b36ba53e47434145a9aedc28fd141a4";
  };

  docker = {
    name = "docker-29.10.0-static";
    url = "https://download.docker.com/linux/static/stable/x86_64/docker-29.10.0.tgz";
    sha256 = "sha256-fB663GXHFs3G01/Hh1ntDkMX3r9Dwh3RLwZvIAaZne8=";
  };

  # Debug-image only. runsc must sit next to gvisor-bin/; later releases
  # no longer bundle sidecars inside the runsc binary.
  gvisor = {
    name = "gvisor-20261116.0-x86_64";
    url = "https://github.com/google/gvisor/releases/download/release-20261116.0/gvisor-x86_64.tar.bz2";
    sha256 = "sha256-4iY4k5RzdZqGAugPK5aptn5K/N9Y8sUdtSsajpAx6kU=";
  };

  nvidiaDebs = [
    { name = "libnvidia-cfg1"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/libnvidia-cfg1_595.112.07-1ubuntu1_amd64.deb"; sha256 = "235e16462eccec3cf55ecbf1d786d73c67d28fe76a42728ad5cb79111d466e51"; }
    nvidiaCompute
    { name = "libnvidia-gpucomp"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/libnvidia-gpucomp_595.112.07-1ubuntu1_amd64.deb"; sha256 = "acfdf29ce30452e2993933587b91703d3fb4e2642dbd171aeafeaf27b36048e2"; }
    { name = "libnvidia-nscq"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/libnvidia-nscq_595.112.07-1ubuntu1_amd64.deb"; sha256 = "a0d5949485da985de5734499f82289b6464597190cbbb49eb5c217d4029b5f5d"; }
    { name = "nvidia-container-toolkit-base"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/nvidia-container-toolkit-base_1.20.2-1_amd64.deb"; sha256 = "479fec58540595e9873839c9518373b26a54b1fb51b0153a261fa29357ea9fee"; }
    { name = "nvidia-fabricmanager"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/nvidia-fabricmanager_595.112.07-1ubuntu1_amd64.deb"; sha256 = "a764c61951f45e2973359cc23d8d0818f460808fae42e22567455a6cf6b9cc08"; }
    { name = "nvidia-firmware"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/nvidia-firmware_595.112.07-1ubuntu1_amd64.deb"; sha256 = "4e0519908df12faf6d293138bb65877d0067dbb529096e7f6965e8892a3756b7"; }
    { name = "nvidia-persistenced"; url = "https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2604/x86_64/nvidia-persistenced_595.112.07-1ubuntu1_amd64.deb"; sha256 = "23081a95cb67026202fbbb4d019252f7a81e120a3bd3b6932ae77a556f0d51a2"; }
  ];

  inherit nvidiaCompute;
}
