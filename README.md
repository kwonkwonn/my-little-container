# Micro Container

A lightweight container runtime that provides:

- **Resource Limiting** using **cgroups**  (working on)
- **Environment Isolation** using **namespaces**  (next step)
- **Virtual Networking** with **veth pairs** and **iptables**  (next step)
- **Minimal Runtime Environment** leveraging **overlayfs (if available)** and OS file system packages  (next step)

> ⚠️ Currently, this project only runs on **RHEL 10**.  
> To test it on other environments, you may need to adjust filesystem structures and related configurations.  
> And this is worked by sole experimental purpose.
> This could let you to know how basic's of container works under the hood 

If you are interested, feel free to **open an issue** or **submit a pull request**.
