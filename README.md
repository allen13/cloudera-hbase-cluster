opentsdb-cluster
=======================

Provisioned opentsdb cluster

####Setup

Requires [Vagrant](https://docs.vagrantup.com/v2/installation/) and [Ansible](http://docs.ansible.com/intro_installation.html)

Install ansible on Ubuntu

    $ sudo apt-get install software-properties-common
    $ sudo apt-add-repository ppa:ansible/ansible
    $ sudo apt-get update
    $ sudo apt-get install ansible

Install ansible on Mac

    $ brew install ansible

####Running

    vagrant up
    bin/export_vagrant_ssh_keys
    bin/provision


####Destroying

    bin/destroy
