# Design
System configuration design.

# Table Of Contents
- [Overview](#overview)
- [Operating Systems](#operating-systems)
- [Automatic Documentation](#automatic-documentation)
- [Modules](#modules)
- [Actions](#actions)
- [Configuration](#configuration)

# Overview
The tool defines 4 primitives which are used to configure a system:

- Module Groups
- Modules
- Actions
- Configuration

Different parts of the system are configured by *modules*. Similar modules are organized into *module groups*. Modules 
use *actions* to change the state of the system. Modules also accept *configuration*.  

Each of these primitives provide automatic documentation. When combined the documentation from these primitives create a
documentation page detailing the steps used to set up a system.

# Operating Systems
[Arch Linux](http://archlinux.org/) is the only linux distribution supported at this time.  

The tool has the concept of actions which allow support for other operating systems to be added.

# Automatic Documentation
The tool automatically generates a documentation site detailing how the system will be set up.  

Module groups will be used to organize related modules in a sidebar. Modules will each have their own page. On this 
page a short description of what the module does is present. Along with a description of what each action the module 
calls does.

# Modules
A module has 2 steps:

- Install
    - Configure relevant part of system using actions
- Verify
	- Ensure the expected changes were made by the module

# Actions
Actions modify the state of the system. They abstract how the system's state is actually modified. Allowing modules to 
use actions and let the implementation of the action run different commands depending on the operating system.

# Configuration
Configuration is used by modules to modify their behavior.
