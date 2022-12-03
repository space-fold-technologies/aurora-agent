```
  ______  __    __ _______   ______  _______   ______  
 /      \|  \  |  \       \ /      \|       \ /      \ 
|  ▓▓▓▓▓▓\ ▓▓  | ▓▓ ▓▓▓▓▓▓▓\  ▓▓▓▓▓▓\ ▓▓▓▓▓▓▓\  ▓▓▓▓▓▓\
| ▓▓__| ▓▓ ▓▓  | ▓▓ ▓▓__| ▓▓ ▓▓  | ▓▓ ▓▓__| ▓▓ ▓▓__| ▓▓
| ▓▓    ▓▓ ▓▓  | ▓▓ ▓▓    ▓▓ ▓▓  | ▓▓ ▓▓    ▓▓ ▓▓    ▓▓
| ▓▓▓▓▓▓▓▓ ▓▓  | ▓▓ ▓▓▓▓▓▓▓\ ▓▓  | ▓▓ ▓▓▓▓▓▓▓\ ▓▓▓▓▓▓▓▓
| ▓▓  | ▓▓ ▓▓__/ ▓▓ ▓▓  | ▓▓ ▓▓__/ ▓▓ ▓▓  | ▓▓ ▓▓  | ▓▓
| ▓▓  | ▓▓\▓▓    ▓▓ ▓▓  | ▓▓\▓▓    ▓▓ ▓▓  | ▓▓ ▓▓  | ▓▓
 \▓▓   \▓▓ \▓▓▓▓▓▓ \▓▓   \▓▓ \▓▓▓▓▓▓ \▓▓   \▓▓\▓▓   \▓▓
 Node Agent
```

# Aurora Node Agent #

This daemon allows makes the node it is installed on a part of the `cluster`. \
As a result, the node becomes a compute resource that can be used to run containers \
at the request of the `aurora-daemon` running on the master node. \

## Installation and Setup ##

### Pre-requisites ###

- Linux system or WSL-2 with golang 1.19.X to compile the code
- Debian Linux or any linux environment with docker and docker swarm installed and systemctl
- Open internal network interface protected from all external communication, this is usually available as an option for most `Virtual Private Server` providers like [Digital Ocean](https://www.digitalocean.com) and [Vultr](https://www.vultr.com)

#### STEPS ####

Edit the contents of in `resources/settings.yml` substituting what is required

```yaml
provider: DOCKER-SWARM
profile-directory: /etc/aurora # Path on file system to good stuff
listen-address: 0.0.0.0
advertise-address: ${SECURED-IP-ADDRESS-HERE}
host: 0.0.0.0
port: 2700
```

On a computer with the desired golang compiler for the target architecture running linux , you can compile the golang source code by running `./full_build.sh`
This might be x86_64, RISC-V, or ARM-64.

This will result into a tar-ball with the target architecture for your servers that you can transfer using

```bash
    scp aurora-agent.tar.gz hostname@IP-ADDRESS:~
    tar -xvf aurora-agent.tar.gz 
    sudo chmod +x aurora-agent  
    sudo mv aurora-agent /usr/local/bin/
```

Create a new group and user

```bash
   sudo groupadd --systemctl aurora
   sudo useradd -s /sbin/nologin --system -g aurora aurora
```

Create a systemd file

```bash
    sudo nano /etc/systemd/system/aurora-agent.service
```

Add the following lines

```bash
    [Unit]
    Description=Aurora Service [Daemon]
    After=network-online.target
    Wants=network-online.target
    [Service]
    Type=simple
    User=aurora
    Group=aurora
    ExecStart=/usr/local/bin/aurora-agent --mode=run
    ExecReload=/bin/kill -HUP $MAINPID
    KillSignal=SIGINT
    TimeoutStopSec=5
    Restart=on-failure
    SyslogIdentifier=aurora-agent
    [Install]
    WantedBy=multi-user.target 
```

Save and close the file when you are finished

Next, reload the system daemon with the following command

```bash
    sudo systemctl daemon-reload
    sudo systemctl start aurora-agent
    sudo systemctl enable aurora-agent
```

And check to see if the daemon is running

```bash
    sudo systemctl status aurora-agent
```

Then the last thing is to add aurora user to the docker user group

```bash
    sudo usermod -aG docker aurora
```

Using the `Aurora CLI client` you can register the node with the following command

```bash
    aurora-client nodes:create ${NODE_NAME} --cluster=${CLUSTER_NAME} --address=${IP-ADDRESS} --provider=${PROVIDER} --description=${DESCRIPTION}
```

#### NB ####

The default cluster name is `default` and the provider currently supported is `DOCKER`. \
You can check if the setup went ok by running the command:

```bash
   aurora-client nodes:list --cluster=${CLUSTER_NAME}
```

Everything should be running ok now

### What's Next ? ###

To use the `PaaS` you can check out [Aurora CLI](https://github.com/space-fold-technologies/aurora-client) which can be used on your developer work station. \

### Features ###
  
- [x] Docker Swarm Integration
- [ ] Kubernetes Integration
- [ ] Podman Integration
- [ ] Agent Node Stats
- [ ] Scheduled container checks
