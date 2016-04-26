VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  # ssh
  config.ssh.username = "cs"
  config.ssh.private_key_path = "~/.ssh/id_rsa"
  config.ssh.forward_agent = true

  # box
  config.vm.box = "nishida"
  config.vm.box_url = "~/Dropbox/Public/CentOS-7.2-x86_64-nishida-virtualbox.box"

  # network
  config.vm.network :private_network, ip: "172.16.99.99"
  # config.vm.network :public_network

  # synced folders
  config.vm.synced_folder ".", "/mnt/vagrant", owner: "nishida", group: "nishida"

  # spec
  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--cpus", "2"]
    vb.customize ["modifyvm", :id, "--memory", "2048"]
  end

end
