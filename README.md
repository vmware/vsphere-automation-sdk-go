# vsphere-automation-sdk-golang

vSphere Automation SDK for GO Lang.

# VMware vSphere Automation SDK for GO Lang
## Table of Contents
- [Abstract](#abstract)
- [Supported vCenter Releases](#supported-vcenter-releases)
- [Table of Contents](#table-of-contents)
- [Quick Start Guide](#quick-start-guide)
- [API Documentation](#api-documentation)
- [Submitting samples](#submitting-samples)
- [Resource Maintenance](#resource-maintenance)
- [Repository Administrator Resources](#repository-administrator-resources)
  - [Board Members](#board-members)
  - [Approval of Additions](#approval-of-additions)
- [VMware Resources](#vmware-resources)

## Abstract
This document describes the vSphere Automation GOLang SDK samples that use the vSphere Automation
GOLang client library. Additionally, some of the samples demonstrate the combined use of the
vSphere Automation and vSphere Web Service APIs.

## Supported OnPrem vCenter Releases:
Please refer to the notes in each sample for detailed compatibility information. 

## VMware Cloud on AWS Support:

## Quick Start Guide
This document will walk you through getting up and running with the GOLang SDK Samples. Prior to running the samples you will need to setup a vCenter test environment, the following steps will take you through this process.
Before you can run the SDK samples we'll need to walk you through the following steps:

### Setting up a vSphere Test Environment
**NOTE:** The samples are intended to be run against a freshly installed **non-Production** vSphere setup as the scripts may make changes to the test environment and in some cases can destroy items when needed.

To run the samples a vSphere test environment is required with the following minimum configuration
* 1 vCenter Server
* 2 ESX hosts
* 1 NFS Datastore with at least 3GB of free capacity

Apart from the above, each individual sample may require additional setup. Please refer to the sample parameters for more information on that.

### Building the Samples
In the root directory of your folder after cloning the repository, run the below maven commands -

### Running the Samples

## API Documentation
The API documentation can be accessed from here: 

[vSphere APIs]().

## Submitting samples

### Developer Certificate of Origin

Before you start working with this project, please read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on as an open-source patch.

### Required Information
The following information must be included in the README.md for the sample.
* Author Name
  * This can include full name, email address or other identifiable piece of information that would allow interested parties to contact author with questions.
* Date
  * Date the sample was originally written
* Minimal/High Level Description
  * What does the sample do ?
* Any KNOWN limitations or dependencies

### Suggested Information
The following information should be included when possible. Inclusion of information provides valuable information to consumers of the resource.
* vSphere version against which the sample was developed/tested
* SDK version against which the sample was developed/tested

### Contribution Process

* Follow the [GitHub process](https://help.github.com/articles/fork-a-repo)
  * Please use one branch per sample or change-set
  * Please use one commit and pull request per sample
  * Please post the sample output along with the pull request
  * If you include a license with your sample, use the project license

### Code Style


## Resource Maintenance
### Maintenance Ownership
Ownership of any and all submitted samples are maintained by the submitter.
### Filing Issues
Any bugs or other issues should be filed within GitHub by way of the repository’s Issue Tracker.
### Resolving Issues
Any community member can resolve issues within the repository, however only the board member can approve the update. Once approved, assuming the resolution involves a pull request, only a board member will be able to merge and close the request.

### VMware Sample Exchange
It is highly recommended to add any and all submitted samples to the VMware Sample Exchange:  <https://code.vmware.com/samples>

Sample Exchange can be allowed to access your GitHub resources, by way of a linking process, where they can be indexed and searched by the community. There are VMware social media accounts which will advertise resources posted to the site and there's no additional accounts needed, as the VMware Sample Exchange uses MyVMware credentials.     

## Repository Administrator Resources
### Board Members

Board members are volunteers from the SDK community and VMware staff members, board members are not held responsible for any issues which may occur from running of samples from this repository.

Members:
* (VMware)

### Approval of Additions
Items added to the repository, including items from the Board members, require 2 votes from the board members before being added to the repository. The approving members will have ideally downloaded and tested the item. When two “Approved for Merge” comments are added from board members, the pull can then be committed to the repository.

## VMware Resources

* [vSphere Automation SDK Overview](http://pubs.vmware.com/vsphere-65/index.jsp#com.vmware.vapi.progguide.doc/GUID-AF73991C-FC1C-47DF-8362-184B6544CFDE.html)
* [VMware Code](https://code.vmware.com/home)
* [VMware Developer Community](https://communities.vmware.com/community/vmtn/developer)
* [VMware vSphere GOLang API Reference documentation]().
* [VMware GOLang forum]()

## Contributing

The vsphere-automation-sdk-go project team welcomes contributions from the community. Before you start working with vsphere-automation-sdk-go, please
read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be
signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on
as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License
vsphere-automation-sdk-go is available under the [BSD-2 license](LICENSE.txt).
