+++
title = "CoreDNS-1.8.5 Release"
description = "CoreDNS-1.8.5 Release Notes."
tags = ["Release", "1.8.5", "Notes"]
release = "1.8.5"
date = 2021-05-28T07:00:00+00:00
author = "coredns"
+++

Blah blah blah

## Brought to You By

Chris O'Haver,
Licht Takeuchi,
mfleader,
Miek Gieben,
Ondřej Benkovský,
Sven Nebel,
Yong Tang.

## Noteworthy Changes

* core: Add -p for port flag (https://github.com/coredns/coredns/pull/4653)
* core: Fix IPv6 case for CIDR format reverse zones (https://github.com/coredns/coredns/pull/4652)
* core: Share plugins among zones in the same server block (https://github.com/coredns/coredns/pull/4593)
* plugin/cache: Unset AD flag when DO is not set for cache miss (https://github.com/coredns/coredns/pull/4736)
* plugin/errors: add configurable log level to errors plugin (https://github.com/coredns/coredns/pull/4718)
* plugin/kubernetes: Add NS+hosts records to xfr response. Add coredns service to test data. (https://github.com/coredns/coredns/pull/4696)
* plugin/log: do not log NOERROR in log plugin when response is not available (https://github.com/coredns/coredns/pull/4725)
* plugin/log: fix closing of codeblock (https://github.com/coredns/coredns/pull/4680)
* plugin/metrics: when no response is written, fallback to status of next plugin in prometheus plugin (https://github.com/coredns/coredns/pull/4727)
* plugin/route53: Fix Route53 plugin cannot retrieve ECS Task Role (https://github.com/coredns/coredns/pull/4669)
* plugin/secondary: doc updates (https://github.com/coredns/coredns/pull/4686)
* plugin/secondary: Retry initial transfer until successful (https://github.com/coredns/coredns/pull/4663)
* plugin/trace: fix rcode tag in case of no response (https://github.com/coredns/coredns/pull/4742)
* plugin/trace: trace plugin can mark traces with error tag (https://github.com/coredns/coredns/pull/4720)
