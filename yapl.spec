# Copyright 2021 Hewlett Packard Enterprise Development LP
Name: yapl
License: MIT License
Summary: Simple pipeline for automation tasks
Version: %(cat .version)
Release: %(echo ${BUILD_METADATA})
Source: %{name}-%{version}.tar.bz2
Vendor: Cray Inc.
Provides: yapl
%ifarch x86_64
    %global GOARCH amd64
%endif
%description
Installs the Yapl GoLang binary onto a Linux system.

%prep
%setup -q

%build
CGO_ENABLED=0
GOOS=linux
GOARCH="%{GOARCH}"
GO111MODULE=on
export CGO_ENABLED GOOS GOARCH GO111MODULE

make build

%install
CGO_ENABLED=0
GOOS=linux
GOARCH="%{GOARCH}"
GO111MODULE=on
export CGO_ENABLED GOOS GOARCH GO111MODULE

mkdir -pv ${RPM_BUILD_ROOT}/usr/bin/
cp -pv bin/yapl ${RPM_BUILD_ROOT}/usr/bin/yapl

%clean

%files
%license LICENSE
%defattr(755,root,root)
/usr/bin/yapl

%changelog