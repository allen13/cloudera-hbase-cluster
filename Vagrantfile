Vagrant.configure("2") do |config|
  config.vm.box = "chef/centos-7.0"

  (1..1).each do |vm_number|
    provision_vm(vm_number, config)
  end
end

def provision_vm(vm_number, config)
  vm_name = "node-%02d" % vm_number

  config.vm.define vm_name do |config|
    config.vm.hostname = vm_name

    config.vm.provider :virtualbox do |vb|
      vb.gui = false
      vb.memory = 1024
      vb.cpus = 1
    end

    config.vm.network :private_network, ip: "172.17.8.#{vm_number+100}"

    config.vm.provision "ansible" do |ansible|
      ansible.playbook = "docker-playbook/site.yml"
    end

  end
end
