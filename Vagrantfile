require 'fileutils'
require 'open-uri'

NODES = 4

Vagrant.configure("2") do |config|
  config.vm.box = "chef/centos-7.0"

  (1..NODES).each do |vm_number|
    create_vm(vm_number, config)
  end
end

def create_vm(vm_number, config)
  vm_name = "node-%02d" % vm_number

  config.vm.define vm_name do |config|
    config.vm.hostname = vm_name

    config.vm.provider :virtualbox do |vb|
      vb.gui = false
      vb.memory = 2048
      vb.cpus = 1
    end

    vm_ip = "172.17.8.#{vm_number+100}"
    config.vm.network :private_network, ip: vm_ip

  end
end
