require 'fileutils'
require 'open-uri'

EXTRA_HBASE_SERVERS = 2

Vagrant.configure("2") do |config|
  create_hbase_vm(1, config, 4096)

  (1..EXTRA_HBASE_SERVERS).each do |vm_number|
    create_hbase_vm(vm_number + 1, config, 1024)
  end
end

def create_hbase_vm(vm_number, config, memory)
  vm_name = "hbase-%02d" % vm_number

  config.vm.define vm_name do |config|
    config.vm.box = "ubuntu/precise64"
    config.vm.hostname = vm_name
    config.vm.provider :virtualbox do |vb|
      vb.gui = false
      vb.memory = memory
      vb.cpus = 1
    end

    vm_ip = "172.17.8.#{vm_number+100}"
    config.vm.network :private_network, ip: vm_ip

  end
end
