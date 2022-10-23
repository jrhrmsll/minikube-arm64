Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-20.04-arm64"

  config.vm.provider "parallels" do |prl|
    prl.linked_clone = false

    prl.name   = "minikube"
    prl.memory = 2048
    prl.cpus   = 2
  end

  config.vm.network "private_network", type: "dhcp"

  config.vm.provision "shell", path: "setup.sh"
end