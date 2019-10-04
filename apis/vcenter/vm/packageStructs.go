/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.vm.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm



// The ``GuestOS`` enumeration class defines the valid guest operating system types used for configuring a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type GuestOS string

const (
    // MS-DOS.
     GuestOS_DOS GuestOS = "DOS"
    // Windows 3.1
     GuestOS_WIN_31 GuestOS = "WIN_31"
    // Windows 95
     GuestOS_WIN_95 GuestOS = "WIN_95"
    // Windows 98
     GuestOS_WIN_98 GuestOS = "WIN_98"
    // Windows Millennium Edition
     GuestOS_WIN_ME GuestOS = "WIN_ME"
    // Windows NT 4
     GuestOS_WIN_NT GuestOS = "WIN_NT"
    // Windows 2000 Professional
     GuestOS_WIN_2000_PRO GuestOS = "WIN_2000_PRO"
    // Windows 2000 Server
     GuestOS_WIN_2000_SERV GuestOS = "WIN_2000_SERV"
    // Windows 2000 Advanced Server
     GuestOS_WIN_2000_ADV_SERV GuestOS = "WIN_2000_ADV_SERV"
    // Windows XP Home Edition
     GuestOS_WIN_XP_HOME GuestOS = "WIN_XP_HOME"
    // Windows XP Professional
     GuestOS_WIN_XP_PRO GuestOS = "WIN_XP_PRO"
    // Windows XP Professional Edition (64 bit)
     GuestOS_WIN_XP_PRO_64 GuestOS = "WIN_XP_PRO_64"
    // Windows Server 2003, Web Edition
     GuestOS_WIN_NET_WEB GuestOS = "WIN_NET_WEB"
    // Windows Server 2003, Standard Edition
     GuestOS_WIN_NET_STANDARD GuestOS = "WIN_NET_STANDARD"
    // Windows Server 2003, Enterprise Edition
     GuestOS_WIN_NET_ENTERPRISE GuestOS = "WIN_NET_ENTERPRISE"
    // Windows Server 2003, Datacenter Edition
     GuestOS_WIN_NET_DATACENTER GuestOS = "WIN_NET_DATACENTER"
    // Windows Small Business Server 2003
     GuestOS_WIN_NET_BUSINESS GuestOS = "WIN_NET_BUSINESS"
    // Windows Server 2003, Standard Edition (64 bit)
     GuestOS_WIN_NET_STANDARD_64 GuestOS = "WIN_NET_STANDARD_64"
    // Windows Server 2003, Enterprise Edition (64 bit)
     GuestOS_WIN_NET_ENTERPRISE_64 GuestOS = "WIN_NET_ENTERPRISE_64"
    // Windows Longhorn (experimental)
     GuestOS_WIN_LONGHORN GuestOS = "WIN_LONGHORN"
    // Windows Longhorn (64 bit) (experimental)
     GuestOS_WIN_LONGHORN_64 GuestOS = "WIN_LONGHORN_64"
    // Windows Server 2003, Datacenter Edition (64 bit) (experimental)
     GuestOS_WIN_NET_DATACENTER_64 GuestOS = "WIN_NET_DATACENTER_64"
    // Windows Vista
     GuestOS_WIN_VISTA GuestOS = "WIN_VISTA"
    // Windows Vista (64 bit)
     GuestOS_WIN_VISTA_64 GuestOS = "WIN_VISTA_64"
    // Windows 7
     GuestOS_WINDOWS_7 GuestOS = "WINDOWS_7"
    // Windows 7 (64 bit)
     GuestOS_WINDOWS_7_64 GuestOS = "WINDOWS_7_64"
    // Windows Server 2008 R2 (64 bit)
     GuestOS_WINDOWS_7_SERVER_64 GuestOS = "WINDOWS_7_SERVER_64"
    // Windows 8
     GuestOS_WINDOWS_8 GuestOS = "WINDOWS_8"
    // Windows 8 (64 bit)
     GuestOS_WINDOWS_8_64 GuestOS = "WINDOWS_8_64"
    // Windows 8 Server (64 bit)
     GuestOS_WINDOWS_8_SERVER_64 GuestOS = "WINDOWS_8_SERVER_64"
    // Windows 10
     GuestOS_WINDOWS_9 GuestOS = "WINDOWS_9"
    // Windows 10 (64 bit)
     GuestOS_WINDOWS_9_64 GuestOS = "WINDOWS_9_64"
    // Windows 10 Server (64 bit)
     GuestOS_WINDOWS_9_SERVER_64 GuestOS = "WINDOWS_9_SERVER_64"
    // Windows Hyper-V
     GuestOS_WINDOWS_HYPERV GuestOS = "WINDOWS_HYPERV"
    // Windows Server 2019. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     GuestOS_WINDOWS_SERVER_2019 GuestOS = "WINDOWS_SERVER_2019"
    // FreeBSD 10 or earlier
     GuestOS_FREEBSD GuestOS = "FREEBSD"
    // FreeBSD 10 x64 or earlier
     GuestOS_FREEBSD_64 GuestOS = "FREEBSD_64"
    // FreeBSD 11
     GuestOS_FREEBSD_11 GuestOS = "FREEBSD_11"
    // FreeBSD 12 or later
     GuestOS_FREEBSD_12 GuestOS = "FREEBSD_12"
    // FreeBSD 11 x64
     GuestOS_FREEBSD_11_64 GuestOS = "FREEBSD_11_64"
    // FreeBSD 12 x64 or later
     GuestOS_FREEBSD_12_64 GuestOS = "FREEBSD_12_64"
    // Red Hat Linux 2.1
     GuestOS_REDHAT GuestOS = "REDHAT"
    // Red Hat Enterprise Linux 2
     GuestOS_RHEL_2 GuestOS = "RHEL_2"
    // Red Hat Enterprise Linux 3
     GuestOS_RHEL_3 GuestOS = "RHEL_3"
    // Red Hat Enterprise Linux 3 (64 bit)
     GuestOS_RHEL_3_64 GuestOS = "RHEL_3_64"
    // Red Hat Enterprise Linux 4
     GuestOS_RHEL_4 GuestOS = "RHEL_4"
    // Red Hat Enterprise Linux 4 (64 bit)
     GuestOS_RHEL_4_64 GuestOS = "RHEL_4_64"
    // Red Hat Enterprise Linux 5
     GuestOS_RHEL_5 GuestOS = "RHEL_5"
    // Red Hat Enterprise Linux 5 (64 bit) (experimental)
     GuestOS_RHEL_5_64 GuestOS = "RHEL_5_64"
    // Red Hat Enterprise Linux 6
     GuestOS_RHEL_6 GuestOS = "RHEL_6"
    // Red Hat Enterprise Linux 6 (64 bit)
     GuestOS_RHEL_6_64 GuestOS = "RHEL_6_64"
    // Red Hat Enterprise Linux 7
     GuestOS_RHEL_7 GuestOS = "RHEL_7"
    // Red Hat Enterprise Linux 7 (64 bit)
     GuestOS_RHEL_7_64 GuestOS = "RHEL_7_64"
    // Red Hat Enterprise Linux 8 (64 bit)
     GuestOS_RHEL_8_64 GuestOS = "RHEL_8_64"
    // CentOS 4/5
     GuestOS_CENTOS GuestOS = "CENTOS"
    // CentOS 4/5 (64-bit)
     GuestOS_CENTOS_64 GuestOS = "CENTOS_64"
    // CentOS 6
     GuestOS_CENTOS_6 GuestOS = "CENTOS_6"
    // CentOS 6 (64-bit)
     GuestOS_CENTOS_6_64 GuestOS = "CENTOS_6_64"
    // CentOS 7
     GuestOS_CENTOS_7 GuestOS = "CENTOS_7"
    // CentOS 7 (64-bit)
     GuestOS_CENTOS_7_64 GuestOS = "CENTOS_7_64"
    // CentOS 8 (64-bit)
     GuestOS_CENTOS_8_64 GuestOS = "CENTOS_8_64"
    // Oracle Linux 4/5
     GuestOS_ORACLE_LINUX GuestOS = "ORACLE_LINUX"
    // Oracle Linux 4/5 (64-bit)
     GuestOS_ORACLE_LINUX_64 GuestOS = "ORACLE_LINUX_64"
    // Oracle Linux 6
     GuestOS_ORACLE_LINUX_6 GuestOS = "ORACLE_LINUX_6"
    // Oracle Linux 6 (64-bit)
     GuestOS_ORACLE_LINUX_6_64 GuestOS = "ORACLE_LINUX_6_64"
    // Oracle Linux 7
     GuestOS_ORACLE_LINUX_7 GuestOS = "ORACLE_LINUX_7"
    // Oracle Linux 7 (64-bit)
     GuestOS_ORACLE_LINUX_7_64 GuestOS = "ORACLE_LINUX_7_64"
    // Oracle Linux 8 (64-bit)
     GuestOS_ORACLE_LINUX_8_64 GuestOS = "ORACLE_LINUX_8_64"
    // Suse Linux
     GuestOS_SUSE GuestOS = "SUSE"
    // Suse Linux (64 bit)
     GuestOS_SUSE_64 GuestOS = "SUSE_64"
    // Suse Linux Enterprise Server 9
     GuestOS_SLES GuestOS = "SLES"
    // Suse Linux Enterprise Server 9 (64 bit)
     GuestOS_SLES_64 GuestOS = "SLES_64"
    // Suse linux Enterprise Server 10
     GuestOS_SLES_10 GuestOS = "SLES_10"
    // Suse Linux Enterprise Server 10 (64 bit) (experimental)
     GuestOS_SLES_10_64 GuestOS = "SLES_10_64"
    // Suse linux Enterprise Server 11
     GuestOS_SLES_11 GuestOS = "SLES_11"
    // Suse Linux Enterprise Server 11 (64 bit)
     GuestOS_SLES_11_64 GuestOS = "SLES_11_64"
    // Suse linux Enterprise Server 12
     GuestOS_SLES_12 GuestOS = "SLES_12"
    // Suse Linux Enterprise Server 12 (64 bit)
     GuestOS_SLES_12_64 GuestOS = "SLES_12_64"
    // Suse Linux Enterprise Server 15 (64 bit)
     GuestOS_SLES_15_64 GuestOS = "SLES_15_64"
    // Novell Linux Desktop 9
     GuestOS_NLD_9 GuestOS = "NLD_9"
    // Open Enterprise Server
     GuestOS_OES GuestOS = "OES"
    // Sun Java Desktop System
     GuestOS_SJDS GuestOS = "SJDS"
    // Mandrake Linux
     GuestOS_MANDRAKE GuestOS = "MANDRAKE"
    // Mandriva Linux
     GuestOS_MANDRIVA GuestOS = "MANDRIVA"
    // Mandriva Linux (64 bit)
     GuestOS_MANDRIVA_64 GuestOS = "MANDRIVA_64"
    // Turbolinux
     GuestOS_TURBO_LINUX GuestOS = "TURBO_LINUX"
    // Turbolinux (64 bit)
     GuestOS_TURBO_LINUX_64 GuestOS = "TURBO_LINUX_64"
    // Ubuntu Linux
     GuestOS_UBUNTU GuestOS = "UBUNTU"
    // Ubuntu Linux (64 bit)
     GuestOS_UBUNTU_64 GuestOS = "UBUNTU_64"
    // Debian GNU/Linux 4
     GuestOS_DEBIAN_4 GuestOS = "DEBIAN_4"
    // Debian GNU/Linux 4 (64 bit)
     GuestOS_DEBIAN_4_64 GuestOS = "DEBIAN_4_64"
    // Debian GNU/Linux 5
     GuestOS_DEBIAN_5 GuestOS = "DEBIAN_5"
    // Debian GNU/Linux 5 (64 bit)
     GuestOS_DEBIAN_5_64 GuestOS = "DEBIAN_5_64"
    // Debian GNU/Linux 6
     GuestOS_DEBIAN_6 GuestOS = "DEBIAN_6"
    // Debian GNU/Linux 6 (64 bit)
     GuestOS_DEBIAN_6_64 GuestOS = "DEBIAN_6_64"
    // Debian GNU/Linux 7
     GuestOS_DEBIAN_7 GuestOS = "DEBIAN_7"
    // Debian GNU/Linux 7 (64 bit)
     GuestOS_DEBIAN_7_64 GuestOS = "DEBIAN_7_64"
    // Debian GNU/Linux 8
     GuestOS_DEBIAN_8 GuestOS = "DEBIAN_8"
    // Debian GNU/Linux 8 (64 bit)
     GuestOS_DEBIAN_8_64 GuestOS = "DEBIAN_8_64"
    // Debian GNU/Linux 9
     GuestOS_DEBIAN_9 GuestOS = "DEBIAN_9"
    // Debian GNU/Linux 9 (64 bit)
     GuestOS_DEBIAN_9_64 GuestOS = "DEBIAN_9_64"
    // Debian GNU/Linux 10
     GuestOS_DEBIAN_10 GuestOS = "DEBIAN_10"
    // Debian GNU/Linux 10 (64 bit)
     GuestOS_DEBIAN_10_64 GuestOS = "DEBIAN_10_64"
    // Debian GNU/Linux 11. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     GuestOS_DEBIAN_11 GuestOS = "DEBIAN_11"
    // Debian GNU/Linux 11 (64 bit). **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     GuestOS_DEBIAN_11_64 GuestOS = "DEBIAN_11_64"
    // Asianux Server 3
     GuestOS_ASIANUX_3 GuestOS = "ASIANUX_3"
    // Asianux Server 3 (64 bit)
     GuestOS_ASIANUX_3_64 GuestOS = "ASIANUX_3_64"
    // Asianux Server 4
     GuestOS_ASIANUX_4 GuestOS = "ASIANUX_4"
    // Asianux Server 4 (64 bit)
     GuestOS_ASIANUX_4_64 GuestOS = "ASIANUX_4_64"
    // Asianux Server 5 (64 bit)
     GuestOS_ASIANUX_5_64 GuestOS = "ASIANUX_5_64"
    // Asianux Server 7 (64 bit)
     GuestOS_ASIANUX_7_64 GuestOS = "ASIANUX_7_64"
    // Asianux Server 8 (64 bit)
     GuestOS_ASIANUX_8_64 GuestOS = "ASIANUX_8_64"
    // OpenSUSE Linux
     GuestOS_OPENSUSE GuestOS = "OPENSUSE"
    // OpenSUSE Linux (64 bit)
     GuestOS_OPENSUSE_64 GuestOS = "OPENSUSE_64"
    // Fedora Linux
     GuestOS_FEDORA GuestOS = "FEDORA"
    // Fedora Linux (64 bit)
     GuestOS_FEDORA_64 GuestOS = "FEDORA_64"
    // CoreOS Linux (64 bit)
     GuestOS_COREOS_64 GuestOS = "COREOS_64"
    // VMware Photon (64 bit)
     GuestOS_VMWARE_PHOTON_64 GuestOS = "VMWARE_PHOTON_64"
    // Linux 2.4x Kernel
     GuestOS_OTHER_24X_LINUX GuestOS = "OTHER_24X_LINUX"
    // Linux 2.4x Kernel (64 bit) (experimental)
     GuestOS_OTHER_24X_LINUX_64 GuestOS = "OTHER_24X_LINUX_64"
    // Linux 2.6x Kernel
     GuestOS_OTHER_26X_LINUX GuestOS = "OTHER_26X_LINUX"
    // Linux 2.6x Kernel (64 bit) (experimental)
     GuestOS_OTHER_26X_LINUX_64 GuestOS = "OTHER_26X_LINUX_64"
    // Linux 3.x Kernel
     GuestOS_OTHER_3X_LINUX GuestOS = "OTHER_3X_LINUX"
    // Linux 3.x Kernel (64 bit)
     GuestOS_OTHER_3X_LINUX_64 GuestOS = "OTHER_3X_LINUX_64"
    // Linux 4.x Kernel
     GuestOS_OTHER_4X_LINUX GuestOS = "OTHER_4X_LINUX"
    // Linux 4.x Kernel (64 bit)
     GuestOS_OTHER_4X_LINUX_64 GuestOS = "OTHER_4X_LINUX_64"
    // Linux 2.2x Kernel
     GuestOS_OTHER_LINUX GuestOS = "OTHER_LINUX"
    // Other Linux
     GuestOS_GENERIC_LINUX GuestOS = "GENERIC_LINUX"
    // Linux (64 bit) (experimental)
     GuestOS_OTHER_LINUX_64 GuestOS = "OTHER_LINUX_64"
    // Solaris 6
     GuestOS_SOLARIS_6 GuestOS = "SOLARIS_6"
    // Solaris 7
     GuestOS_SOLARIS_7 GuestOS = "SOLARIS_7"
    // Solaris 8
     GuestOS_SOLARIS_8 GuestOS = "SOLARIS_8"
    // Solaris 9
     GuestOS_SOLARIS_9 GuestOS = "SOLARIS_9"
    // Solaris 10 (32 bit) (experimental)
     GuestOS_SOLARIS_10 GuestOS = "SOLARIS_10"
    // Solaris 10 (64 bit) (experimental)
     GuestOS_SOLARIS_10_64 GuestOS = "SOLARIS_10_64"
    // Solaris 11 (64 bit)
     GuestOS_SOLARIS_11_64 GuestOS = "SOLARIS_11_64"
    // OS/2
     GuestOS_OS2 GuestOS = "OS2"
    // eComStation 1.x
     GuestOS_ECOMSTATION GuestOS = "ECOMSTATION"
    // eComStation 2.0
     GuestOS_ECOMSTATION_2 GuestOS = "ECOMSTATION_2"
    // Novell NetWare 4
     GuestOS_NETWARE_4 GuestOS = "NETWARE_4"
    // Novell NetWare 5.1
     GuestOS_NETWARE_5 GuestOS = "NETWARE_5"
    // Novell NetWare 6.x
     GuestOS_NETWARE_6 GuestOS = "NETWARE_6"
    // SCO OpenServer 5
     GuestOS_OPENSERVER_5 GuestOS = "OPENSERVER_5"
    // SCO OpenServer 6
     GuestOS_OPENSERVER_6 GuestOS = "OPENSERVER_6"
    // SCO UnixWare 7
     GuestOS_UNIXWARE_7 GuestOS = "UNIXWARE_7"
    // Mac OS 10.5
     GuestOS_DARWIN GuestOS = "DARWIN"
    // Mac OS 10.5 (64 bit)
     GuestOS_DARWIN_64 GuestOS = "DARWIN_64"
    // Mac OS 10.6
     GuestOS_DARWIN_10 GuestOS = "DARWIN_10"
    // Mac OS 10.6 (64 bit)
     GuestOS_DARWIN_10_64 GuestOS = "DARWIN_10_64"
    // Mac OS 10.7
     GuestOS_DARWIN_11 GuestOS = "DARWIN_11"
    // Mac OS 10.7 (64 bit)
     GuestOS_DARWIN_11_64 GuestOS = "DARWIN_11_64"
    // Mac OS 10.8 (64 bit)
     GuestOS_DARWIN_12_64 GuestOS = "DARWIN_12_64"
    // Mac OS 10.9 (64 bit)
     GuestOS_DARWIN_13_64 GuestOS = "DARWIN_13_64"
    // Mac OS 10.10 (64 bit)
     GuestOS_DARWIN_14_64 GuestOS = "DARWIN_14_64"
    // Mac OS 10.11 (64 bit)
     GuestOS_DARWIN_15_64 GuestOS = "DARWIN_15_64"
    // Mac OS 10.12 (64 bit)
     GuestOS_DARWIN_16_64 GuestOS = "DARWIN_16_64"
    // Mac OS 10.13 (64 bit)
     GuestOS_DARWIN_17_64 GuestOS = "DARWIN_17_64"
    // Mac OS 10.14 (64 bit)
     GuestOS_DARWIN_18_64 GuestOS = "DARWIN_18_64"
    // Mac OS 10.15 (64 bit). **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     GuestOS_DARWIN_19_64 GuestOS = "DARWIN_19_64"
    // VMware ESX 4
     GuestOS_VMKERNEL GuestOS = "VMKERNEL"
    // VMware ESX 5
     GuestOS_VMKERNEL_5 GuestOS = "VMKERNEL_5"
    // VMware ESX 6
     GuestOS_VMKERNEL_6 GuestOS = "VMKERNEL_6"
    // VMware ESX 6.5
     GuestOS_VMKERNEL_65 GuestOS = "VMKERNEL_65"
    // VMware ESX 7. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     GuestOS_VMKERNEL_7 GuestOS = "VMKERNEL_7"
    // Amazon Linux 2 (64 bit)
     GuestOS_AMAZONLINUX2_64 GuestOS = "AMAZONLINUX2_64"
    // CRX Pod 1. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     GuestOS_CRXPOD_1 GuestOS = "CRXPOD_1"
    // Other Operating System
     GuestOS_OTHER GuestOS = "OTHER"
    // Other Operating System (64 bit) (experimental)
     GuestOS_OTHER_64 GuestOS = "OTHER_64"
)

func (g GuestOS) GuestOS() bool {
    switch g {
        case GuestOS_DOS:
            return true
        case GuestOS_WIN_31:
            return true
        case GuestOS_WIN_95:
            return true
        case GuestOS_WIN_98:
            return true
        case GuestOS_WIN_ME:
            return true
        case GuestOS_WIN_NT:
            return true
        case GuestOS_WIN_2000_PRO:
            return true
        case GuestOS_WIN_2000_SERV:
            return true
        case GuestOS_WIN_2000_ADV_SERV:
            return true
        case GuestOS_WIN_XP_HOME:
            return true
        case GuestOS_WIN_XP_PRO:
            return true
        case GuestOS_WIN_XP_PRO_64:
            return true
        case GuestOS_WIN_NET_WEB:
            return true
        case GuestOS_WIN_NET_STANDARD:
            return true
        case GuestOS_WIN_NET_ENTERPRISE:
            return true
        case GuestOS_WIN_NET_DATACENTER:
            return true
        case GuestOS_WIN_NET_BUSINESS:
            return true
        case GuestOS_WIN_NET_STANDARD_64:
            return true
        case GuestOS_WIN_NET_ENTERPRISE_64:
            return true
        case GuestOS_WIN_LONGHORN:
            return true
        case GuestOS_WIN_LONGHORN_64:
            return true
        case GuestOS_WIN_NET_DATACENTER_64:
            return true
        case GuestOS_WIN_VISTA:
            return true
        case GuestOS_WIN_VISTA_64:
            return true
        case GuestOS_WINDOWS_7:
            return true
        case GuestOS_WINDOWS_7_64:
            return true
        case GuestOS_WINDOWS_7_SERVER_64:
            return true
        case GuestOS_WINDOWS_8:
            return true
        case GuestOS_WINDOWS_8_64:
            return true
        case GuestOS_WINDOWS_8_SERVER_64:
            return true
        case GuestOS_WINDOWS_9:
            return true
        case GuestOS_WINDOWS_9_64:
            return true
        case GuestOS_WINDOWS_9_SERVER_64:
            return true
        case GuestOS_WINDOWS_HYPERV:
            return true
        case GuestOS_WINDOWS_SERVER_2019:
            return true
        case GuestOS_FREEBSD:
            return true
        case GuestOS_FREEBSD_64:
            return true
        case GuestOS_FREEBSD_11:
            return true
        case GuestOS_FREEBSD_12:
            return true
        case GuestOS_FREEBSD_11_64:
            return true
        case GuestOS_FREEBSD_12_64:
            return true
        case GuestOS_REDHAT:
            return true
        case GuestOS_RHEL_2:
            return true
        case GuestOS_RHEL_3:
            return true
        case GuestOS_RHEL_3_64:
            return true
        case GuestOS_RHEL_4:
            return true
        case GuestOS_RHEL_4_64:
            return true
        case GuestOS_RHEL_5:
            return true
        case GuestOS_RHEL_5_64:
            return true
        case GuestOS_RHEL_6:
            return true
        case GuestOS_RHEL_6_64:
            return true
        case GuestOS_RHEL_7:
            return true
        case GuestOS_RHEL_7_64:
            return true
        case GuestOS_RHEL_8_64:
            return true
        case GuestOS_CENTOS:
            return true
        case GuestOS_CENTOS_64:
            return true
        case GuestOS_CENTOS_6:
            return true
        case GuestOS_CENTOS_6_64:
            return true
        case GuestOS_CENTOS_7:
            return true
        case GuestOS_CENTOS_7_64:
            return true
        case GuestOS_CENTOS_8_64:
            return true
        case GuestOS_ORACLE_LINUX:
            return true
        case GuestOS_ORACLE_LINUX_64:
            return true
        case GuestOS_ORACLE_LINUX_6:
            return true
        case GuestOS_ORACLE_LINUX_6_64:
            return true
        case GuestOS_ORACLE_LINUX_7:
            return true
        case GuestOS_ORACLE_LINUX_7_64:
            return true
        case GuestOS_ORACLE_LINUX_8_64:
            return true
        case GuestOS_SUSE:
            return true
        case GuestOS_SUSE_64:
            return true
        case GuestOS_SLES:
            return true
        case GuestOS_SLES_64:
            return true
        case GuestOS_SLES_10:
            return true
        case GuestOS_SLES_10_64:
            return true
        case GuestOS_SLES_11:
            return true
        case GuestOS_SLES_11_64:
            return true
        case GuestOS_SLES_12:
            return true
        case GuestOS_SLES_12_64:
            return true
        case GuestOS_SLES_15_64:
            return true
        case GuestOS_NLD_9:
            return true
        case GuestOS_OES:
            return true
        case GuestOS_SJDS:
            return true
        case GuestOS_MANDRAKE:
            return true
        case GuestOS_MANDRIVA:
            return true
        case GuestOS_MANDRIVA_64:
            return true
        case GuestOS_TURBO_LINUX:
            return true
        case GuestOS_TURBO_LINUX_64:
            return true
        case GuestOS_UBUNTU:
            return true
        case GuestOS_UBUNTU_64:
            return true
        case GuestOS_DEBIAN_4:
            return true
        case GuestOS_DEBIAN_4_64:
            return true
        case GuestOS_DEBIAN_5:
            return true
        case GuestOS_DEBIAN_5_64:
            return true
        case GuestOS_DEBIAN_6:
            return true
        case GuestOS_DEBIAN_6_64:
            return true
        case GuestOS_DEBIAN_7:
            return true
        case GuestOS_DEBIAN_7_64:
            return true
        case GuestOS_DEBIAN_8:
            return true
        case GuestOS_DEBIAN_8_64:
            return true
        case GuestOS_DEBIAN_9:
            return true
        case GuestOS_DEBIAN_9_64:
            return true
        case GuestOS_DEBIAN_10:
            return true
        case GuestOS_DEBIAN_10_64:
            return true
        case GuestOS_DEBIAN_11:
            return true
        case GuestOS_DEBIAN_11_64:
            return true
        case GuestOS_ASIANUX_3:
            return true
        case GuestOS_ASIANUX_3_64:
            return true
        case GuestOS_ASIANUX_4:
            return true
        case GuestOS_ASIANUX_4_64:
            return true
        case GuestOS_ASIANUX_5_64:
            return true
        case GuestOS_ASIANUX_7_64:
            return true
        case GuestOS_ASIANUX_8_64:
            return true
        case GuestOS_OPENSUSE:
            return true
        case GuestOS_OPENSUSE_64:
            return true
        case GuestOS_FEDORA:
            return true
        case GuestOS_FEDORA_64:
            return true
        case GuestOS_COREOS_64:
            return true
        case GuestOS_VMWARE_PHOTON_64:
            return true
        case GuestOS_OTHER_24X_LINUX:
            return true
        case GuestOS_OTHER_24X_LINUX_64:
            return true
        case GuestOS_OTHER_26X_LINUX:
            return true
        case GuestOS_OTHER_26X_LINUX_64:
            return true
        case GuestOS_OTHER_3X_LINUX:
            return true
        case GuestOS_OTHER_3X_LINUX_64:
            return true
        case GuestOS_OTHER_4X_LINUX:
            return true
        case GuestOS_OTHER_4X_LINUX_64:
            return true
        case GuestOS_OTHER_LINUX:
            return true
        case GuestOS_GENERIC_LINUX:
            return true
        case GuestOS_OTHER_LINUX_64:
            return true
        case GuestOS_SOLARIS_6:
            return true
        case GuestOS_SOLARIS_7:
            return true
        case GuestOS_SOLARIS_8:
            return true
        case GuestOS_SOLARIS_9:
            return true
        case GuestOS_SOLARIS_10:
            return true
        case GuestOS_SOLARIS_10_64:
            return true
        case GuestOS_SOLARIS_11_64:
            return true
        case GuestOS_OS2:
            return true
        case GuestOS_ECOMSTATION:
            return true
        case GuestOS_ECOMSTATION_2:
            return true
        case GuestOS_NETWARE_4:
            return true
        case GuestOS_NETWARE_5:
            return true
        case GuestOS_NETWARE_6:
            return true
        case GuestOS_OPENSERVER_5:
            return true
        case GuestOS_OPENSERVER_6:
            return true
        case GuestOS_UNIXWARE_7:
            return true
        case GuestOS_DARWIN:
            return true
        case GuestOS_DARWIN_64:
            return true
        case GuestOS_DARWIN_10:
            return true
        case GuestOS_DARWIN_10_64:
            return true
        case GuestOS_DARWIN_11:
            return true
        case GuestOS_DARWIN_11_64:
            return true
        case GuestOS_DARWIN_12_64:
            return true
        case GuestOS_DARWIN_13_64:
            return true
        case GuestOS_DARWIN_14_64:
            return true
        case GuestOS_DARWIN_15_64:
            return true
        case GuestOS_DARWIN_16_64:
            return true
        case GuestOS_DARWIN_17_64:
            return true
        case GuestOS_DARWIN_18_64:
            return true
        case GuestOS_DARWIN_19_64:
            return true
        case GuestOS_VMKERNEL:
            return true
        case GuestOS_VMKERNEL_5:
            return true
        case GuestOS_VMKERNEL_6:
            return true
        case GuestOS_VMKERNEL_65:
            return true
        case GuestOS_VMKERNEL_7:
            return true
        case GuestOS_AMAZONLINUX2_64:
            return true
        case GuestOS_CRXPOD_1:
            return true
        case GuestOS_OTHER:
            return true
        case GuestOS_OTHER_64:
            return true
        default:
            return false
    }
}




// The ``GuestOSFamily`` enumeration class defines the valid guest operating system family types reported by a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type GuestOSFamily string

const (
    // Windows operating system
     GuestOSFamily_WINDOWS GuestOSFamily = "WINDOWS"
    // Linux operating system
     GuestOSFamily_LINUX GuestOSFamily = "LINUX"
    // Novell Netware
     GuestOSFamily_NETWARE GuestOSFamily = "NETWARE"
    // Solaris operating system
     GuestOSFamily_SOLARIS GuestOSFamily = "SOLARIS"
    // Mac OS operating system
     GuestOSFamily_DARWIN GuestOSFamily = "DARWIN"
    // Other operating systems
     GuestOSFamily_OTHER GuestOSFamily = "OTHER"
)

func (g GuestOSFamily) GuestOSFamily() bool {
    switch g {
        case GuestOSFamily_WINDOWS:
            return true
        case GuestOSFamily_LINUX:
            return true
        case GuestOSFamily_NETWARE:
            return true
        case GuestOSFamily_SOLARIS:
            return true
        case GuestOSFamily_DARWIN:
            return true
        case GuestOSFamily_OTHER:
            return true
        default:
            return false
    }
}










